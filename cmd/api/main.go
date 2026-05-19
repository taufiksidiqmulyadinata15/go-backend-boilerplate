package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/config"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/database"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/server"
)

func main() {
	cfg := config.LoadConfig()

	db := database.ConnectDB(cfg)

	database.AutoMigrate(db)
	database.SeedAdmin(db)

	app := server.NewServer(cfg, db)

	go func() {
		log.Printf("Starting %s on port %s", cfg.AppName, cfg.AppPort)

		if err := app.Start(":" + cfg.AppPort); err != nil {
			log.Println("Server stopped:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get database instance:", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Println("Failed to close database connection:", err)
		return
	}

	log.Println("Server exited properly")
}
