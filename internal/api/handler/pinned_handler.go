package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/mahdi-cpp/photokit/internal/application"
	collection "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/pinned"
	"net/http"
	"sort"
)

type PinnedHandler struct {
	manager *application.AppManager
}

func NewPinnedHandler(manager *application.AppManager) *PinnedHandler {
	return &PinnedHandler{
		manager: manager,
	}
}

func (handler *PinnedHandler) Create(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item pinned.Pinned
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item2, err := userManager.GetCollections().Pinned.Collection.Create(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *PinnedHandler) Update(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var updateOptions pinned.UpdateOptions
	if err := c.ShouldBindJSON(&updateOptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item, err := userManager.GetCollections().Pinned.Collection.Get(updateOptions.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	pinned.Update(item, updateOptions)

	item2, err := userManager.GetCollections().Pinned.Collection.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *PinnedHandler) Delete(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item pinned.Pinned
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err = userManager.GetCollections().Pinned.Collection.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "Delete item with id:"+item.ID)
}

func (handler *PinnedHandler) GetList(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userManager.GetCollections().Pinned.Collection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	items = sortPinnedCollectionByIndex(items)

	// Create collection list without interface constraint
	result := collection.PHCollectionList[*pinned.Pinned]{
		Collections: make([]*collection.PHCollection[*pinned.Pinned], len(items)),
	}

	for i, item := range items {
		assets, _ := userManager.GetCollections().Pinned.PhotoAssetList[item.ID]
		result.Collections[i] = &collection.PHCollection[*pinned.Pinned]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (handler *PinnedHandler) GetCollectionListWith(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	// Get only visible items
	items, err := userManager.GetCollections().Pinned.Collection.GetList(func(a *pinned.Pinned) bool {
		return a.Icon == ""
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func sortPinnedCollectionByIndex(items []*pinned.Pinned) []*pinned.Pinned {
	sort.Slice(items, func(i, j int) bool {
		a := items[i]
		b := items[j]
		return a.Index < b.Index
	})

	return items
}
