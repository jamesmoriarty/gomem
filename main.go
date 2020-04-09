package main

type Process struct {
	ID     uint32
	Name   string
	Handle uintptr
}

func GetFromProcessName(name string) (*Process, error) {
	pid, err := getProcessID(name)

	if err != nil {
		return nil, err
	}

	process := Process{ID: pid, Name: name}

	return &process, nil
}

func (p *Process) Open() (uintptr, error) {
	handle, err := openProcess(PROCESS_ALL_ACCESS, false, p.ID)
	
	if err != nil {
		return 0, err
	}

	p.Handle = handle

	return handle, err
}

func main() {

}
