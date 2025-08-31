package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/photokit/internal/helpers"

	"github.com/mahdi-cpp/photokit/internal/application"
	collection "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/shared_album"
)

type SharedAlbumHandler struct {
	manager *application.AppManager
}

func NewSharedAlbumHandler(manager *application.AppManager) *SharedAlbumHandler {
	return &SharedAlbumHandler{
		manager: manager,
	}
}

func (handler *SharedAlbumHandler) Create(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var item shared_album.SharedAlbum
	if err := c.ShouldBindJSON(&item); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item2, err := userManager.GetCollections().SharedAlbums.Collection.Create(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *SharedAlbumHandler) Update(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var updateOptions shared_album.UpdateOptions
	if err := c.ShouldBindJSON(&updateOptions); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item, err := userManager.GetCollections().SharedAlbums.Collection.Get(updateOptions.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	shared_album.Update(item, updateOptions)

	item2, err := userManager.GetCollections().SharedAlbums.Collection.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *SharedAlbumHandler) Delete(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var item shared_album.SharedAlbum
	if err := c.ShouldBindJSON(&item); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err = userManager.GetCollections().SharedAlbums.Collection.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "delete ok")
}

func (handler *SharedAlbumHandler) GetList(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userManager.GetCollections().SharedAlbums.Collection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := collection.PHCollectionList[*shared_album.SharedAlbum]{
		Collections: make([]*collection.PHCollection[*shared_album.SharedAlbum], len(items)),
	}

	for i, item := range items {
		assets, _ := userManager.GetCollections().SharedAlbums.PhotoAssetList[item.ID]
		result.Collections[i] = &collection.PHCollection[*shared_album.SharedAlbum]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
