package command

type ICommand interface {
	Execute() error
}
