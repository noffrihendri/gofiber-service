package router

import (
	"faber/databases"
	"faber/internal/handlers"
	"faber/internal/repositories"
	"faber/internal/usecases"
	"log"
	"os"
	"runtime/debug"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Router() *fiber.App {
	app := fiber.New()

	// Initialize dependencies
	db := databases.NewDBPostgres()
	productRepository := repositories.NewProductRepository(db)
	productUsecase := usecases.NewProductUsecase(productRepository)
	productHandler := handlers.NewProductHandler(productUsecase)
	app.Use(customRecover())
	//app.Use(recover.New())

	app.Get("/api/list", productHandler.GetProduk)
	app.Get("/api/list/:id", productHandler.GetProductById)
	app.Get("/metrics", monitor.New())

	// Custom 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		// Set the status code to 404
		c.Status(fiber.StatusNotFound)
		// Return a custom 404 JSON response
		return c.JSON(fiber.Map{
			"error":   "Resource not found",
			"message": "The requested URL was not found on this server",
		})
	})

	return app
}

// Custom recover middleware to log panics to a file
func customRecover() fiber.Handler {
	now := time.Now()
	logName := "-" + now.Format("2006-02-01-15")
	// Open log file
	file, err := os.OpenFile("./logruntime/panic/errors"+logName+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// Create a new logger
	logger := log.New(file, "", log.LstdFlags)

	return func(c *fiber.Ctx) error {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic and stack trace
				logger.Printf("panic: %v\n%s", r, debug.Stack())
				// Return a 500 status code and JSON response
				c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "Internal Server Error",
					"message": "An unexpected error occurred",
				})
			}
		}()
		return c.Next()
	}
}
