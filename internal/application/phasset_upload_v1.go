package application

//func (userStorage *UserManager) UploadAsset(userID int, file multipart.File, header *multipart.FileHeader) (*build_asset.PHAsset, error) {
//
//	// Check file size
//	//if header.Size > userStorage.config.MaxUploadSize {
//	//	return nil, ErrFileTooLarge
//	//}
//
//	// Read file content
//	fileBytes, err := io.ReadAll(file)
//	if err != nil {
//		return nil, fmt.Errorf("failed to read file: %w", err)
//	}
//
//	// Handler build_asset filename
//	ext := filepath.Ext(header.Filename)
//	filename := fmt.Sprintf("%d%s", 1, ext)
//	assetPath := filepath.Join(userStorage.config.AssetsDir, filename)
//
//	// Save build_asset file
//	if err := os.WriteFile(assetPath, fileBytes, 0644); err != nil {
//		return nil, fmt.Errorf("failed to save build_asset: %w", err)
//	}
//
//	// Initialize the ImageExtractor with the path to exiftool
//	extractor := asset_create.NewMetadataExtractor("/usr/local/bin/exiftool")
//
//	// Extract metadata
//	width, height, camera, err := extractor.ExtractMetadata(assetPath)
//	if err != nil {
//		log.Printf("Metadata extraction failed: %v", err)
//	}
//	mediaType := asset_create.GetMediaType(ext)
//
//	// Handler build_asset
//	build_asset := &build_asset.PHAsset{
//		ID:           userStorage.lastID,
//		UserID:       userID,
//		Filename:     filename,
//		CreationDate: time.Now(),
//		MediaType:    mediaType,
//		PixelWidth:   width,
//		PixelHeight:  height,
//		CameraModel:  camera,
//	}
//
//	// Save metadata
//	if err := userStorage.metadata.SaveMetadata(build_asset); err != nil {
//		// Clean up build_asset file if metadata save fails
//		os.Remove(assetPath)
//		return nil, fmt.Errorf("failed to save metadata: %w", err)
//	}
//
//	// Add to indexes
//	//userStorage.addToIndexes(build_asset)
//
//	// Update stats
//	userStorage.statsMu.Lock()
//	userStorage.stats.TotalAssets++
//	userStorage.stats.Uploads24h++
//	userStorage.statsMu.Unlock()
//
//	return build_asset, nil
//}
