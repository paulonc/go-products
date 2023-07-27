package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/paulonc/go-products/backend/db/sqlc"
)

type createProductRequest struct {
	Name  string `json:"name" binding:"required"`
	Price int32  `json:"price" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Price <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "The price must be greater than zero."})
		return
	}

	arg := db.CreateProductParams{
		Name:  req.Name,
		Price: req.Price,
	}

	product, err := server.store.CreateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

type getProductRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

type deleteProductRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteProduct(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, true)
}

type updateProductRequest struct {
	ID    int32  `json:"id" binding:"required"`
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var req updateProductRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateProductParams{
		ID:    req.ID,
		Name:  req.Name,
		Price: req.Price,
	}

	product, err := server.store.UpdateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (server *Server) getProducts(ctx *gin.Context) {
	products, err := server.store.GetProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}