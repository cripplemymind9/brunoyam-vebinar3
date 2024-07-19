package main

import (
	"github.com/cripplemymind9/brunoyam-vebinar3/classwork/internal/repository"
	"github.com/cripplemymind9/brunoyam-vebinar3/classwork/internal/server"
	"log"
)

func main() {
	repo, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	server := server.NewServer(":8080", repo)
	
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}