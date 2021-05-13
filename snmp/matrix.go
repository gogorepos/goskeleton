package snmp

type Switch struct {
	LocalIP string // 本地 IP 地址
	If      []If   // 接口
}

type If struct {
	Description  string // 接口信息描述
	Type         string // 接口类型
	MTU          string // 发送和接受的最大 IP 数据报 byte
	Speed        string // 带宽 bps
	PhysAddress  string // 物理地址
	Status       string // 操作状态
	InOctet      string // 收到的字节数
	OutOctet     string // 发送的字节数
	InUcastPkts  string // 收到的数据包个数
	OutUcastPkts string // 发送的数据包个数
}
