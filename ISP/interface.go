package isp

type Command interface {
	Execute()
}

type CommandWithInput interface {
	HandleInput()
}
