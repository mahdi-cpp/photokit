package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mahdi-cpp/photokit/internal/application"
	"github.com/mahdi-cpp/photokit/internal/collections/phasset"
	"github.com/mahdi-cpp/photokit/internal/helpers"
)

type AssetHandler struct {
	manager *application.AppManager
}

func NewAssetHandler(manager *application.AppManager) *AssetHandler {
	return &AssetHandler{manager: manager}
}

func (handler *AssetHandler) Create(c *gin.Context) {
}

func (handler *AssetHandler) Upload(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
		return
	}
	// Recommended: Log the error if the file can't be closed properly.
	defer func() {
		if err := file.Close(); err != nil {
			// Here, you should log the error.
			// For example, using the standard "log" package:
			log.Printf("error closing file: %v", err)
		}
	}()

	// Handler person_test metadata
	asset := &phasset.PHAsset{
		UserID: userID,
		FileInfo: phasset.FileInfo{
			FileType: header.Filename,
		},
	}

	//userManager, err := handler.manager.GetUserManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//}

	//person_test, err = userManager.UploadAsset(person_test.UserID, file, header)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Processing failed"})
	//	return
	//}

	//person_test, err := handler.manager.Upload(c, userID, file, header)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Processing failed"})
	//	return
	//}

	c.JSON(http.StatusCreated, asset)
}

func (handler *AssetHandler) Update(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var updateOptions phasset.UpdateOptions
	if err := c.ShouldBindJSON(&updateOptions); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	asset, err := userManager.UpdateAssets(updateOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userManager.UpdateCollections()

	c.JSON(http.StatusCreated, asset)
}

func (handler *AssetHandler) UpdateAll(c *gin.Context) {
	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	var updateOptions phasset.UpdateOptions
	if err := c.ShouldBindJSON(&updateOptions); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	allAssets, err := userManager.GetAllAssets()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, asset := range allAssets {
		updateOptions.AssetIds = append(updateOptions.AssetIds, asset.ID)
	}

	asset, err := userManager.UpdateAssets(updateOptions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userManager.UpdateCollections()

	c.JSON(http.StatusCreated, asset)
}

func (handler *AssetHandler) Get(c *gin.Context) {

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

	asset, err := userManager.GetCollections().Assets.Get(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Asset not found"})
		return
	}

	c.JSON(http.StatusOK, asset)
}

func (handler *AssetHandler) Search(c *gin.Context) {

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	query := c.Query("query")
	//mediaType := c.Query("type")

	var dateRange []time.Time
	if start := c.Query("start"); start != "" {
		if t, err := time.Parse("2006-01-02", start); err == nil {
			dateRange = append(dateRange, t)
		}
	}
	if end := c.Query("end"); end != "" {
		if t, err := time.Parse("2006-01-02", end); err == nil {
			dateRange = append(dateRange, t)
		}
	}

	filters := phasset.SearchOptions{
		UserID:    userID,
		TextQuery: query,
		//MediaType: phasset.MediaType(mediaType),
	}

	if len(dateRange) > 0 {
		filters.CreatedAfter = &dateRange[0]
	}
	if len(dateRange) > 1 {
		filters.CreatedBefore = &dateRange[1]
	}

	//person_test, _, err := s.repo.Search(ctx, filters)
	//return person_test, err

	//person_test, _, err := handler.manager.Search(c, filters)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
	//	return
	//}

	//c.JSON(http.StatusOK, person_test)
}

func (handler *AssetHandler) Delete(c *gin.Context) {

	//userID := c.Query("userID")
	//userID, err := strconv.Atoi(userID)
	//if err != nil {
	//	c.JSON(400, gin.H{"error": "userID must be an integer"})
	//	return
	//}

	//var request phasset.Delete
	//if err := c.ShouldBindJSON(&request); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	return
	//}
	//
	//userManager, err := handler.manager.GetUserManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//}
	//
	//err = userManager.DeleteAsset(request.AssetID)
	//if err != nil {
	//	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, "successful delete person_test with id: "+request.AssetID)
}

func (handler *AssetHandler) Filters(c *gin.Context) {

	fmt.Println(c.GetHeader("X-User-ID"))

	userID, ok := helpers.GetUserID(c)
	if !ok {
		helpers.AbortWithUserIDInvalid(c)
		return
	}

	fmt.Println("12: ", userID)

	var with *phasset.SearchOptions
	if err := c.ShouldBindJSON(&with); err != nil {
		helpers.AbortWithRequestInvalid(c)
		return
	}

	userManager, err := handler.manager.GetUserManager(c, userID)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items, err := userManager.FetchAssets(with)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed user FetchAssets"})
		return
	}

	fmt.Println("Filters count", len(items))

	//result := asset.PHFetchResult[*phasset.PHAsset]{
	//	Items:  items,
	//	Total:  total,
	//	Limit:  100,
	//	Offset: 100,
	//}
	//c.JSON(http.StatusOK, items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}
