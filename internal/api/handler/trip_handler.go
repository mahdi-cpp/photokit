package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-pkg/shared_model"
	"github.com/mahdi-cpp/photokit/internal/domain/model"
	"github.com/mahdi-cpp/photokit/internal/storage"
	"net/http"
	"strconv"
)

type TripHandler struct {
	userStorageManager *storage.MainStorageManager
}

func NewTripHandler(userStorageManager *storage.MainStorageManager) *TripHandler {
	return &TripHandler{
		userStorageManager: userStorageManager,
	}
}

func (handler *TripHandler) Create(c *gin.Context) {

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

	newItem, err := userStorage.TripManager.Create(&model.Trip{Title: request.Title})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	update := shared_model.AssetUpdate{AssetIds: request.AssetIds, AddTrips: []int{newItem.ID}}
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

func (handler *TripHandler) Update(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var itemHandler model.TripHandler
	if err := c.ShouldBindJSON(&itemHandler); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//collectionManager, err := handler.userStorageManager.GetTripManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	item, err := userStorage.TripManager.Get(itemHandler.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	model.UpdateTrip(item, itemHandler)

	item2, err := userStorage.TripManager.Update(item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *TripHandler) Delete(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	var item model.Trip
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//collectionManager, err := handler.userStorageManager.GetTripManager(c, 4)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	err = userStorage.TripManager.Delete(item.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, "Delete item with id:"+strconv.Itoa(item.ID))
}

func (handler *TripHandler) GetCollectionList(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	items, err := userStorage.TripManager.GetAllSorted("creationDate", "1asc")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Create collection list without interface constraint
	result := shared_model.PHCollectionList[*model.Trip]{
		Collections: make([]*shared_model.PHCollection[*model.Trip], len(items)),
	}

	for i, item := range items {
		assets, _ := userStorage.TripManager.GetItemAssets(item.ID)
		result.Collections[i] = &shared_model.PHCollection[*model.Trip]{
			Item:   item,
			Assets: assets,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (handler *TripHandler) GetCollectionListWith(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//collectionManager, err := userStorage.TripManager.GetAll()
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	// Get only visible items
	items, err := userStorage.TripManager.GetList(func(a *model.Trip) bool {
		return !a.IsCollection
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, items)
}
