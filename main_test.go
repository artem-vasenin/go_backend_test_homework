package main

import (
	"os"
	"testing"
)

var (
	requiredFiles = []string{"README.md", "main.go", "main_test.go", "go.mod"}
)

func TestFiles(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Error(err.Error())
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		t.Errorf("Please check permissions for the dir - %s ", dir)
	}
req:
	for _, reqFile := range requiredFiles {
		for _, file := range files {
			if reqFile == file.Name() {
				continue req
			}
		}
		t.Errorf("File doesn't exist - %s", reqFile)
	}
}
