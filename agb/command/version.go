package command

import (
	"agb/config"
	"fmt"
)

// VersionCommand is a representation of "version" command.
type VersionCommand struct{}

// NewVersionCommand creates new instance of VersionCommand
func NewVersionCommand() *VersionCommand {
	return &VersionCommand{}
}

// Execute runs "version" command's logic.
func (vc *VersionCommand) Execute() error {
	version_config := config.NewVersionConfig()
	app_version, go_version := version_config.AppVersion, version_config.GoVersion

	fmt.Printf("agb version %s, built with %s\n", app_version, go_version)
	return nil
}
