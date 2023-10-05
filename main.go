package main

import (
	routes "auth/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run()
}