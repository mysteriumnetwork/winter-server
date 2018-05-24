package main

import (
	"../core"
	"log"
)

type CommandStart struct {
	GameCommand
}

func (CommandStart) Execute(g *Game, connCommand core.ConnCommand, io core.CommandIO)  {
	args := connCommand.Split()
	clientName := args[1]
	g.SetClientName(io, clientName)
	log.Println("Set client name:", clientName)
	if !g.isStarted {
		if core.LOG_INFO {
			log.Println("Starting game")
		}
		g.isStarted = true
		go g.ShowZombies(io)
		go g.MoveZombies(io)
	}
}