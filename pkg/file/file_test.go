package file

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	type TestStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	tests := []struct {
		name      string
		filepath  string
		want      *TestStruct
		expectErr bool
	}{
		{
			name:     "Valid JSON file",
			filepath: "testdata/valid.json",
			want: &TestStruct{
				Name: "John",
				Age:  30,
			},
			expectErr: false,
		},
		{
			name:      "Non-existent file",
			filepath:  "testdata/nonexistent.json",
			want:      nil,
			expectErr: true,
		},
		{
			name:      "Invalid JSON file",
			filepath:  "testdata/invalid.json",
			want:      nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadJSONFile[TestStruct](tt.filepath)

			if (!tt.expectErr && err != nil) || (tt.expectErr && err == nil) {
				t.Errorf("ReadFile() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && *got != *tt.want {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func setup() {
	os.MkdirAll("testdata", os.ModePerm)
	os.WriteFile("testdata/valid.json", []byte(`{"name": "John", "age": 30}`), os.ModePerm)
	os.WriteFile("testdata/invalid.json", []byte(`{"name": "John", "age":`), os.ModePerm)
}

func teardown() {
	os.RemoveAll("testdata")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
