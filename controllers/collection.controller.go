package controllers

import (
	"CMinting/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gorm.io/gorm"
	"strconv"
)

type CollectionController struct {
	DB *gorm.DB
}

func NewCollectionController(DB *gorm.DB) CollectionController {
	return CollectionController{DB}
}

func (cc *CollectionController) UploadCollection(ctx *gin.Context) {
	log.Print("UploadCollection")

	id := uuid.New()
	name := ctx.PostForm("name")
	description := ctx.PostForm("description")
	recipient_addr := ctx.PostForm("payout_wallet")
	fee_type := ctx.PostForm("fee_type")
	fee := ctx.PostForm("fee")
	collection_type := ctx.PostForm("collection_type")
	collection_ipfs := ctx.PostForm("collection_ipfs")
	representative_ipfs := ctx.PostForm("representative_ipfs")
	collection_count := ctx.PostForm("collection_count")
	collection_price := ctx.PostForm("collection_price")
	
	now := time.Now()

	minted_info := ""
	count, err := strconv.Atoi(collection_count)
	if err != nil {
		panic(err)
	}
	for i := 0; i < count; i++ {
		minted_info += "A"
	}
			
	newCollection := models.Collection{
		ID:                 id,
		Name:               name,
		Description:        description,
		RecipientAddr:      recipient_addr,
		FeeType:            fee_type,
		Fee:                fee,
		Type:				collection_type,
		CollectionIpfs:     collection_ipfs,
		RepresentativeIpfs: representative_ipfs,
		Count:              collection_count,
		Price:              collection_price,
		MintedInfo:			minted_info,
		CreatedAt:          now,
	}

	log.Print("UploadCollection_newCollection--->", newCollection)
	result := cc.DB.Create(&newCollection)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"collection": &newCollection,
	})
}

func (cc *CollectionController) GetAllCollections(ctx *gin.Context) {

	collections := []models.Collection{}

	result := cc.DB.Find(&collections)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":      "success",
		"collections": collections,
	})
}

func (cc *CollectionController) MintedCollection(ctx *gin.Context) {
	id, _ := uuid.Parse(ctx.PostForm("id"))
	minted_info := ctx.PostForm("minted_info")

	log.Print("minted_info-------------------->", minted_info)

	currentCollection := models.Collection{}
	result := cc.DB.First(&currentCollection, "ID = ?", id)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error})
		return
	}

	currentCollection.MintedInfo = minted_info;
	
	cc.DB.Save(&currentCollection);

	ctx.JSON(http.StatusOK, gin.H{
		"status": 		"success",
		"collection":	currentCollection,
	})
}