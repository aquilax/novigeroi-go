package main

type Game struct {
	running bool
	server *Server
}

func NewGame () *Game {
	return &Game{}
}

func (g *Game) Start() {
	g.running = true;
	g.server = NewServer()
	g.server.StartServer();
}
