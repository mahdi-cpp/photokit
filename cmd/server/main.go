package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mahdi-cpp/photokit/internal/api/handler"
	"github.com/mahdi-cpp/photokit/internal/application"
)

func main() {

	applicationManager, err := application.NewApplicationManager()
	if err != nil {
		log.Fatal(err)
	}

	// Wait for initial user list with 10 second timeout
	if err := applicationManager.WaitForInitialUserList(10 * time.Second); err != nil {
		log.Printf("Warning: %v", err)
		// You might choose to continue or exit based on your requirements
	}

	fmt.Println("execute after get users ---------------------------")

	//if !utils.CheckVersionIsUpToDate(2) {
	//upgrade.Start(applicationManager.AccountManager)
	//upgrade_v3.Start(applicationManager.AccountManager)
	//}

	ginInit()

	assetHandler := handler.NewAssetHandler(applicationManager)
	assetRoute(assetHandler)

	albumHandler := handler.NewAlbumHandler(applicationManager)
	RegisterAlbumRoutes(albumHandler)

	tripHandler := handler.NewTripHandler(applicationManager)
	tripRoute(tripHandler)

	sharedAlbumHandler := handler.NewSharedAlbumHandler(applicationManager)
	sharedAlbumRoute(sharedAlbumHandler)

	personHandler := handler.NewPersonsHandler(applicationManager)
	personRoute(personHandler)

	villageHandler := handler.NewVillageHandler(applicationManager)
	villageRoute(villageHandler)

	pinnedHandler := handler.NewPinnedHandler(applicationManager)
	pinnedRoute(pinnedHandler)

	cameraHandler := handler.NewCameraHandler(applicationManager)
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
