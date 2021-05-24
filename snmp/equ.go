package snmp

import (
	"github.com/gogorepos/goskeleton/util/client"
	"github.com/tidwall/gjson"
)

type EquEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Equ struct {
	IP   string
	Mac  string
	Type string
	Name string
}

type Equs struct {
	Map  map[string]int
	Data []*Equ
}

var command = map[string]interface{}{
	"idList":  []string{},
	"mac":     "y",
	"devType": "y",
	"ip":      "y",
	"name":    "y",
}

var equSlice = []string{
	"hscmd-get-rx-config",
	"hscmd-get-tx-config",
	"hscmd-get-sprx-config",
	"hscmd-get-tx-config",
	"hscmd-get-vs10-config",
	"hscmd-get-pwr750-config",
	"hscmd-get-rt10t-config",
	"hscmd-get-usb-rx-config",
	"hscmd-get-usb-tx-config",
}

func removeEmptyMac(results []gjson.Result) []gjson.Result {
	result := results[:0]
	for _, r := range results {
		if r.Get("mac").String() != "" {
			result = append(result, r)
		}
	}
	return result
}

func Get() (*Equs, error) {
	e := &EquEngine{
		Scheduler:   &EquScheduler{},
		WorkerCount: 5,
	}
	return e.Get()
}

func (e *EquEngine) Get() (*Equs, error) {
	in := make(chan string)
	out := make(chan gjson.Result)
	equs := &Equs{
		Map:  make(map[string]int),
		Data: make([]*Equ, 0, 50),
	}
	count := 0
	e.Scheduler.ConfigureWorkerChan(in)
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}
	for _, header := range equSlice {
		e.Scheduler.Submit(header)
	}
	for {
		result := <-out
		r := removeEmptyMac(result.Array())
		equsMapLen := len(equs.Map)
		for i, g := range r {
			mac := g.Get("mac").String()
			equ := Equ{
				IP:   g.Get("ip").String(),
				Mac:  mac,
				Type: g.Get("devType").String(),
				Name: g.Get("name").String(),
			}
			equs.Map[mac] = equsMapLen + i
			equs.Data = append(equs.Data, &equ)
		}
		count++
		if count >= len(equSlice) {
			break
		}
	}

	return equs, nil
}

func createWorker(in chan string, out chan gjson.Result) {
	go func() {
		for {
			header := <-in
			result, err := equWorker(header)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func equWorker(header string) (gjson.Result, error) {
	result, err := client.IpsscSend(header, command)
	if err != nil {
		return gjson.Result{}, err
	}
	return result.Get("parameter"), nil
}
