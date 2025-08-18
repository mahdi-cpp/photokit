package main

import (
	"github.com/mahdi-cpp/photokit/internal/api/handler"
	"github.com/mahdi-cpp/photokit/internal/application"
	"log"
)

func main() {

	userStorageManager, err := application.NewApplicationManager()
	if err != nil {
		log.Fatal(err)
	}

	ginInit()

	assetHandler := handler.NewAssetHandler(userStorageManager)
	assetRoute(assetHandler)

	albumHandler := handler.NewAlbumHandler(userStorageManager)
	albumRoute(albumHandler)

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

	api.POST("create", h.Create)
	api.POST("upload", h.Upload)
	api.GET("get:id", h.Get)
	api.POST("update", h.Update)
	api.POST("update_all", h.UpdateAll)
	api.POST("delete", h.Delete)
	api.POST("filters", h.Filters)
}

func albumRoute(h *handler.AlbumHandler) {

	api := router.Group("/api/v1/album")

	api.POST("create", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetListV2)
}

func pinnedRoute(h *handler.PinnedHandler) {

	api := router.Group("/api/v1/pinned")

	api.POST("create", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetList)
}

func sharedAlbumRoute(h *handler.SharedAlbumHandler) {

	api := router.Group("/api/v1/shared_album")

	api.POST("create", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetList)
}

func tripRoute(h *handler.TripHandler) {

	api := router.Group("/api/v1/trip")

	api.POST("create", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("list", h.GetCollectionList)
}

func personRoute(h *handler.PersonHandler) {

	api := router.Group("/api/v1/person")

	api.POST("create", h.Create)
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
