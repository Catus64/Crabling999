package gitrepo

import (
	"fmt"
	"time"

	//"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5"
)

// Repo wraps the go-git repository so the rest of the app
// never imports go-git directly.
type Repo struct {
	repo *git.Repository
}

// Open opens the repository at the given path.
func Open(path string) (*Repo, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("open repo %q: %w", path, err)
	}
	return &Repo{repo: r}, nil
}

type Commit struct {
	Hash    string    `json:"hash"`
	Author  string    `json:"author"`
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
}

// ListCommits returns up to limit commits, newest first, starting from HEAD.
func (r *Repo) ListCommits(limit int) ([]Commit, error) {
	iter, err := r.repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, fmt.Errorf("read log: %w", err)
	}
	defer iter.Close()

	var commits []Commit
	for len(commits) < limit {
		c, err := iter.Next()
		if err != nil {
			break // io.EOF means no more commits
		}
		commits = append(commits, Commit{
			Hash:    c.Hash.String(),
			Author:  c.Author.Name,
			Message: c.Message,
			Date:    c.Author.When,
		})
	}
	return commits, nil
}
