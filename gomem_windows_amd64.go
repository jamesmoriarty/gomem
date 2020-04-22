package gomem

import (
	"github.com/jamesmoriarty/gomem/internal/kernel32"
	"github.com/jamesmoriarty/gomem/internal/user32"
)

// Process is a struct representing a windows process.
type Process struct {
	ID     uint32
	Name   string
	Handle uintptr
}

// GetFromProcessName converts a process name to a Process struct.
func GetFromProcessName(name string) (*Process, error) {
	pid, err := kernal32.GetProcessID(name)

	if err != nil {
		return nil, err
	}

	process := Process{ID: pid, Name: name}

	return &process, nil
}

// Open process handle.
func (p *Process) Open() (uintptr, error) {
	handle, err := kernal32.OpenProcess(kernal32.PROCESS_ALL_ACCESS, false, p.ID)

	if err != nil {
		return 0, err
	}

	p.Handle = handle

	return handle, err
}

// Read process memory.
func (p *Process) Read(offset uintptr, buffer *uintptr, length uintptr) error {
	_, err := kernal32.ReadProcessMemory(p.Handle, offset, buffer, length)

	if err != nil {
		return err
	}

	return nil
}

// Write process memory.
func (p *Process) Write(offset uintptr, buffer *uintptr, length uintptr) error {
	_, err := kernal32.WriteProcessMemory(p.Handle, offset, buffer, length)

	if err != nil {
		return err
	}

	return nil
}

// GetModule address.
func (p *Process) GetModule(name string) (uintptr, error) {
	ptr, err := kernal32.GetModule(name, p.ID)

	if err != nil {
		return ptr, err
	}

	return ptr, nil
}

// IsKeyDown https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
func IsKeyDown(v int) bool {
	return user32.GetAsyncKeyState(v) > 0
}
