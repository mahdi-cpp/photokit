package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-pkg/shared_model"
	"github.com/mahdi-cpp/photokit/internal/domain/model"
	"github.com/mahdi-cpp/photokit/internal/storage"
	"net/http"
)

type VillageHandler struct {
	userStorageManager *storage.MainStorageManager
}

func NewVillageHandler(userStorageManager *storage.MainStorageManager) *VillageHandler {
	return &VillageHandler{
		userStorageManager: userStorageManager,
	}
}

func (handler *VillageHandler) GetList(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		c.JSON(400, gin.H{"error": "userID must be an integer"})
		return
	}

	userStorage, err := handler.userStorageManager.GetUserStorage(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	items, err := userStorage.VillageManager.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println("villages: ", len(items))

	result := shared_model.PHCollectionList[*model.Village]{
		Collections: make([]*shared_model.PHCollection[*model.Village], len(items)),
	}

	for i, item := range items {
		result.Collections[i] = &shared_model.PHCollection[*model.Village]{
			Item: item,
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
