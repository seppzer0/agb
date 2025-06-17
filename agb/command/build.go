package command

import (
	"agb/core"
	"agb/manager"
)

// BuildCommand is a representation of "build" command.
type BuildCommand struct {
	LinuxKernelVersion string
	AndroidVersion     int
	PatchVersion       string
	DefconfigPath      string
	ClangUrl           string
	SourceLocation     string
	ModulesUrl         string
	KernelSu           bool
}

// NewBuildCommand creates new instance of BuildCommand.
func NewBuildCommand(
	lkv string,
	av int,
	pv string,
	dp string,
	cu string,
	sloc string,
	murl string,
	ksu bool,
) Command {
	return &BuildCommand{
		LinuxKernelVersion: lkv,
		AndroidVersion:     av,
		PatchVersion:       pv,
		DefconfigPath:      dp,
		ClangUrl:           cu,
		SourceLocation:     sloc,
		ModulesUrl:         murl,
		KernelSu:           ksu,
	}
}

// Execute runs "build" command's logic.
func (bc *BuildCommand) Execute() error {
	// TODO: Better do it as DI
	resource_manager := manager.NewResourceManager(
		bc.ClangUrl,
		bc.SourceLocation,
	)

	// TODO: Better do it as DI
	kernel_builder := core.NewGkiBuilder(
		bc.LinuxKernelVersion,
		bc.AndroidVersion,
		bc.PatchVersion,
		bc.DefconfigPath,
		bc.SourceLocation,
		bc.ModulesUrl,
		bc.KernelSu,
		resource_manager,
	)

	if err := kernel_builder.Prepare(); err != nil {
		//return fmt.Errorf("%s returned: %v", logPrefix, err)
		return err
	}

	if err := kernel_builder.Build(); err != nil {
		return err
	}

	return nil
}
