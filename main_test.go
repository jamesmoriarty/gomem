package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSanity(t *testing.T) {
	main()
}

func TestGetFromProcessName(t *testing.T) {
	name := executableName()

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	if process.ID == 0 {
		t.Errorf("unexpected process id")
	}

	if process.Name != name {
		t.Errorf("unexpected process name")
	}
}

func TestProcessOpen(t *testing.T) {
	name := executableName()

	process, _ := GetFromProcessName(name)

	handle, err := process.Open()

	if err != nil {
		t.Errorf(err.Error())
	}

	if handle == 0 {
		t.Errorf("unable to open handle")
	}
}

func executableName() string {
	path, _ := os.Executable()

	return filepath.Base(path)
}
