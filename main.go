package main

type Process struct {
	ID uint32
	Name string
	Handle uint32
}

func GetFromProcessName(name string) (*Process, error) {
	pid, err := getProcessID(name)

	if err != nil {
		return nil, err
	}

	process := Process{ID: pid, Name: name}

	return &process, nil
}


func main(){

}