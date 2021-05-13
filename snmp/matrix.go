package snmp

import (
	"sync"

	"github.com/gogorepos/goskeleton/util/mac"
)

type Switch struct {
	LocalIP string // 本地 IP 地址
	mu      sync.Mutex
	IfUnit  []IfUnit // 接口
}

type IfUnit struct {
	Description  string // 接口信息描述
	MTU          int    // 发送和接受的最大 IP 数据报 byte
	Speed        uint   // 带宽 bps
	Mac          string // MAC 地址
	Status       int    // 操作状态
	InOctet      uint   // 收到的字节数
	OutOctet     uint   // 发送的字节数
	InUcastPkts  uint   // 收到的数据包个数
	OutUcastPkts uint   // 发送的数据包个数
}

func GetSwitch(ip string) (map[string]*Switch, error) {
	var (
		switchs = make(map[string]*Switch)
		ipMap   = []string{ip}
	)

	switchs[ip], _ = findSwitch(ip, ipMap)
	return switchs, nil
}

func findSwitch(ip string, ipMap []string) (*Switch, error) {
	snmp := NewSNMP(ip)
	if err := snmp.Connect(); err != nil {
		return nil, err
	}
	defer snmp.Close()
	n, err := snmp.Get(IfNumberOid)
	if err != nil {
		return nil, err
	}
	s := &Switch{mu: sync.Mutex{}, LocalIP: ip, IfUnit: make([]IfUnit, n.Int())}
	wg := sync.WaitGroup{}
	wg.Add(9)
	go getInfoByName(snmp, "IfDescOid", s, &wg)
	go getInfoByName(snmp, "IfMTUOid", s, &wg)
	go getInfoByName(snmp, "IfSpeedOid", s, &wg)
	go getInfoByName(snmp, "IfMacOid", s, &wg)
	go getInfoByName(snmp, "IfStatusOid", s, &wg)
	go getInfoByName(snmp, "IfInOctetOid", s, &wg)
	go getInfoByName(snmp, "IfOutOctetOid", s, &wg)
	go getInfoByName(snmp, "IfInUcastPktsOid", s, &wg)
	go getInfoByName(snmp, "IfOutUcastPktsOid", s, &wg)
	wg.Wait()
	return s, nil
}

func getInfoByName(snmp *SNMP, t string, s *Switch, wg *sync.WaitGroup) {
	defer wg.Done()
	r, err := snmp.Walk(OidMap[t])
	if err != nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, result := range r {
		switch t {
		case "IfDescOid":
			s.IfUnit[i].Description = result.String()
		case "IfMTUOid":
			s.IfUnit[i].MTU = result.Int()
		case "IfStatusOid":
			s.IfUnit[i].Status = result.Int()
		case "IfSpeedOid":
			s.IfUnit[i].Speed = result.Uint()
		case "IfInOctetOid":
			s.IfUnit[i].InOctet = result.Uint()
		case "IfOutOctetOid":
			s.IfUnit[i].OutOctet = result.Uint()
		case "IfInUcastPktsOid":
			s.IfUnit[i].InUcastPkts = result.Uint()
		case "IfOutUcastPktsOid":
			s.IfUnit[i].OutUcastPkts = result.Uint()
		case "IfMacOid":
			s.IfUnit[i].Mac = mac.Mac(result.UintSlice()).String()
		}
	}
}
