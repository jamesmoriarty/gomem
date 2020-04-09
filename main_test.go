package main

import (
	"path/filepath"
	"testing"
	"os"
)

func TestSanity(t *testing.T) {
	main()
}

func TestGetFromProcessName(t *testing.T) {
	name := executableName()

	process, err := GetFromProcessName(name)

	
	if process.ID == 0 {
		t.Errorf("unexpected process id")
	}

	if process.Name != name {
		t.Errorf("unexpected process name")
	}

	if err != nil {
		t.Errorf("unable to get process by name")
	}
}

func executableName() string {
	path, _ := os.Executable()

	return filepath.Base(path)
}
