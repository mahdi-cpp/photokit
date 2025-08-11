package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-pkg/shared_model"
	"github.com/mahdi-cpp/photokit/internal/domain/model"
	"github.com/mahdi-cpp/photokit/internal/storage"
	"net/http"
)

type SharedAlbumHandler struct {
	userStorageManager *storage.MainStorageManager
}

func NewSharedAlbumHandler(userStorageManager *storage.MainStorageManager) *SharedAlbumHandler {
	return &SharedAlbumHandler{
		userStorageManager: userStorageManager,
	}
}

func (handler *SharedAlbumHandler) Create(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item model.SharedAlbum
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item2, err := userStorage.SharedAlbumManager.Create(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *SharedAlbumHandler) Update(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var itemHandler model.SharedAlbumHandler
	if err := c.ShouldBindJSON(&itemHandler); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item, err := userStorage.SharedAlbumManager.Get(itemHandler.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	model.UpdateSharedAlbum(item, itemHandler)

	item2, err := userStorage.SharedAlbumManager.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *SharedAlbumHandler) Delete(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item model.SharedAlbum
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err = userStorage.SharedAlbumManager.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "delete ok")
}

func (handler *SharedAlbumHandler) GetList(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userStorage.SharedAlbumManager.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := shared_model.PHCollectionList[*model.SharedAlbum]{
		Collections: make([]*shared_model.PHCollection[*model.SharedAlbum], len(items)),
	}

	for i, item := range items {
		assets, _ := userStorage.SharedAlbumManager.GetItemAssets(item.ID)
		result.Collections[i] = &shared_model.PHCollection[*model.SharedAlbum]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
