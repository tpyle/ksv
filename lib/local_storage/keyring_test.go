package localstorage_test

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	localstorage "github.com/tpyle/ksv/lib/local_storage"
	"github.com/zalando/go-keyring"
)

func TestKeyringStorage_LoadConfig(t *testing.T) {
	config := map[string]interface{}{
		"storage_size": 10,
	}

	ks := &localstorage.KeyringStorage{}
	err := ks.LoadConfig(config)
	assert.NoError(t, err)
	assert.Equal(t, 10, ks.Config.StorageSize)
}

func TestKeyringStorage_Load(t *testing.T) {
	keyring.MockInit()

	data := strings.Repeat("a", 5000) // larger data to test chunking
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	chunks := localstorage.ChunkString(encodedData, localstorage.ChunkSize)
	for i, chunk := range chunks {
		err := keyring.Set(localstorage.KeyringService, fmt.Sprintf("ksv_chunk_%d", i), chunk)
		assert.NoError(t, err)
	}

	ks := &localstorage.KeyringStorage{}
	loadedData, err := ks.Load()
	assert.NoError(t, err)
	assert.Equal(t, data, string(loadedData))
}

func TestKeyringStorage_Save(t *testing.T) {
	keyring.MockInit()

	data := strings.Repeat("a", 5000) // larger data to test chunking

	ks := &localstorage.KeyringStorage{}
	err := ks.Save([]byte(data))
	assert.NoError(t, err)

	var loadedDataBuilder strings.Builder
	for i := 0; ; i++ {
		chunk, err := keyring.Get(localstorage.KeyringService, fmt.Sprintf("ksv_chunk_%d", i))
		if err != nil {
			if err == keyring.ErrNotFound {
				break
			}
			assert.NoError(t, err)
		}
		loadedDataBuilder.WriteString(chunk)
	}

	decodedData, err := base64.StdEncoding.DecodeString(loadedDataBuilder.String())
	assert.NoError(t, err)
	assert.Equal(t, data, string(decodedData))
}
