package main

import (
	"unsafe"
	"github.com/jamesmoriarty/gomem/internal/process"
)

// Process is a struct representing a windows process.
type Process struct {
	ID     uint32
	Name   string
	Handle uintptr
}

// GetFromProcessName converts a process name to a Process struct.
func GetFromProcessName(name string) (*Process, error) {
	pid, err := process.GetProcessID(name)

	if err != nil {
		return nil, err
	}

	process := Process{ID: pid, Name: name}

	return &process, nil
}

// Open opens a Process handle for read/write.
// It returns uintptr representing a windows handle.
func (p *Process) Open() (uintptr, error) {
	handle, err := process.OpenProcess(process.PROCESS_ALL_ACCESS, false, p.ID)
	
	if err != nil {
		return 0, err
	}

	p.Handle = handle

	return handle, err
}

// Read process memory.
func (p *Process) Read(offset uintptr, bytes uintptr) (*[]byte, error) {
	buffer := make([]byte, bytes, bytes)
	ptr := uintptr(unsafe.Pointer(&buffer))

	process.ReadProcessMemory(p.Handle, offset, &ptr, bytes)
	// _, err := readProcessMemory(p.Handle, offset, &ptr, bytes)

	// if err != nil {
	// 	return nil, err
	// }

	return &buffer, nil
}


func main() {

}
