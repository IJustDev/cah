package main

import (
	"github.com/royalzsoftware/cah/src"
)

type lightPlugin struct {
}

func (i lightPlugin) Init() {
	src.PlayerJoinedEvent.Register(func())
}

var CAHPlugin lightPlugin
