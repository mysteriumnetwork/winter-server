package main

import "../client/gamer"

type dtoStart struct {
	SessionID int
	Error  string
}

type dtoStatus struct {
	Width  int
	Height int
	Commands []string
	Game gamer.Gamer
	Error string
}
