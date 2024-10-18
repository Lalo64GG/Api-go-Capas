package  controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	
	"firstAPI-go/src/services"
	"firstAPI-go/src/models"
)

type ProductController struct {
	 ProductService *services.ProductService
}

func NewProductController(productService *services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}


func (ctrl *ProductController) CreateProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBind(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Reponse{
			Success: false,
            Message: "Invalid Request",
			Error: err.Error(),
			Data: nil,
		})
	}

	err := ctrl.ProductService.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Reponse{
			Success: false,
			Message: "Error in the server",
			Error: err.Error(),
            Data: nil,
		})
	}

	ctx.JSON(http.StatusOK, models.Reponse{
		Success: true,
		Message: "Product created successfully",
        Error: "",
        Data: product,
	})

}