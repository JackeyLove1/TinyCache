package TinyCache

import (
	"testing"
)

func TestGroup_Do(t *testing.T) {
	var g Group
	v, err := g.Do("Key", func() (interface{}, error) {
		return "bar", nil
	})

	if v != "bar" || err != nil {
		t.Errorf("Do v = %v, error = %v\n", v, err)
	}
}
