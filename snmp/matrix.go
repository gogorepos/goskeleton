package snmp

import (
	"github.com/gogorepos/goskeleton/util/mac"
	"sync"
)

type Switch struct {
	LocalIP string // 本地 IP 地址
	mu      sync.Mutex
	IfUnit  []IfUnit // 接口
}

type IfUnit struct {
	Description  string  // 接口信息描述
	MTU          int     // 发送和接受的最大 IP 数据报 byte
	Speed        uint    // 带宽 bps
	PhysAddress  mac.Mac // MAC 地址
	Status       int     // 操作状态
	InOctet      uint    // 收到的字节数
	OutOctet     uint    // 发送的字节数
	InUcastPkts  uint    // 收到的数据包个数
	OutUcastPkts uint    // 发送的数据包个数
}

func GetSwitch(ip string) (*Switch, error) {
	s := Switch{LocalIP: ip, mu: sync.Mutex{}}
	snmp := NewSNMP(ip)
	if err := snmp.Connect(); err != nil {
		return nil, err
	}
	defer snmp.Close()
	num, err := snmp.Get(IfNumberOid)
	if err != nil {
		return nil, err
	}
	s.IfUnit = make([]IfUnit, num.Int())

	wg := sync.WaitGroup{}
	wg.Add(9)
	go getIfDescr(snmp, &s, &wg)
	go getIfMTU(snmp, &s, &wg)
	go getIfMac(snmp, &s, &wg)
	go getIfInOctet(snmp, &s, &wg)
	go getIfOutOctet(snmp, &s, &wg)
	go getIfInUcastPkts(snmp, &s, &wg)
	go getIfOutUcastPkts(snmp, &s, &wg)
	go getIfStatus(snmp, &s, &wg)
	go getIfSpeed(snmp, &s, &wg)
	wg.Wait()
	return &s, nil
}

func getIfDescr(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfDescrOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].Description = result.String()
		s.mu.Unlock()
	}
}

func getIfStatus(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfOperStatusOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].Status = result.Int()
		s.mu.Unlock()
	}
}

func getIfMTU(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfOperStatusOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].MTU = result.Int()
		s.mu.Unlock()
	}
}

func getIfSpeed(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfOperStatusOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].Speed = result.Uint()
		s.mu.Unlock()
	}
}

func getIfInOctet(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfOperStatusOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].InOctet = result.Uint()
		s.mu.Unlock()
	}
}

func getIfOutOctet(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfOperStatusOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].OutOctet = result.Uint()
		s.mu.Unlock()
	}
}

func getIfInUcastPkts(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfOperStatusOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].InUcastPkts = result.Uint()
		s.mu.Unlock()
	}
}

func getIfOutUcastPkts(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfOperStatusOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].OutUcastPkts = result.Uint()
		s.mu.Unlock()
	}
}

func getIfMac(snmp *SNMP, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(IfPhysAddressOid)
	if err != nil {
		return
	}
	for i, result := range r {
		s.mu.Lock()
		s.IfUnit[i].PhysAddress = result.UintSlice()
		s.mu.Unlock()
	}
}
