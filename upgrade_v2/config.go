package upgrade

//var usersDir = "/app/iris/com.iris.photos/users"
//var assetsDir = "/app/iris/com.iris.photos/users/018f3a8b-1b32-729a-f7e5-5467c1b2d3e4/assets"
//var thumbnailsDir = "/app/iris/com.iris.photos/users/018f3a8b-1b32-729a-f7e5-5467c1b2d3e4/thumbnails"
//var metadataDir = "/app/iris/com.iris.photos/users/018f3a8b-1b32-729a-f7e5-5467c1b2d3e4/metadata"

const (
	// The user ID is a constant
	userID = "018f3a8b-1b32-729a-f7e5-5467c1b2d3e4"

	// Define all directories as constants using filepath.Join
	// Note: You can't join paths with `const` directly in Go,
	// so you need to do this at the variable declaration level.
	usersDir = "/app/iris/com.iris.photos/users"

	currentUserPath = usersDir + "/" + userID // A good intermediate constant
	assetsDir       = currentUserPath + "/assets"
	thumbnailsDir   = currentUserPath + "/thumbnails"
	metadataDir     = currentUserPath + "/metadata"
)

const (
	currentVersion = "v1"
	newVersion     = "v3"
)
