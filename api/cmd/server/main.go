package main

import (
	"log/slog"
	"os"

	"github.com/Catus64/repo-observer/api/internal/router"
)

func setupLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
}

func main() {

	setupLogger()

	slog.Info("starting server", "repo_path", os.Getenv("REPO_PATH"))

	r := router.New()

	if err := r.Run(":8080"); err != nil {
		slog.Error("failed to run server", "error", err)
		os.Exit(1)
	}

}
