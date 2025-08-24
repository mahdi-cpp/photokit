package main

import (
	"fmt"
	"github.com/mahdi-cpp/photokit/internal/api/handler"
	"github.com/mahdi-cpp/photokit/internal/application"
	"github.com/mahdi-cpp/photokit/upgrade_v3"
	"log"
	"time"
)

func main() {

	userStorageManager, err := application.NewApplicationManager()
	if err != nil {
		log.Fatal(err)
	}

	// Wait for initial user list with 10 second timeout
	if err := userStorageManager.WaitForInitialUserList(10 * time.Second); err != nil {
		log.Printf("Warning: %v", err)
		// You might choose to continue or exit based on your requirements
	}

	fmt.Println("execute after get users ---------------------------")

	//if !utils.CheckVersionIsUpToDate(2) {
	//upgrade.Start(userStorageManager.AccountManager)
	upgrade_v3.Start(userStorageManager.AccountManager)
	//}

	ginInit()

	assetHandler := handler.NewAssetHandler(userStorageManager)
	assetRoute(assetHandler)

	albumHandler := handler.NewAlbumHandler(userStorageManager)
	RegisterAlbumRoutes(albumHandler)

	tripHandler := handler.NewTripHandler(userStorageManager)
	tripRoute(tripHandler)

	sharedAlbumHandler := handler.NewSharedAlbumHandler(userStorageManager)
	sharedAlbumRoute(sharedAlbumHandler)

	personHandler := handler.NewPersonsHandler(userStorageManager)
	personRoute(personHandler)

	villageHandler := handler.NewVillageHandler(userStorageManager)
	villageRoute(villageHandler)

	pinnedHandler := handler.NewPinnedHandler(userStorageManager)
	pinnedRoute(pinnedHandler)

	cameraHandler := handler.NewCameraHandler(userStorageManager)
	cameraRoute(cameraHandler)

	startServer(router)
}

func assetRoute(h *handler.AssetHandler) {

	api := router.Group("/api/v1/assets")

	api.POST("thumbnail", h.Create)
	api.POST("upload", h.Upload)
	api.GET("get:id", h.Get)
	api.POST("update", h.Update)
	api.POST("update_all", h.UpdateAll)
	api.POST("delete", h.Delete)
	api.POST("filters", h.Filters)
}

func RegisterAlbumRoutes(h *handler.AlbumHandler) {

	api := router.Group("/api/v1/album")

	api.POST("thumbnail", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetAll)
	api.POST("search", h.GetBySearchOptions)
}

func pinnedRoute(h *handler.PinnedHandler) {

	api := router.Group("/api/v1/pinned")

	api.POST("thumbnail", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetList)
}

func sharedAlbumRoute(h *handler.SharedAlbumHandler) {

	api := router.Group("/api/v1/shared_album")

	api.POST("thumbnail", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetList)
}

func tripRoute(h *handler.TripHandler) {

	api := router.Group("/api/v1/trip")

	api.POST("thumbnail", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetCollectionList)
}

func personRoute(h *handler.PersonHandler) {

	api := router.Group("/api/v1/person")

	api.POST("thumbnail", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetCollectionList)
}

func cameraRoute(h *handler.CameraHandler) {

	api := router.Group("/api/v1/camera")

	api.POST("/list", h.GetList)
}

func villageRoute(h *handler.VillageHandler) {

	api := router.Group("/api/v1/village")

	api.POST("list", h.GetList)
}
