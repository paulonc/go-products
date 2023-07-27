package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/paulonc/go-products/backend/db/sqlc"
)

type Server struct {
	store  *db.ExecuteStore
	router *gin.Engine
}

func InstanceServer(store *db.ExecuteStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"}
	router.Use(cors.New(config))


	router.POST("/product", server.createProduct)
	router.GET("/product/:id", server.getProduct)
	router.GET("/products", server.getProducts)
	router.PUT("/product/:id", server.updateProduct)
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