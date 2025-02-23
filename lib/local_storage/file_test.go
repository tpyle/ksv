package localstorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileStorage_LoadConfig(t *testing.T) {
	fs := &FileStorage{}
	config := map[string]interface{}{
		"file_path": "testfile.txt",
	}

	err := fs.LoadConfig(config)
	assert.NoError(t, err)
	assert.Equal(t, "testfile.txt", fs.Config.FilePath)
}

func TestFileStorage_LoadConfig_MissingFilePath(t *testing.T) {
	fs := &FileStorage{}
	config := map[string]interface{}{}

	err := fs.LoadConfig(config)
	assert.Error(t, err)
	assert.EqualError(t, err, "file_path is required for file storage")
}

func TestFileStorage_Load(t *testing.T) {
	filePath := "testfile.txt"
	content := []byte("test content")
	os.WriteFile(filePath, content, 0644)
	defer os.Remove(filePath)

	fs := &FileStorage{Config: &FileStorageConfig{FilePath: filePath}}

	data, err := fs.Load()
	assert.NoError(t, err)
	assert.Equal(t, content, data)
}

func TestFileStorage_Save(t *testing.T) {
	filePath := "testfile.txt"
	defer os.Remove(filePath)

	fs := &FileStorage{Config: &FileStorageConfig{FilePath: filePath}}

	content := []byte("test content")
	err := fs.Save(content)
	assert.NoError(t, err)

	result, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, content, result)
}
