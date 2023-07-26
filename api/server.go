package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/paulonc/go-products/db/sqlc"
)

type Server struct {
	store  *db.ExecuteStore
	router *gin.Engine
}

func InstanceServer(store *db.ExecuteStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/product", server.createProduct)
	router.GET("/product/:id", server.getProduct)
	router.GET("/products", server.getProducts)
	router.PUT("/product", server.updateProduct)
	router.DELETE("/product/:id", server.deleteProduct)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"api has one error:": err.Error()}
}