package routes

import (
	"CMinting/controllers"

	"github.com/gin-gonic/gin"
)

type CollectionRouteController struct {
	collectionController controllers.CollectionController
}

func NewRouteCollectionController(collectionController controllers.CollectionController) CollectionRouteController {
	return CollectionRouteController{collectionController}
}

func (cc *CollectionRouteController) CollectionRoute(rg *gin.RouterGroup) {

	router := rg.Group("collection")
	router.POST("/", cc.collectionController.UploadCollection)
	router.GET("/", cc.collectionController.GetAllCollections)
	router.PUT("/minted", cc.collectionController.MintedCollection)
}