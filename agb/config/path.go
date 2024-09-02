package config

import (
	"os"
	"path/filepath"
)

// DirectoryConfig represents a config containing internal paths to multiple resources.
type DirectoryConfig struct {
	RootPath         string
	KernelSourcePath string
	KernelBuildPath  string
	ClangPath        string
	Anykernel3Path   string
}

// NewDirectoryConfig returns new instance of DirectoryConfig
func NewDirectoryConfig() *DirectoryConfig {
	// get the root path and use it as a base for everything
	root, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	return &DirectoryConfig{
		RootPath:         root,
		KernelSourcePath: filepath.Join(root, "kernel_source"),
		KernelBuildPath:  filepath.Join(root, "kernel_build"),
		ClangPath:        filepath.Join(root, "clang"),
		Anykernel3Path:   filepath.Join(root, "anykernel3"),
	}
}
