package command

// CleanCommand is a representation of "version" command.
type CleanCommand struct{}

// NewCleanCommand creates new instance of CleanCommand
func NewCleanCommand() *CleanCommand {
	return &CleanCommand{}
}
