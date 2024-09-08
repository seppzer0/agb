package core

import (
	cerror "agb/error"
	"agb/manager"
	"agb/tool"
)

// GkiBuilder is a core module responsible for building GKI.
type GkiBuilder struct {
	LinuxKernelVersion float64
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
	lkv float64,
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

// validateEnd checks that everything necessary is present in build environment.
func (gb *GkiBuilder) validateEnv() bool {
	return true
}

// CleanEnvironment resets the build environment.
func (gb *GkiBuilder) CleanEnvironment() error {
	return nil
}

// patchAnyKernel3 patches AnyKernel3's files to package the kernel image.
func (gb *GkiBuilder) patchAnykernel3() error {
	return nil
}

// AddKsu introduces KernelSU support into the kernel (GKI mode).
func (gb *GkiBuilder) addKsu() error {
	return nil
}

// determineKernelVersion determines Linux kernel version directly from sources.
func (gb *GkiBuilder) determineKernelVersion() float64 {
	return 3.14
}

// Patch applies all necessary modifications for the kernel build.
func (gb *GkiBuilder) Patch() error {
	return nil
}

// Prepare runs all the preparations for the build.
func (gb *GkiBuilder) Prepare() error {
	// for regular GKI sources, separate Clang is not required
	if !(gb.LinuxKernelVersion >= 5.10) {
		if err := gb.resourceManager.GetCompiler(); err != nil {
			return err
		}
	}

	if err := gb.resourceManager.CleanKernelSource(); err != nil {
		return err
	}

	if err := gb.resourceManager.GetSource(gb.AndroidVersion, gb.LinuxKernelVersion, gb.PatchVersion); err != nil {
		return err
	}

	//if gb.determineKernelVersion() != gb.LinuxKernelVersion {
	//	return cerror.ErrGeneric{Message: "Specified Linux Kernel version does not match the actual one."}
	//}

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

	out, err := tool.RunCmd(cmd)
	if err != nil {
		return cerror.ErrCommandRun{Command: cmd, Output: out}
	}

	return nil
}

// Packages collects the kernel binary.
func (gb *GkiBuilder) Package() error {
	return nil
}
