package isp

type ICommand interface {
	Execute()
}

type ICommandWithInput interface {
	ICommand
	HandleInput()
}
