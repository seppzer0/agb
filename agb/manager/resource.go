package manager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"agb/config"
	"agb/tool"

	cerror "agb/error"
)

// ResourceManager manages resources such as toolchain, kernel source, PATHs etc.
type ResourceManager struct {
	ClangUrl        string
	SourceLocation  string
	directoryConfig *config.DirectoryConfig
	gitManager      *GitManager
	fileManager     *FileManager
}

// NewResourceManager creates new instance of ResourceManager.
func NewResourceManager(cu string, sloc string) *ResourceManager {
	dc := config.NewDirectoryConfig()
	gm := NewGitManager()
	fm := NewFileManager()

	return &ResourceManager{
		ClangUrl:        cu,
		SourceLocation:  sloc,
		directoryConfig: dc,
		gitManager:      gm,
		fileManager:     fm,
	}
}

// ExportToPath exports specified path to PATH.
func (rm *ResourceManager) ExportToPath(path string) error {
	pathVar := "PATH"
	updatedPath := fmt.Sprintf("%s:%s", os.Getenv(pathVar), path)

	os.Setenv(pathVar, updatedPath)
	if !strings.Contains(os.Getenv(pathVar), path) {
		return cerror.ErrGeneric{Message: fmt.Sprintf("Path %s not exported to PATH", path)}
	}

	return nil
}

// GetCompiler sets up Clang compiler.
func (rm *ResourceManager) GetCompiler() error {
	fmt.Printf("Downloading Clang from: %s\n", rm.ClangUrl)

	targz_name := filepath.Join(rm.directoryConfig.RootPath, "clang.tar.gz")
	clang_bin := filepath.Join(rm.directoryConfig.ClangPath, "bin")

	if _, err := os.Stat(targz_name); err != nil {
		if err := rm.fileManager.Download(rm.ClangUrl, targz_name); err != nil {
			return err
		}
	} else {
		tool.Mnote("Compiler tar.gz already downloaded\n")
	}

	if _, err := tool.RunCmd(fmt.Sprintf("%s --version", filepath.Join(clang_bin, "clang"))); err != nil {
		tool.Mnote("Cleaning dirty directory..")
		rm.fileManager.Delete(rm.directoryConfig.ClangPath)
		os.MkdirAll(rm.directoryConfig.ClangPath, os.ModePerm)
		tool.Mdone()

		tool.Mnote("Unpacking..")
		if err := rm.fileManager.UnpackTarGz(targz_name, rm.directoryConfig.ClangPath); err != nil {
			return err
		}
		tool.Mdone()
	} else {
		tool.Mnote("Compiler already unpacked and functional\n")
	}

	if err := rm.ExportToPath(clang_bin); err != nil {
		return err
	}

	return nil
}

// GetSource downloads kernel source.
func (rm *ResourceManager) GetSource(av int, lkv float64, pv string) error {
	if _, err := tool.RunCmd("repo --version"); err != nil {
		return cerror.ErrGeneric{Message: "repo tool not installed."}
	}

	if _, err := os.Stat(rm.directoryConfig.KernelSourcePath); err == nil {
		tool.Mnote("A kernel source directory already detected, using it for the build..\n")
	} else {
		tool.Mnote("Downloading kernel source..")
		os.MkdirAll(rm.directoryConfig.KernelSourcePath, os.ModePerm)

		// Google's GKI source is addressed if no custom URL is specified
		if rm.SourceLocation == "" {
			cmd_repo := fmt.Sprintf(
				"repo init --depth=1 --u https://android.googlesource.com/kernel/manifest -b common-android%d-%.1f --repo-rev=v2.16",
				av, lkv,
			)
			if pv != "" {
				cmd_repo = strings.Replace(
					cmd_repo,
					fmt.Sprintf("common-android%d-%.1f", av, lkv),
					fmt.Sprintf("common-android%d-%.1f-%s", av, lkv, pv),
					1,
				)
			}

			// NOTE: includes git config action to prevent repo from stalling;
			//       see https://groups.google.com/g/repo-discuss/c/T_JouBm-vBU/m/Jg1SLlRs2t4J
			procs, _ := tool.RunCmd("nproc --all")
			cmds := []string{
				"git config --global color.ui false",
				cmd_repo,
				"repo --version",
				fmt.Sprintf("repo --trace sync -c -j%s --no-tags", procs),
			}

			for _, cmd := range cmds {
				out, err := tool.RunCmdWDir(cmd, rm.directoryConfig.KernelSourcePath)
				if err != nil {
					return cerror.ErrCommandRun{Command: cmd, Output: out}
				}
			}
		} else {
			if err := rm.gitManager.Clone(rm.SourceLocation, rm.directoryConfig.KernelSourcePath, true); err != nil {
				return err
			}
		}
	}

	tool.Mdone()
	return nil
}

// CleanKernelSource cleans the directory with kernel sources from potential artifacts.
func (rm *ResourceManager) CleanKernelSource() error {
	return rm.gitManager.Reset(rm.directoryConfig.KernelSourcePath)
}
