package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/mahdi-cpp/photokit/internal/application"
	asset "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/camera"
	"net/http"
)

type CameraHandler struct {
	manager *application.AppManager
}

func NewCameraHandler(manager *application.AppManager) *CameraHandler {
	return &CameraHandler{
		manager: manager,
	}
}

//func (handler *CameraHandler) Create(c *gin.Context) {
//
//	userID, err := account.GetUserId(c)
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
//	userManager, err := handler.manager.GetUserManager(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	item2, err := userManager.CameraManager.Create(&item)
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
//	userID, err := account.GetUserId(c)
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
//	userManager, err := handler.manager.GetUserManager(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	item, err := userManager.CameraManager.Get(itemHandler.ID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	model.UpdateCamera(item, itemHandler)
//
//	item2, err := userManager.CameraManager.Update(item)
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
//	userID, err := account.GetUserId(c)
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
//	userManager, err := handler.manager.GetUserManager(c, userID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//	}
//
//	err = userManager.CameraManager.Delete(item.ID)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//		return
//	}
//
//	c.JSON(http.StatusCreated, "delete ok")
//}

func (handler *CameraHandler) GetList(c *gin.Context) {

	userID, err := account.GetUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//items, err := userManager.CameraManager.GetAllSorted("creationDate", "a2sc")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	var a []*asset.PHCollection[camera.Camera]
	result := userManager.GetAllCameras()
	for _, camera := range result {
		a = append(a, camera)
	}

	//result := model.PHCollectionList[*model.Camera]{
	//	collection: make([]*model.PHCollection[*model.Camera], len(items)),
	//}

	//for i, item := range items {
	//	//build_asset, _ := userManager.CameraManager.GetItemAssets(item.ID)
	//	result.collection[i] = &model.PHCollection[*model.Camera]{
	//		Item:   item,
	//		Assets: build_asset,
	//	}
	//}
	cc := gin.H{"collections": a}

	c.JSON(http.StatusOK, gin.H{"data": cc})
}
