package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/mahdi-cpp/photokit/internal/application"
	collection "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/phasset"
	"github.com/mahdi-cpp/photokit/internal/collections/trip"
	"net/http"
)

type TripHandler struct {
	manager *application.AppManager
}

func NewTripHandler(manager *application.AppManager) *TripHandler {
	return &TripHandler{
		manager: manager,
	}
}

func (handler *TripHandler) Create(c *gin.Context) {

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

	newItem, err := userManager.GetCollections().Trips.Collection.Create(&trip.Trip{Title: request.Title})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	updateOptions := phasset.UpdateOptions{
		AssetIds: request.AssetIds,
		AddTrips: []string{newItem.ID},
	}
	_, err = userManager.UpdateAssets(updateOptions)
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

func (handler *TripHandler) Update(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var updateOptions trip.UpdateOptions
	if err := c.ShouldBindJSON(&updateOptions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//collectionManager, err := handler.manager.GetTripManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	item, err := userManager.GetCollections().Trips.Collection.Get(updateOptions.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	trip.Update(item, updateOptions)

	item2, err := userManager.GetCollections().Trips.Collection.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *TripHandler) Delete(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item trip.Trip
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//collectionManager, err := handler.manager.GetTripManager(c, 4)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	err = userManager.GetCollections().Trips.Collection.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "Delete item with id:"+item.ID)
}

func (handler *TripHandler) GetCollectionList(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	items, err := userManager.GetCollections().Trips.Collection.GetAllSorted("creationDate", "1asc")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Create collection list without interface constraint
	result := collection.PHCollectionList[*trip.Trip]{
		Collections: make([]*collection.PHCollection[*trip.Trip], len(items)),
	}

	for i, item := range items {
		assets, _ := userManager.GetCollections().Trips.PhotoAssetList[item.ID]
		result.Collections[i] = &collection.PHCollection[*trip.Trip]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (handler *TripHandler) GetCollectionListWith(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//collectionManager, err := userManager.GetCollections().Trips.GetAll()
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	// Get only visible items
	items, err := userManager.GetCollections().Trips.Collection.GetList(func(a *trip.Trip) bool {
		return !a.IsCollection
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, items)
}
