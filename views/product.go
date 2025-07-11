package views

import (
	"net/http"
	"strconv"

	"github.com/farhapartex/ainventory/controller"
	"github.com/farhapartex/ainventory/dto"
	"github.com/farhapartex/ainventory/utils"
	"github.com/gin-gonic/gin"
)

func ProductCategoryListAPIView(ctx *gin.Context, ac *controller.AuthController) {
	_, err := utils.GetAuthenticatedUser(ctx)
	if err != nil {
		utils.HandleAuthError(ctx, err)
		return
	}

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	resp, err := ac.ProductCategoryList(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func ProductCategoryCreateAPIView(ctx *gin.Context, ac *controller.AuthController) {
	var request dto.ProductCategoryRequestDTO
	if err := ctx.ShouldBindJSON((&request)); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	response, err := ac.CreateProductCategoryController(request)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func ProductCategoryUpdateAPIView(ctx *gin.Context, ac *controller.AuthController) {
	categoryIDStr := ctx.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	var request dto.ProductCategoryRequestDTO
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid input",
		})
		return
	}

	response, err := ac.UpdateProductCategoryController(uint(categoryID), request)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func ProductCategoryDeleteAPIView(ctx *gin.Context, ac *controller.AuthController) {
	categoryIDStr := ctx.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid category ID",
		})
		return
	}

	err = ac.DeleteProductCategoryController(uint(categoryID))
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}

func SupplierListAPIView(ctx *gin.Context, ac *controller.AuthController) {
	// user, err := utils.GetAuthenticatedUser(ctx)
	// if err != nil {
	// 	utils.HandleAuthError(ctx, err)
	// 	return
	// }
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	queryParams := dto.ListQueryDTO{
		Page:     page,
		PageSize: pageSize,
		Status:   ctx.Query("status"),
		Search:   ctx.Query("search"),
		SortBy:   ctx.Query("sortBy"),
		SortDir:  ctx.Query("sortDir"),
	}

	resp, err := ac.SupplierList(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
