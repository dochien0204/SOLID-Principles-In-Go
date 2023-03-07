package main

import (
	isp "solid-principle/ISP"
	lsp "solid-principle/LSP"
	ocp "solid-principle/OCP"
	srp "solid-principle/SRP"
)

func main() {
	srp.Run()
	ocp.Run()
	lsp.Run()
	isp.Run()
}
