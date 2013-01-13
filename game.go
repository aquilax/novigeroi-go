package main

type Game struct {
	running bool
	server  *Server
	db *Database
	user *User
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() {
	// Connect to database
	g.db = NewDatabase()
	g.db.Connect()

	// Start HTTP Server
	g.server = NewServer()

	g.user = NewUser()
	g.user.Init()

	g.server.StartServer()

	// It's all good
	g.running = true
}
