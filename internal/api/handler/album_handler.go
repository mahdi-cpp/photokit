package handler

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/mahdi-cpp/photokit/internal/application"
	collection "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/album"
	"github.com/mahdi-cpp/photokit/internal/collections/phasset"
	"net/http"
)

type AlbumHandler struct {
	manager *application.AppManager
}

func NewAlbumHandler(manager *application.AppManager) *AlbumHandler {
	return &AlbumHandler{
		manager: manager,
	}
}

func (handler *AlbumHandler) Create(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var request collection.CollectionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	newItem, err := userManager.GetCollections().Album.Collection.Create(&album.Album{Title: request.Title})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	update := phasset.UpdateOptions{
		AssetIds:  request.AssetIds,
		AddAlbums: []string{newItem.ID},
	}
	_, err = userManager.UpdateAssets(update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userManager.UpdateCollections()

	c.JSON(http.StatusCreated, CollectionResponse{
		ID:    newItem.ID,
		Title: newItem.Title,
	})
}

type CollectionResponse struct {
	ID    string `json:"id"`
	Title string `json:"name"`
}

func (handler *AlbumHandler) Update(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var updateOptions album.UpdateOptions
	if err := c.ShouldBindJSON(&updateOptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item, err := userManager.GetCollections().Album.Collection.Get(updateOptions.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	album.Update(item, updateOptions)

	item2, err := userManager.GetCollections().Album.Collection.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *AlbumHandler) Delete(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item album.Album
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err = userManager.GetCollections().Album.Collection.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "delete ok")
}

func (handler *AlbumHandler) GetAll(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userManager.GetCollections().Album.Collection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := collection.PHCollectionList[*album.Album]{
		Collections: make([]*collection.PHCollection[*album.Album], len(items)),
	}

	for i, item := range items {
		assets, _ := userManager.GetCollections().Album.PhotoAssetList[item.ID]
		result.Collections[i] = &collection.PHCollection[*album.Album]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (handler *AlbumHandler) GetBySearchOptions(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var searchOptions album.SearchOptions
	if err := c.ShouldBindJSON(&searchOptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		fmt.Println("Invalid request")
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userManager.GetCollections().Album.Collection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	filterItems := album.Search(items, searchOptions)

	result := collection.PHCollectionList[*album.Album]{
		Collections: make([]*collection.PHCollection[*album.Album], len(filterItems)),
	}

	for i, item := range items {
		assets, _ := userManager.GetCollections().Album.PhotoAssetList[item.ID]
		result.Collections[i] = &collection.PHCollection[*album.Album]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
