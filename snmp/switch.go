package snmp

import (
	"github.com/gosnmp/gosnmp"
)

// GetIfNumber 获取交换机端口个数
func (s *SNMP) GetIfNumber() (int, error) {
	num, err := s.Get(IfNumberOid)
	if err != nil {
		return 0, err
	}
	return num.Int(), nil
}

// GetMacAddress 获取连接接口的设备的 MAC 地址
// 返回值为 接口索引 ==> mac 地址，一个端口的 mac 地址可能有多个
func (s *SNMP) GetMacAddress() (map[int]string, error) {
	// 端口对应 mac 地址
	portToMac := make(map[int]string)
	// mac 地址数组
	macSlice := make([]string, 0)
	// 获取所有 mac 地址
	if err := s.WalkFunc(IfMacOid, func(u gosnmp.SnmpPDU) error {
		r := Result{u}
		macSlice = append(macSlice, r.MacString())
		return nil
	}); err != nil {
		return nil, err
	}
	// 获取 mac 地址对应端口
	r, err := s.BulkWalk(IfMacPortOid)
	if err != nil {
		return nil, err
	}
	for i, result := range r {
		index := result.Int()
		// 如果一个端口对应多个 mac 地址，则该端口可能连接其他交换机
		if _, ok := portToMac[index]; ok {
			portToMac[index] = "other"
		} else {
			portToMac[index] = macSlice[i]
		}
	}
	return portToMac, nil
}
