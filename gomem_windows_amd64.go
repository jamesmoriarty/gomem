package gomem

import (
	"unsafe"

	kernal32 "github.com/jamesmoriarty/gomem/internal/kernel32"
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
func (p *Process) Read(offset uintptr, buffer uintptr, length uintptr) error {
	_, err := kernal32.ReadProcessMemory(p.Handle, offset, buffer, length)

	return err
}

// Read byte from process memory.
func (p *Process) ReadByte(offset uintptr) (byte, error) {
	var (
		value    byte
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read uint32 from process memory.
func (p *Process) ReadUInt32(offset uintptr) (uint32, error) {
	var (
		value    uint32
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Read uint64 from process memory.
func (p *Process) ReadUInt64(offset uintptr) (uint64, error) {
	var (
		value    uint64
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	err := p.Read(offset, valuePtr, unsafe.Sizeof(value))

	return value, err
}

// Write process memory.
func (p *Process) Write(offset uintptr, buffer uintptr, length uintptr) error {
	_, err := kernal32.WriteProcessMemory(p.Handle, offset, buffer, length)

	return err
}

// Write byte to process memory.
func (p *Process) WriteByte(offset uintptr, value byte) error {
	var (
		valuePtr = (uintptr)(unsafe.Pointer(&value))
	)

	return p.Write(offset, valuePtr, unsafe.Sizeof(value))
}

// GetModule address.
func (p *Process) GetModule(name string) (uintptr, error) {
	ptr, err := kernal32.GetModule(name, p.ID)

	return ptr, err
}

// IsKeyDown https://docs.microsoft.com/en-gb/windows/win32/inputdev/virtual-key-codes
func IsKeyDown(v int) bool {
	return user32.GetAsyncKeyState(v) > 0
}
