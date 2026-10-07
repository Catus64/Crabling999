package main

import (
	"log"
	"os"

	"github.com/Catus64/repo-observer/api/internal/gitrepo"
	"github.com/Catus64/repo-observer/api/internal/router"
)

func main() {
	repoPath := os.Getenv("REPO_PATH")
	if repoPath == "" {
		log.Fatal("REPO_PATH is not set")
	}

	repo, err := gitrepo.Open(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	r := router.New(repo)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
