package main

import (
	"github.com/mahdi-cpp/photokit/internal/api/handler"
	"github.com/mahdi-cpp/photokit/internal/storage"
	"log"
)

func main() {

	userStorageManager, err := storage.NewMainStorageManager()
	if err != nil {
		log.Fatal(err)
	}

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

	searchHandler := handler.NewSearchHandler(userStorageManager)
	searchRoute(searchHandler)

	pinnedHandler := handler.NewPinnedHandler(userStorageManager)
	pinnedRoute(pinnedHandler)

	cameraHandler := handler.NewCameraHandler(userStorageManager)
	cameraRoute(cameraHandler)

	startServer(router)
}

func assetRoute(assetHandler *handler.AssetHandler) {

	api := router.Group("/api/v1/assets")

	// Asset routes
	api.POST("create", assetHandler.Create)
	api.POST("upload", assetHandler.Upload)
	api.GET("get:id", assetHandler.Get)
	api.POST("update", assetHandler.Update)
	api.POST("update_all", assetHandler.UpdateAll)
	api.POST("delete", assetHandler.Delete)
	api.POST("filters", assetHandler.Filters)

	//http://localhost:8080/api/v1/assets/download/thumbnail/map_270.jpg
	api.GET("download/:filename", assetHandler.OriginalDownload)
	api.GET("download/thumbnail/:filename", assetHandler.TinyImageDownload)
	api.GET("download/icons/:filename", assetHandler.IconDownload)
}

func albumRoute(albumHandler *handler.AlbumHandler) {

	api := router.Group("/api/v1/album")

	api.POST("create", albumHandler.Create)
	api.POST("update", albumHandler.Update)
	api.POST("delete", albumHandler.Delete)
	api.POST("list", albumHandler.GetListV2)
}

func pinnedRoute(pinnedHandler *handler.PinnedHandler) {

	api := router.Group("/api/v1/pinned")

	api.POST("create", pinnedHandler.Create)
	api.POST("update", pinnedHandler.Update)
	api.POST("delete", pinnedHandler.Delete)
	api.POST("list", pinnedHandler.GetList)
}

func sharedAlbumRoute(sharedAlbumHandler *handler.SharedAlbumHandler) {

	api := router.Group("/api/v1/shared_album")

	api.POST("create", sharedAlbumHandler.Create)
	api.POST("update", sharedAlbumHandler.Update)
	api.POST("delete", sharedAlbumHandler.Delete)
	api.POST("list", sharedAlbumHandler.GetList)
}

func tripRoute(tripHandler *handler.TripHandler) {

	api := router.Group("/api/v1/trip")

	api.POST("create", tripHandler.Create)
	api.POST("update", tripHandler.Update)
	api.POST("delete", tripHandler.Delete)
	api.POST("list", tripHandler.GetCollectionList)
}

func personRoute(personHandler *handler.PersonHandler) {

	api := router.Group("/api/v1/person")

	api.POST("create", personHandler.Create)
	api.POST("update", personHandler.Update)
	api.POST("delete", personHandler.Delete)
	api.POST("list", personHandler.GetCollectionList)
}

func cameraRoute(cameraHandler *handler.CameraHandler) {

	api := router.Group("/api/v1/camera")

	api.POST("/list", cameraHandler.GetList)
}

func searchRoute(searchHandler *handler.SearchHandler) {
	api := router.Group("/api/v1/search")

	api.GET("/", searchHandler.Search)
	api.POST("/filters", searchHandler.Filters)
}

func villageRoute(villageHandler *handler.VillageHandler) {

	api := router.Group("/api/v1/village")

	api.POST("list", villageHandler.GetList)
}
