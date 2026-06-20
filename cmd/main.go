package main

import (
	"log/slog"
	"pizza-tracker-go/internal/models"
    "os"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dbModel, err := models.InitDB(cfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialized database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database Initialized Successfully")

	RegisterCustomValidators()

	h := NewHandler(dbModel)

	router := gin.Default()

	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to load Templates", "error", err)
		os.Exit(1)
	}
	setupRoutes(router, h)

	slog.Info()("Server Starting", "url", "http://localhost:"+cfg.Port)

	router.Run(":" +cfg.Port)
}
