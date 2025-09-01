# PhotoKit

### Build Your Project From Its Directory

```
/opt/github.com/mahdi-cpp/PhotoKit/cmd/server
```

### generic search

https://chat.deepseek.com/a/chat/s/456940d9-d8e5-42e4-9ffa-ca1c01128744

## git

```
git remote set-url origin https://github.com/mahdi-cpp/PhotoKit.git
```


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


broker-service/
├── cmd/
│ └── main.go
├── internal/
│ └── directory1
│ └── directory2
│ └── directory3