package main

import "ProjectAnalysis/pkg/server"

func main() {
	server.Run(&server.AverageServer{})
}
