package main

import (
	"fmt"
	"strconv"
	"github.com/jamesmoriarty/gomem/internal/process"
)

// PtrToHex converts uintptr to hex string.
func PtrToHex(ptr uintptr) string {
	s := fmt.Sprintf("%d", ptr)
	n, _ := strconv.Atoi(s)
	h := fmt.Sprintf("0x%x", n)
	return h
}

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
// It returns a *[]byte with the memory contents.
func (p *Process) Read(offset uintptr, bytes uintptr) (*uintptr, error) {
	var buffer uintptr

	// process.ReadProcessMemory(p.Handle, offset, &ptr, bytes)
	_, err := process.ReadProcessMemory(p.Handle, offset, &buffer, bytes)

	if err != nil {
		return nil, err
	}

	return &buffer, nil
}

func main() {

}
