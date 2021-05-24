package main

import (
	"github.com/gogorepos/goskeleton/snmp"
)

func main() {
	e := &snmp.EquEngine{
		Scheduler:   &snmp.EquScheduler{},
		WorkerCount: 5,
	}
	e.Get()
}
