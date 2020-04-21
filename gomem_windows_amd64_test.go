package gomem

import (
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

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
		t.Errorf("unexpected handle id")
	}
}

func TestProcessRead(t *testing.T) {
	name := executableName()

	var buffer uintptr
	bufferPtr := &buffer

	value := 42
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	process.Open()
	err = process.Read(valuePtr, bufferPtr, unsafe.Sizeof(value))

	if err != nil {
		t.Errorf(err.Error())
	}

	if (int)(*bufferPtr) != value {
		t.Errorf("unexpected value")
	}
}

func TestProcessWrite(t *testing.T) {
	name := executableName()

	buffer := (uintptr)(43)
	bufferPtr := &buffer

	value := 42
	valuePtr := (uintptr)(unsafe.Pointer(&value))

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	process.Open()
	err = process.Write(valuePtr, bufferPtr, unsafe.Sizeof(value))

	if err != nil {
		t.Errorf(err.Error())
	}

	if (int)(buffer) != value {
		t.Errorf("unexpected value")
	}
}

func TestGetModuleNotFound(t *testing.T) {
	name := executableName()

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	ptr, err := process.GetModule("client.dll")

	if err.Error() != "not found" {
		t.Errorf(err.Error())
	}

	if (ptr) == 0 {
		t.Errorf("unexpected value")
	}
}

func TestIsKeyDown(t *testing.T) { 
	value := IsKeyDown(0x20) // https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes

	if value != false {
		t.Errorf("unexpected value")
	}
}

func executableName() string {
	path, _ := os.Executable()

	return filepath.Base(path)
}
