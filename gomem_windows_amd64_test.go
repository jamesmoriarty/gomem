package gomem

import (
	"os"
	"path/filepath"
	"testing"
	"runtime"
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

	var bufferValue int
	bufferPtr := (uintptr)(unsafe.Pointer(&bufferValue))

	offsetValue := 42
	offsetPtr := (uintptr)(unsafe.Pointer(&offsetValue))

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	process.Open()
	err = process.Read(offsetPtr, bufferPtr, unsafe.Sizeof(offsetValue))

	if err != nil {
		t.Errorf(err.Error())
	}

	if (int)(bufferValue) != 42 {
		t.Errorf("unexpected value")
	}
}

func TestProcessWrite(t *testing.T) {
	name := executableName()

	var bufferValue = 43
	bufferPtr := (uintptr)(unsafe.Pointer(&bufferValue))

	offsetValue := 42
	offsetPtr := (uintptr)(unsafe.Pointer(&offsetValue))

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	process.Open()
	err = process.Write(offsetPtr, bufferPtr, unsafe.Sizeof(bufferValue))

	if err != nil {
		t.Errorf(err.Error())
	}

	if (int)(offsetValue) != 43 {
		t.Errorf("unexpected value")
	}

	runtime.KeepAlive(&bufferValue)
	runtime.KeepAlive(&offsetValue)
}

func TestGetModuleNotFound(t *testing.T) {
	name := executableName()

	process, err := GetFromProcessName(name)

	if err != nil {
		t.Errorf(err.Error())
	}

	process.Open()
	ptr, err := process.GetModule("unknown.dll")

	if err.Error() != "module not found" {
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
