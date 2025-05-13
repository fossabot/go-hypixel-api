package hypixel

import (
	"testing"
)

func TestParams_String(t *testing.T) {
	params := Params{
		"key1": "value1",
	}

	t.Run("Delete key", func(t *testing.T) {
		params.Del("key1")
		if params.Has("key1") {
			t.Errorf("expected to not have key1")
		}
	})

	t.Run("Set and Get key", func(t *testing.T) {
		params.Set("key2", "value2")
		if params.Get("key2") != "value2" {
			t.Errorf("expected 'value2', got %s", params.Get("key2"))
		}
	})

	t.Run("Generate URL string", func(t *testing.T) {
		url := "https://example.com/v1"
		full := params.String(url)
		expected := "https://example.com/v1?key2=value2"
		if full != expected {
			t.Errorf("expected %s, got %s", expected, full)
		}
	})
}
