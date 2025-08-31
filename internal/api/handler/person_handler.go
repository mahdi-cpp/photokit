package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/photokit/internal/helpers"

	"github.com/mahdi-cpp/photokit/internal/application"
	collection "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/person"
	"github.com/mahdi-cpp/photokit/internal/collections/phasset"
)

type PersonHandler struct {
	manager *application.AppManager
}

func NewPersonsHandler(manager *application.AppManager) *PersonHandler {
	return &PersonHandler{
		manager: manager,
	}
}

func (handler *PersonHandler) Create(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var request collection.CollectionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	newItem, err := userManager.GetCollections().Persons.Collection.Create(&person.Person{Title: request.Title})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	updateOptions := phasset.UpdateOptions{
		AssetIds:   request.AssetIds,
		AddPersons: []string{newItem.ID},
	}
	_, err = userManager.UpdateAssets(updateOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userManager.UpdateCollections()

	//c.JSON(http.StatusCreated, CollectionResponse{
	//	ID:    newItem.ID,
	//	Title: newItem.Title,
	//})
}

func (handler *PersonHandler) Update(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var updateOptions person.UpdateOptions
	if err := c.ShouldBindJSON(&updateOptions); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	item, err := userManager.GetCollections().Persons.Collection.Get(updateOptions.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	person.Update(item, updateOptions)

	item2, err := userManager.GetCollections().Persons.Collection.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *PersonHandler) Delete(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var item person.Person
	if err := c.ShouldBindJSON(&item); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	err = userManager.GetCollections().Persons.Collection.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "delete ok")
}

func (handler *PersonHandler) GetCollectionList(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userManager.GetCollections().Persons.Collection.GetAllSorted("creationDate", "1asc")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result := collection.PHCollectionList[*person.Person]{
		Collections: make([]*collection.PHCollection[*person.Person], len(items)),
	}

	for i, item := range items {
		assets, _ := userManager.GetCollections().Persons.PhotoAssetList[item.ID]
		result.Collections[i] = &collection.PHCollection[*person.Person]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (handler *PersonHandler) GetCollectionListWith(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	// Get only visible items
	items, err := userManager.GetCollections().Persons.Collection.GetList(func(a *person.Person) bool {
		return !a.IsCollection
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, items)
}
