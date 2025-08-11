package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-pkg/shared_model"
	"github.com/mahdi-cpp/photokit/internal/domain/model"
	"github.com/mahdi-cpp/photokit/internal/storage"
	"net/http"
)

type CameraHandler struct {
	userStorageManager *storage.MainStorageManager
}

func NewCameraHandler(userStorageManager *storage.MainStorageManager) *CameraHandler {
	return &CameraHandler{
		userStorageManager: userStorageManager,
	}
}

//func (handler *CameraHandler) Create(c *gin.Context) {
//
//	userID, err := getUserId(c)
//	if err != nil {
//		c.JSON(400, gin.H{"error": "userID must be an integer"})
//		return
//	}
//
//	var item model.Camera
//	if err := c.ShouldBindJSON(&item); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//		return
//	}
//
//	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	item2, err := userStorage.CameraManager.Create(&item)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//		return
//	}
//
//	c.JSON(http.StatusCreated, item2)
//}
//
//func (handler *CameraHandler) Update(c *gin.Context) {
//
//	userID, err := getUserId(c)
//	if err != nil {
//		c.JSON(400, gin.H{"error": "userID must be an integer"})
//		return
//	}
//
//	var itemHandler model.CameraHandler
//	if err := c.ShouldBindJSON(&itemHandler); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//		return
//	}
//
//	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	item, err := userStorage.CameraManager.Get(itemHandler.ID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	model.UpdateCamera(item, itemHandler)
//
//	item2, err := userStorage.CameraManager.Update(item)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//		return
//	}
//
//	c.JSON(http.StatusCreated, item2)
//}
//
//func (handler *CameraHandler) Delete(c *gin.Context) {
//
//	userID, err := getUserId(c)
//	if err != nil {
//		c.JSON(400, gin.H{"error": "userID must be an integer"})
//		return
//	}
//
//	var item model.Camera
//	if err := c.ShouldBindJSON(&item); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
//		return
//	}
//
//	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	err = userStorage.CameraManager.Delete(item.ID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//		return
//	}
//
//	c.JSON(http.StatusCreated, "delete ok")
//}

func (handler *CameraHandler) GetList(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//items, err := userStorage.CameraManager.GetAllSorted("creationDate", "a2sc")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	var a []*shared_model.PHCollection[model.Camera]
	result := userStorage.GetAllCameras()
	for _, camera := range result {
		a = append(a, camera)
	}

	//result := model.PHCollectionList[*model.Camera]{
	//	Collections: make([]*model.PHCollection[*model.Camera], len(items)),
	//}

	//for i, item := range items {
	//	//assets, _ := userStorage.CameraManager.GetItemAssets(item.ID)
	//	result.Collections[i] = &model.PHCollection[*model.Camera]{
	//		Item:   item,
	//		Assets: assets,
	//	}
	//}
	cc := gin.H{"collections": a}

	c.JSON(http.StatusOK, gin.H{"data": cc})
}
