```file.go
AlbumManager, err = collection.NewCollectionManager[*model.Album]("albums.json", false)
AlbumManager, err = collection.NewCollectionManager[*model.Album]("albums_directory", false)

albums_directory
----1.json
----2.json
----3.json
```

albums_directory/
├── 1/
│ └── metadata.json
├── 2/
│ └── metadata.json
└── 3/
└── metadata.json

albums_directory/
├── 1.json
├── 2.json
└── 3.json