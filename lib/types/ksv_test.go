package types

import (
	"encoding/base64"
	"encoding/json"
	"reflect"
	"testing"
)

func TestKSV_MarshalJSON(t *testing.T) {
	ksv := &KSV{
		Namespaces: map[string][]byte{
			"namespace1": []byte("value1"),
		},
		DefaultNamespace: Namespace{},
	}

	data, err := ksv.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON() error = %v", err)
	}

	var aux struct {
		Namespaces       map[string]string `json:"namespaces"`
		DefaultNamespace Namespace         `json:"default_namespace"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	if !reflect.DeepEqual(aux.DefaultNamespace, ksv.DefaultNamespace) {
		t.Errorf("DefaultNamespace = %v, want %v", aux.DefaultNamespace, ksv.DefaultNamespace)
	}

	for key, val := range aux.Namespaces {
		decodedVal, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			t.Fatalf("DecodeString() error = %v", err)
		}
		if !reflect.DeepEqual(decodedVal, ksv.Namespaces[key]) {
			t.Errorf("Namespaces[%v] = %v, want %v", key, decodedVal, ksv.Namespaces[key])
		}
	}
}

func TestKSV_UnmarshalJSON(t *testing.T) {
	data := []byte(`{
		"namespaces": {
			"namespace1": "dmFsdWUx"
		},
		"default_namespace": {
			"name": "default"
		}
	}`)

	ksv := &KSV{}
	if err := ksv.UnmarshalJSON(data); err != nil {
		t.Fatalf("UnmarshalJSON() error = %v", err)
	}

	want := &KSV{
		Namespaces: map[string][]byte{
			"namespace1": []byte("value1"),
		},
		DefaultNamespace: Namespace{},
	}

	if !reflect.DeepEqual(ksv, want) {
		t.Errorf("UnmarshalJSON() = %v, want %v", ksv, want)
	}
}
