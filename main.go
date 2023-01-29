package main

import (
	"log"
	"net/http"
	"runtime"

	"CMinting/controllers"
	"CMinting/initializers"
	"CMinting/routes"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	CollectionController      controllers.CollectionController
	CollectionRouteController routes.CollectionRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	CollectionController = controllers.NewCollectionController(initializers.DB)
	CollectionRouteController = routes.NewRouteCollectionController(CollectionController)

	server = gin.Default()
}

func main() {
	// config, err := initializers.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("🚀 Could not load environment variables", err)
	// }
	server.StaticFile("/57eca921daeb014452a8.module.wasm", "../Cardano-NFT-marketplace/build/57eca921daeb014452a8.module.wasm")
	server.LoadHTMLGlob("../Cardano-NFT-marketplace/build/index.html")
	server.Static("/static", "../Cardano-NFT-marketplace/build/static")
	server.Static("/images", "../Cardano-NFT-marketplace/build/images")
	server.Static("/monkeyimages", "../Cardano-NFT-marketplace/build/monkeyimages")
	
	
	// server.Static("/", "../CMinting_frontend/dist/")
	// server.StaticFile("/659.bundle.js", "../CMinting_frontend/dist/659.bundle.js")
	// server.StaticFile("/806.bundle.js", "../CMinting_frontend/dist/806.bundle.js")
	// server.StaticFile("/74ab6ba62b0c6b572c34.module.wasm", "../CMinting_frontend/dist/74ab6ba62b0c6b572c34.module.wasm")
	// server.GET("/", func(c *gin.Context) {
	// 	c.Redirect(http.StatusMovedPermanently, "public/")
	// })
	// server.StaticFS("/public", http.Dir("../CMinting_frontend/dist/"))

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	server.GET("/nftmint", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	server.GET("/createcollection", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	router := server.Group("/api")
	CollectionRouteController.CollectionRoute(router)

	log.Fatal(server.Run(":8000"))
}
