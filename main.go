package main

import (
	"log"
	"os"
	"sync"

	"github.com/Ndraaa15/musiku/cmd/server"
	"github.com/joho/godotenv"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		os.Exit(server.Run())
	}()
	wg.Wait()
}
