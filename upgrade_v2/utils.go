package upgrade

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// readJsonFilesFromDirectory reads all .json file_working from the specified directory
// and unmarshal them into a slice of MyData.
func readJsonFilesFromDirectory(dirPath string) ([]PHAssetV1, error) {

	var assetsV1 []PHAssetV1

	// Read all directory entries
	files, err := os.ReadDir(dirPath) // Use os.ReadDir(dirPath) for Go 1.16+
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip subdirectories
		}

		if strings.HasSuffix(file.Name(), ".json") {
			filePath := filepath.Join(dirPath, file.Name())

			// Read file content
			content, err := os.ReadFile(filePath) // Use os.ReadFile(filePath) for Go 1.16+
			if err != nil {
				return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
			}

			var data PHAssetV1
			// Unmarshal JSON content
			if err := json.Unmarshal(content, &data); err != nil {
				return nil, fmt.Errorf("failed to unmarshal JSON from %s: %w", filePath, err)
			}
			assetsV1 = append(assetsV1, data)
		}
	}
	return assetsV1, nil
}

func CreateDirectory(dirPath string) error {

	// This path includes nested directories that might not exist
	nestedDirPath := dirPath

	// Create the directory path, including any missing parent directories.
	// Permissions are the same as for os.Mkdir.
	err := os.MkdirAll(nestedDirPath, 0755)
	if err != nil {
		log.Fatalf("Error creating nested directories %s: %v\n", nestedDirPath, err)
	}
	fmt.Printf("Nested directories '%s' created successfully (or already existed).\n", nestedDirPath)

	return nil
}

func WriteData[T any](filePath string, data *T) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	tempFile := filePath + ".tmp"
	if err := os.WriteFile(tempFile, jsonData, 0644); err != nil {
		return err
	}

	return os.Rename(tempFile, filePath)
}

// DeleteNestedDirectory attempts to delete the specified directory and all its contents.
func DeleteNestedDirectory(dirPath string) error {
	// os.RemoveAll deletes the path and any children it contains.
	// It does not return an error if the path does not exist.
	return os.RemoveAll(dirPath)
}

func copyMatchingFields(dst, src interface{}) {

	dstVal := reflect.ValueOf(dst).Elem()
	srcVal := reflect.ValueOf(src).Elem()

	// Handle destination pointer allocation
	if dstVal.Kind() == reflect.Ptr && dstVal.IsNil() {
		dstVal.Set(reflect.New(dstVal.Type().Elem()))
	}

	// Get underlying struct values
	var dstStruct, srcStruct reflect.Value

	if dstVal.Kind() == reflect.Ptr {
		dstStruct = dstVal.Elem()
	} else {
		dstStruct = dstVal
	}

	if srcVal.Kind() == reflect.Ptr {
		srcStruct = srcVal.Elem()
	} else {
		srcStruct = srcVal
	}

	// Iterate through source fields
	for i := 0; i < srcStruct.NumField(); i++ {

		srcField := srcStruct.Field(i)
		srcType := srcStruct.Type().Field(i)

		// Skip unexported fields
		if srcType.PkgPath != "" {
			continue
		}

		// Find matching field in destination
		dstField := dstStruct.FieldByName(srcType.Name)
		if !dstField.IsValid() || !dstField.CanSet() {
			continue
		}

		if srcField.Type() == dstField.Type() {
			dstField.Set(srcField)
		}

		//// Handle special conversions
		//switch {
		//// Convert int slices to string slices
		//case (srcType.Name == "Albums" || srcType.Name == "Trips" || srcType.Name == "Persons") &&
		//	srcField.Kind() == reflect.Slice && dstField.Kind() == reflect.Slice:
		//
		//	if srcField.Len() > 0 && srcField.Index(0).Kind() == reflect.Int {
		//		for j := 0; j < srcField.Len(); j++ {
		//			intVal := srcField.Index(j).Int()
		//			strVal := strconv.Itoa(int(intVal))
		//			dstField.Set(reflect.Append(dstField, reflect.ValueOf(strVal)))
		//		}
		//	}
		//
		//// Skip ID field (we'll handle separately)
		//case srcType.Name == "ID":
		//	continue
		//
		//// Direct copy for identical types
		//case srcField.Type() == dstField.Type():
		//	dstField.Set(srcField)
		//
		//// Convert int to string for UserID
		//case srcType.Name == "UserID" && dstField.Kind() == reflect.String:
		//	if srcField.Kind() == reflect.Int {
		//		dstField.SetString(strconv.Itoa(int(srcField.Int())))
		//	}
		//}
	}
}

// ============== HELPER FUNCTIONS ==============
func backupFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	backupPath := fmt.Sprintf("%s.%s.bak", path, time.Now().Format("20060102-150405"))
	return os.WriteFile(backupPath, data, 0644)
}

func intSliceToStrSlice(nums []int) []string {
	strs := make([]string, len(nums))
	for i, n := range nums {
		strs[i] = strconv.Itoa(n)
	}
	return strs
}

func writeFile(path string, data []byte) error {
	tmpFile, err := os.CreateTemp("", "upgrade_v2-*.json")
	if err != nil {
		return err
	}
	//defer os.Remove(tmpFile.Name())

	defer func() {
		// Explicitly ignore the error returned by os.Remove
		_ = os.Remove(tmpFile.Name())
	}()

	if _, err := tmpFile.Write(data); err != nil {
		err := tmpFile.Close()
		if err != nil {
			return err
		}
	}

	if err := tmpFile.Close(); err != nil {
		return err
	}

	return os.Rename(tmpFile.Name(), path)
}

func RenameFile(oldName string, newName string) error {

	// Rename the file
	err := os.Rename(oldName, newName)
	if err != nil {
		fmt.Printf("Error renaming file: %v\n", err)
		return err
	}
	fmt.Printf("Renamed %s to %s successfully!\n", oldName, newName)

	// Verify the new file exists and the old one doesn't
	if _, err = os.Stat(newName); os.IsNotExist(err) {
		fmt.Printf("Error: %s does not exist after rename.\n", newName)
	} else {
		fmt.Printf("Verified: %s exists.\n", newName)
	}

	if _, err = os.Stat(oldName); !os.IsNotExist(err) {
		fmt.Printf("Error: %s still exists after rename.\n", oldName)
	} else {
		fmt.Printf("Verified: %s no longer exists.\n", oldName)
	}

	return nil
}

// RenameDirectory renames a directory from oldPath to newPath.
// It returns an error if the operation fails.
func RenameDirectory(oldPath, newPath string) error {
	// os.Rename can be used for both files and directories.
	// If newPath already exists and is a directory, the behavior depends on the OS:
	// - On Unix-like systems, if newPath is an empty directory, it's replaced.
	//   If newPath is a non-empty directory, os.Rename will typically fail with an "directory not empty" error.
	// - On Windows, if newPath exists, it must be an empty directory for the rename to succeed.
	// It's generally safer to ensure newPath does not exist, or handle potential conflicts.
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return fmt.Errorf("failed to rename directory from '%s' to '%s': %w", oldPath, newPath, err)
	}
	//fmt.Printf("Successfully renamed directory from '%s' to '%s'\n", oldPath, newPath)
	return nil
}

// IsDirectoryExist checks if a path exists and is a directory.
// It returns true if the path exists and is a directory, false otherwise.
func IsDirectoryExist(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		// The path exists, now check if it's a directory.
		return info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		// The path does not exist.
		return false, nil
	}
	// For any other error (e.g., permission issues), return the error.
	return false, err
}

// contains checks if an element exists in a slice.
// This is a common helper if not using Go 1.21+ 'slices' package.
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
