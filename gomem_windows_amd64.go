package gomem

import (
	"github.com/jamesmoriarty/gomem/internal/sys"
)

// Process is a struct representing a windows process.
type Process struct {
	ID     uint32
	Name   string
	Handle uintptr
}

// GetFromProcessName converts a process name to a Process struct.
func GetFromProcessName(name string) (*Process, error) {
	pid, err := sys.GetProcessID(name)

	if err != nil {
		return nil, err
	}

	process := Process{ID: pid, Name: name}

	return &process, nil
}

// Open process handle.
func (p *Process) Open() (uintptr, error) {
	handle, err := sys.OpenProcess(sys.PROCESS_ALL_ACCESS, false, p.ID)

	if err != nil {
		return 0, err
	}

	p.Handle = handle

	return handle, err
}

// Read process memory.
func (p *Process) Read(offset uintptr, buffer *uintptr, length uintptr) error {
	_, err := sys.ReadProcessMemory(p.Handle, offset, buffer, length)

	if err != nil {
		return err
	}

	return nil
}

// Write process memory.
func (p *Process) Write(offset uintptr, buffer *uintptr, length uintptr) error {
	_, err := sys.WriteProcessMemory(p.Handle, offset, buffer, length)

	if err != nil {
		return err
	}

	return nil
}

// GetModule address.
func (p *Process) GetModule(name string) (uintptr, error) {
	ptr, err := sys.GetModule(name, p.ID)

	if err != nil {
		return ptr, err
	}

	return ptr, nil
}
