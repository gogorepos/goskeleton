package snmp

type EquScheduler struct {
	workerChan chan string
}

func (e *EquScheduler) ConfigureWorkerChan(c chan string) {
	e.workerChan = c
}

func (e *EquScheduler) Submit(header string) {
	go func() {
		e.workerChan <- header
	}()
}
