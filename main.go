package main

import (
	"context"
	"faber/internal/infrastructures"
	"faber/internal/router"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

const (
	SERVERPORT = ":8083"
)

var (
	appLogger *infrastructures.LogDir
)

func main() {

	app := router.Router()
	//app.Use(customRecover())
	// Changing TimeZone & TimeFormat
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${ip} ${locals:requestid} ${pid} ${status} ${latency}  ${method} ${path} ${error}\n",
		TimeFormat: "2006/01/02 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	// Custom File Writer
	now := time.Now()
	logName := "applog-" + now.Format("2006-02-01-15")
	file, err := os.OpenFile("./logruntime/"+logName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output: file,
	}))

	log.SetOutput(file)

	go func() {
		if err := app.Listen(SERVERPORT); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()
	gracefulShutdown(app)
}

func gracefulShutdown(server *fiber.App) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("\nShutting down gracefully...")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(); err != nil {
		fmt.Printf("Error shutting down: %v\n", err)
	}
	fmt.Println("Running cleanup tasks...")
	os.Exit(0)
}
