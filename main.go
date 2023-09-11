package main

import (
	"log"
	"os"
	"sync"

	"github.com/Ndraaa15/musiku/cmd/server"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("[musiku-main] failed to load .env file. Error : %v\n", err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		os.Exit(server.Run())
	}()
	wg.Wait()
}
