package depricated

//func (userStorage *UserManager) UploadAsset(userID int, file multipart.File, header *multipart.FileHeader) (*person_test.PHAsset, error) {
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
//	// Handler person_test filename
//	ext := filepath.Ext(header.Filename)
//	filename := fmt.Sprintf("%d%s", 1, ext)
//	assetPath := filepath.Join(userStorage.config.AssetsDir, filename)
//
//	// Save person_test file
//	if err := os.WriteFile(assetPath, fileBytes, 0644); err != nil {
//		return nil, fmt.Errorf("failed to save person_test: %w", err)
//	}
//
//	// Initialize the ImageExtractor with the path to exiftool_v1
//	extractor := asset_create.NewMetadataExtractor("/usr/local/bin/exiftool_v1")
//
//	// Extract metadata
//	width, height, camera, err := extractor.ExtractMetadata(assetPath)
//	if err != nil {
//		log.Printf("Metadata extraction failed: %v", err)
//	}
//	mediaType := asset_create.GetMediaType(ext)
//
//	// Handler person_test
//	person_test := &person_test.PHAsset{
//		ID:           userStorage.lastID,
//		UserID:       userID,
//		Filename:     filename,
//		CreationDate: time.Now(),
//		MediaType:    mediaType,
//		ImageWidth:   width,
//		PixelHeight:  height,
//		CameraModel:  camera,
//	}
//
//	// Save metadata
//	if err := userStorage.metadata.SaveMetadata(person_test); err != nil {
//		// Clean up person_test file if metadata save fails
//		os.Remove(assetPath)
//		return nil, fmt.Errorf("failed to save metadata: %w", err)
//	}
//
//	// Add to indexes
//	//userStorage.addToIndexes(person_test)
//
//	// Update stats
//	userStorage.statsMu.Lock()
//	userStorage.stats.TotalAssets++
//	userStorage.stats.Uploads24h++
//	userStorage.statsMu.Unlock()
//
//	return person_test, nil
//}
