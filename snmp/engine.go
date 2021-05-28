package snmp

import (
	"encoding/json"
)

// Run 获取 <ips> 地址交换机接口信息
func Run(ips ...string) []byte {
	var (
		// 已访问的 ip 表
		connectedIpMap = make(map[string]bool)
		// ip 对应接口信息表
		switchIpToUnit = make(map[string][]IfUnit)
	)
	// 从服务器获取所有设备信息
	equs, _ := Get()
	for len(ips) > 0 {
		// 获取 IP 地址队列第一个元素
		ip := ips[0]
		ips = ips[1:]
		// 获取对应交换机信息
		ifUnitSlice, err := Fetch(ip)
		if err != nil {
			continue
		}
		// 标记已访问
		connectedIpMap[ip] = true
		// 根据设备信息的 mac 地址与接口中的 mac 地址对比，如果存在则添加设备信息
		// 并找出其中连接其他交换机的接口，将 IP 地址加入队列
		for i, unit := range ifUnitSlice {
			if index, ok := equs.Map[unit.Mac]; ok {
				e := equs.Data[index]
				ifUnitSlice[i].ID = e.ID
				ifUnitSlice[i].IP = e.IP
				ifUnitSlice[i].Type = e.Type
				ifUnitSlice[i].Name = e.Name
			}
			remoteIP := unit.IP
			if unit.Type == "switch" && remoteIP != "" && !connectedIpMap[remoteIP] {
				ips = append(ips, remoteIP)
			}
		}
		switchIpToUnit[ip] = ifUnitSlice
	}
	// JSON 序列化，以更好的格式
	r, _ := json.MarshalIndent(switchIpToUnit, "", "  ")
	return r
}

// RunWithSpeed 获取 <ips> 地址交换机接口信息
func RunWithSpeed(ips ...string) []byte {
	var (
		// 已访问的 ip 表
		connectedIpMap = make(map[string]bool)
		// ip 对应接口信息表
		switchIpToUnit = make(map[string][]IfUnit)
	)
	// 从服务器获取所有设备信息
	equs, _ := Get()
	for len(ips) > 0 {
		// 获取 IP 地址队列第一个元素
		ip := ips[0]
		ips = ips[1:]
		// 获取对应交换机信息
		ifUnitSlice, err := FetchWithSpeed(ip)
		if err != nil {
			continue
		}
		// 标记已访问
		connectedIpMap[ip] = true
		// 根据设备信息的 mac 地址与接口中的 mac 地址对比，如果存在则添加设备信息
		// 并找出其中连接其他交换机的接口，将 IP 地址加入队列
		for i, unit := range ifUnitSlice {
			if index, ok := equs.Map[unit.Mac]; ok {
				e := equs.Data[index]
				ifUnitSlice[i].ID = e.ID
				ifUnitSlice[i].IP = e.IP
				ifUnitSlice[i].Type = e.Type
				ifUnitSlice[i].Name = e.Name
			}
			remoteIP := unit.IP
			if unit.Type == "switch" && remoteIP != "" && !connectedIpMap[remoteIP] {
				ips = append(ips, remoteIP)
			}
		}
		switchIpToUnit[ip] = ifUnitSlice
	}
	// JSON 序列化，以更好的格式
	r, _ := json.MarshalIndent(switchIpToUnit, "", "  ")
	return r
}
