package snmp

type Scheduler interface {
	Submit(string)
	ConfigureWorkerChan(chan string)
}

type IfUnit struct {
	ID          string
	IP          string
	Description string
	Type        string
	Name        string
	Mac         string
	Speed       int
	Status      int
}
