package handlers

import (
	"faber/internal/infrastructures"
	"faber/internal/usecases"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type ProductHandler struct {
	ProductUsecase usecases.ProductUsecaseContract
	name           string
}

type ProductHandlerContract interface {
	GetProduk(w *fiber.Ctx) error
	GetProductById(w *fiber.Ctx) error
}

func NewProductHandler(productHandler usecases.ProductUsecaseContract) ProductHandlerContract {
	return &ProductHandler{
		name:           "Product Handler",
		ProductUsecase: productHandler,
	}
}

var (
	appLogger *infrastructures.LogDir
)

func (p *ProductHandler) GetProduk(c *fiber.Ctx) error {
	fmt.Println("masuk disini")

	result, err := p.ProductUsecase.GetProduct()

	if err != nil {
		c.Status(http.StatusNotFound)
		log.Info("error log", err)
		log.Debugf("[%s][Read] is executed\n", err)
		data := map[string]interface{}{
			"responseCode":    "200",
			"responseMessage": err,
		}
		return c.JSON(data)
	}

	data := map[string]interface{}{
		"responseCode":    "200",
		"responseMessage": result,
	}
	return c.JSON(data)
}

func (p *ProductHandler) GetProductById(c *fiber.Ctx) error {
	result, err := p.ProductUsecase.GetProductById(c.Params("id"))
	if err != nil {
		c.Status(http.StatusNotFound)
		// data := map[string]interface{}{
		// 	"responseCode":    "404",
		// 	"responseMessage": "Product Not Found",
		// }
		//	return c.JSON(data)
	}
	data := map[string]interface{}{
		"responseCode":    "200",
		"responseMessage": result.ID,
	}
	return c.JSON(data)

}
