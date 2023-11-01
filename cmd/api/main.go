package main

import "github.com/CrysLef/cripto-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}
