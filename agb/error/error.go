package error

import "fmt"

// ErrGeneric represents errors of a generic type.
type ErrGeneric struct {
	Message string
}

func (err ErrGeneric) Error() string {
	return fmt.Sprint(err.Message)
}

func (err ErrGeneric) Is(target error) bool {
	_, ok := target.(ErrGeneric)
	return ok
}

// ErrCommanRun represents errors for shell-run commands.
type ErrCommandRun struct {
	Command string
	Output  string
}

func (err ErrCommandRun) Error() string {
	return fmt.Sprintf("Could not execute '%s' command.", err.Command)
}

func (err ErrCommandRun) Is(target error) bool {
	_, ok := target.(ErrCommandRun)
	return ok
}
