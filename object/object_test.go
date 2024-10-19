package object

import (
	"slices"
	"testing"
)

func TestGetKeys(t *testing.T) {
	data := make(map[string]string, 0)
	data["Name"] = "Tom"
	data["Sex"] = "Boy"

	rst := GetKeys(data)
	expected := []string{"Name", "Sex"}

	if !slices.Equal(rst, expected) {
		t.Fatalf("expect %v, but got %v", expected, rst)
	}
}

func TestGetValues(t *testing.T) {
	data := make(map[string]string, 0)
	data["Name"] = "Tom"
	data["Sex"] = "Boy"

	rst := GetValues(data)
	expected := []string{"Tom", "Boy"}

	if !slices.Equal(rst, expected) {
		t.Fatalf("expect %v, but got %v", expected, rst)
	}
}
