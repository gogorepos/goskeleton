package snmp

import (
	"log"
	"strconv"
	"sync"
	"time"
)

// Fetch 根据 <IP> 获取对应交换机的接口信息
func Fetch(ip string) ([]IfUnit, error) {
	snmp, err := NewSNMP(ip)
	if err != nil {
		return nil, err
	}
	defer snmp.Close()
	// 获取接口数量
	ifNumber, err := snmp.GetIfNumber()
	if err != nil {
		return nil, err
	}
	// 接口设备列表
	ifUnitSlice := make([]IfUnit, ifNumber)
	// 接口描述哈希表 接口描述 ==> 对应接口列表的下标
	ifDescToIndex := make(map[string]int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	getIfUtilization(snmp, ifUnitSlice, &wg)
	// 获取每个接口的描述
	r, err := snmp.Walk(IfDescOid)
	if err != nil {
		return nil, err
	}
	// 遍历结果，保存接口描述，并在哈希表中记录每个接口描述对应的下标
	for i, result := range r {
		description := result.String()
		ifUnitSlice[i].Description = description
		ifDescToIndex[description] = i
	}
	// 获取每个接口的状态
	r, err = snmp.Walk(IfStatusOid)
	if err != nil {
		return nil, err
	}
	for i, result := range r {
		ifUnitSlice[i].Status = result.Int()
	}
	// 获取每个接口的带宽
	r, err = snmp.Walk(IfSpeedOid)
	if err != nil {
		return nil, err
	}
	for i, result := range r {
		ifUnitSlice[i].Speed = result.Int()
	}
	// 获取每个接口的物理地址
	r, err = snmp.Walk(IfPMacOid)
	if err != nil {
		return nil, err
	}
	for i, result := range r {
		ifUnitSlice[i].Mac = result.MacString()
	}
	// 获取端口和其 mac 地址映射表
	portToMac, err := snmp.GetMacAddress()
	if err != nil {
		return nil, err
	}
	for i, _ := range ifUnitSlice {
		if m, ok := portToMac[i+1]; ok {
			ifUnitSlice[i].Mac = m
		}
	}
	// 获取交换机连接其他交换机的端口数
	r, err = snmp.Walk(OccupiedPortOid)
	if err != nil {
		return nil, err
	}
	count := len(r)
	for i := 1; i <= count; i++ {
		iString := strconv.Itoa(i)
		// 获取本地端口
		num, err := snmp.GetNext(IndexLocalPortOid + iString)
		if err != nil {
			continue
		}
		// 获取本地端口 ID
		num, err = snmp.Get(IndexLocalIDOid + iString)
		if err != nil {
			log.Printf("id %v", err)
			continue
		}
		id := num.String()
		// 获取本地端口描述
		num, err = snmp.Get(IndexLocalDesOid + iString)
		if err != nil {
			continue
		}
		description := num.String()
		num, err = snmp.GetNext(IndexRemoteIPOid + iString)
		if err != nil {
			continue
		}
		if index, ok := ifDescToIndex[description]; ok {
			ifUnitSlice[index].Type = "switch"
			ifUnitSlice[index].Name = "Switch"
			ifUnitSlice[index].ID = id
			ifUnitSlice[index].IP = num.String()
		}
	}
	wg.Wait()
	return ifUnitSlice, nil
}

func FetchWithSpeed(ip string) ([]IfUnit, error) {
	snmp, err := NewSNMP(ip)
	if err != nil {
		return nil, err
	}
	defer snmp.Close()
	// 获取接口数量
	ifNumber, err := snmp.GetIfNumber()
	if err != nil {
		return nil, err
	}
	// 接口设备列表
	ifUnitSlice := make([]IfUnit, ifNumber)
	// 接口描述哈希表 接口描述 ==> 对应接口列表的下标
	ifDescToIndex := make(map[string]int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	getIfUtilization(snmp, ifUnitSlice, &wg)
	// 获取每个接口的描述
	r, err := snmp.Walk(IfDescOid)
	if err != nil {
		return nil, err
	}
	// 遍历结果，保存接口描述，并在哈希表中记录每个接口描述对应的下标
	for i, result := range r {
		description := result.String()
		ifUnitSlice[i].Description = description
		ifDescToIndex[description] = i
	}
	// 获取每个接口的状态
	r, err = snmp.Walk(IfStatusOid)
	if err != nil {
		return nil, err
	}
	for i, result := range r {
		ifUnitSlice[i].Status = result.Int()
	}
	// 获取每个接口的带宽
	r, err = snmp.Walk(IfSpeedOid)
	if err != nil {
		return nil, err
	}
	for i, result := range r {
		ifUnitSlice[i].Speed = result.Int()
	}
	// 获取每个接口的物理地址
	r, err = snmp.Walk(IfPMacOid)
	if err != nil {
		return nil, err
	}
	for i, result := range r {
		ifUnitSlice[i].Mac = result.MacString()
	}
	// 获取端口和其 mac 地址映射表
	portToMac, err := snmp.GetMacAddress()
	if err != nil {
		return nil, err
	}
	for i, _ := range ifUnitSlice {
		if m, ok := portToMac[i+1]; ok {
			ifUnitSlice[i].Mac = m
		}
	}
	// 获取交换机连接其他交换机的端口数
	r, err = snmp.Walk(OccupiedPortOid)
	if err != nil {
		return nil, err
	}
	count := len(r)
	for i := 1; i <= count; i++ {
		iString := strconv.Itoa(i)
		// 获取本地端口
		num, err := snmp.GetNext(IndexLocalPortOid + iString)
		if err != nil {
			continue
		}
		// 获取本地端口 ID
		num, err = snmp.Get(IndexLocalIDOid + iString)
		if err != nil {
			log.Printf("id %v", err)
			continue
		}
		id := num.String()
		// 获取本地端口描述
		num, err = snmp.Get(IndexLocalDesOid + iString)
		if err != nil {
			continue
		}
		description := num.String()
		num, err = snmp.GetNext(IndexRemoteIPOid + iString)
		if err != nil {
			continue
		}
		if index, ok := ifDescToIndex[description]; ok {
			ifUnitSlice[index].Type = "switch"
			ifUnitSlice[index].Name = "Switch"
			ifUnitSlice[index].ID = id
			ifUnitSlice[index].IP = num.String()
		}
	}
	wg.Wait()
	return ifUnitSlice, nil
}

// getIfUtilization 获取每个端口的速率
func getIfUtilization(s *SNMP, ifUnitSlice []IfUnit, wg *sync.WaitGroup) {
	defer wg.Done()
	var startInOctets []int
	var endInOctets []int
	var startOutOctets []int
	var endOutOctets []int
	// 获取端口接受字节数
	r, _ := s.Walk(IfInOctetOid)
	for _, result := range r {
		startInOctets = append(startInOctets, result.Int())
	}
	// 获取端口发送字节数
	r, _ = s.Walk(IfOutOctetOid)
	for _, result := range r {
		startOutOctets = append(startOutOctets, result.Int())
	}
	// 暂停 1 S
	time.Sleep(time.Second)
	// 再次获取
	r, _ = s.Walk(IfInOctetOid)
	for _, result := range r {
		endInOctets = append(endInOctets, result.Int())
	}
	// 获取端口发送字节数
	r, _ = s.Walk(IfOutOctetOid)
	for _, result := range r {
		endOutOctets = append(endOutOctets, result.Int())
	}
	// 结合带宽计算
	for i, octet := range startInOctets {
		ifUnitSlice[i].InSpeed = endInOctets[i] - octet
	}
	for i, octet := range startOutOctets {
		speed := ifUnitSlice[i].Speed
		if speed == 0 {
			continue
		}
		ifUnitSlice[i].OutSpeed = endOutOctets[i] - octet
	}
}
