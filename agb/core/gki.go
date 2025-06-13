package core

import (
	cerror "agb/error"
	"agb/manager"
	"agb/tool"
	"strconv"
)

// GkiBuilder is a core module responsible for building GKI.
type GkiBuilder struct {
	LinuxKernelVersion string
	AndroidVersion     int
	PatchVersion       string
	DefconfigPath      string
	SourceLocation     string
	ModulesPath        string
	KernelSu           bool
	resourceManager    *manager.ResourceManager
}

// NewGkiBuilder creates new instance of GkiBuilder.
func NewGkiBuilder(
	lkv string,
	av int,
	pv string,
	dp string,
	sloc string,
	murl string,
	ksu bool,
	rm *manager.ResourceManager,
) *GkiBuilder {
	return &GkiBuilder{
		LinuxKernelVersion: lkv,
		AndroidVersion:     av,
		PatchVersion:       pv,
		DefconfigPath:      dp,
		SourceLocation:     sloc,
		ModulesPath:        murl,
		KernelSu:           ksu,
		resourceManager:    rm,
	}
}

// AddKsu introduces KernelSU support into the kernel (GKI mode).
func (gb *GkiBuilder) addKsu() error {
	return nil
}

// Prepare runs all the preparations for the build.
func (gb *GkiBuilder) Prepare() error {
	// by default, assume that Clang is not required
	clang_required := false

	lkv_float, err := strconv.ParseFloat(gb.LinuxKernelVersion, 32)
	if err != nil {
		return cerror.ErrGeneric{Message: "Could not convert kernel version to float"}
	}

	// for regular GKI sources, separate Clang is not required
	if !(lkv_float >= 5.10) {
		clang_required = true
		if err := gb.resourceManager.GetCompiler(); err != nil {
			return err
		}
	}

	if err := gb.resourceManager.ValidateEnv(clang_required); err != nil {
		return err
	}

	if err := gb.resourceManager.CleanArtifacts(); err != nil {
		return err
	}

	if err := gb.resourceManager.GetSource(gb.AndroidVersion, gb.LinuxKernelVersion, gb.PatchVersion); err != nil {
		return err
	}

	return nil
}

// Build launches GKI kernel build.
func (gb *GkiBuilder) Build() error {
	var cmd string

	if gb.AndroidVersion >= 14 {
		cmd = "tools/bazel run --disk_cache=~/.cache/bazel --config=fast --config=stamp --lto=thin //common:kernel_aarch64_dist -- --dist_dir=dist"
	} else {
		cmd = "LTO=thin BUILD_CONFIG=common/build.config.gki.aarch64 build/build.sh CC=\"clang\""
	}

	out, err := tool.RunCmdWDir(cmd, "kernel_source")
	if err != nil {
		return cerror.ErrCommandRun{Command: cmd, Output: out}
	}

	return nil
}

// Packages collects the kernel binary.
func (gb *GkiBuilder) Package() error {
	return nil
}
