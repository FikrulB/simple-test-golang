package main

import (
	"fmt"
	"golang-test/cmd"
	"golang-test/cmd/migrate"
	"golang-test/cmd/seeder"
	"golang-test/configs"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "migrate":
			migrate.Run()
		case "seeder":
			seeder.Run()
		default:
			fmt.Println("Unknown command:", command)
		}
		return
	}

	env := configs.NewEnv()

	e := cmd.InitializeRouter()

	port := fmt.Sprintf(":%s", env.AppPort)
	if err := e.Start(port); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
