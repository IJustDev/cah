package main

import (
	"fmt"

	"github.com/royalzsoftware/cah/src"
)

type ijustdevPlugin struct {
}

func (i ijustdevPlugin) Init() {
	fmt.Println("Hello World")

	src.PlayerJoinedEvent.Register(func(player src.Player) {
		fmt.Println("Player joined: " + player.Name)
	})

}

var CAHPlugin ijustdevPlugin
