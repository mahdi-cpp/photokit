package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-pkg/shared_model"
	"github.com/mahdi-cpp/photokit/internal/domain/model"
	"github.com/mahdi-cpp/photokit/internal/storage"
	"net/http"
)

type PersonHandler struct {
	userStorageManager *storage.MainStorageManager
}

func NewPersonsHandler(userStorageManager *storage.MainStorageManager) *PersonHandler {
	return &PersonHandler{
		userStorageManager: userStorageManager,
	}
}

func (handler *PersonHandler) Create(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var request shared_model.CollectionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	newItem, err := userStorage.PersonManager.Create(&model.Person{Title: request.Title})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	update := shared_model.AssetUpdate{AssetIds: request.AssetIds, AddPersons: []int{newItem.ID}}
	_, err = userStorage.UpdateAsset(update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userStorage.UpdateCollections()

	c.JSON(http.StatusCreated, shared_model.CollectionResponse{
		ID:    newItem.ID,
		Title: newItem.Title,
	})
}

func (handler *PersonHandler) Update(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var itemHandler model.PersonHandler
	if err := c.ShouldBindJSON(&itemHandler); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item, err := userStorage.PersonManager.Get(itemHandler.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	model.UpdatePerson(item, itemHandler)

	item2, err := userStorage.PersonManager.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *PersonHandler) Delete(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item model.Person
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err = userStorage.PersonManager.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "delete ok")
}

func (handler *PersonHandler) GetCollectionList(c *gin.Context) {
	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userStorage.PersonManager.GetAllSorted("creationDate", "1asc")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := shared_model.PHCollectionList[*model.Person]{
		Collections: make([]*shared_model.PHCollection[*model.Person], len(items)),
	}

	for i, item := range items {
		assets, _ := userStorage.PersonManager.GetItemAssets(item.ID)
		result.Collections[i] = &shared_model.PHCollection[*model.Person]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (handler *PersonHandler) GetCollectionListWith(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	// Get only visible items
	items, err := userStorage.PersonManager.GetList(func(a *model.Person) bool {
		return !a.IsCollection
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, items)
}
