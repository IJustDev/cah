package main

import (
	cah "github.com/royalzsoftware/cah/src"
	ws "github.com/royalzsoftware/cah/ws"
)

func main() {

	f := &cah.PluginManager{}
	f.LoadAllPlugins("./")

	e := cah.NewEventListener()
	e.Subscribe()

	ws.Start()
}
