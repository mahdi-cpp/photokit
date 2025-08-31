package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/photokit/internal/application"
	asset "github.com/mahdi-cpp/photokit/internal/collections"
	"github.com/mahdi-cpp/photokit/internal/collections/village"
	"github.com/mahdi-cpp/photokit/internal/helpers"
)

type VillageHandler struct {
	manager *application.AppManager
}

func NewVillageHandler(manager *application.AppManager) *VillageHandler {
	return &VillageHandler{
		manager: manager,
	}
}

func (handler *VillageHandler) GetList(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	items, err := userManager.GetCollections().Villages.Collection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println("villages: ", len(items))

	result := asset.PHCollectionList[*village.Village]{
		Collections: make([]*asset.PHCollection[*village.Village], len(items)),
	}

	for i, item := range items {
		result.Collections[i] = &asset.PHCollection[*village.Village]{
			Item: item,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
