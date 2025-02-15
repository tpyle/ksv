package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestKSV_Validate(t *testing.T) {
	validSite := Site{
		Url:  "https://example.com",
		Name: "Example",
	}
	invalidSite := Site{}

	ksv := KSV{
		DefaultNamespace: Namespace{
			Sites: []Site{
				validSite,
				invalidSite,
			},
		},
	}

	errs := ksv.Validate()
	if len(errs) == 0 {
		t.Errorf("Expected validation errors, got none")
	}
}

func TestKSV_MarshalUnmarshalJSON(t *testing.T) {
	ksv := KSV{
		Namespaces: map[string][]byte{
			"namespace1": []byte("data1"),
			"namespace2": []byte("data2"),
		},
		DefaultNamespace: Namespace{},
	}

	data, err := json.Marshal(ksv)
	if err != nil {
		t.Fatalf("MarshalJSON failed: %v", err)
	}

	var unmarshaledKSV KSV
	if err := json.Unmarshal(data, &unmarshaledKSV); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if !reflect.DeepEqual(ksv, unmarshaledKSV) {
		t.Errorf("Expected %v, got %v", ksv, unmarshaledKSV)
	}
}
