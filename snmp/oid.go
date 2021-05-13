package snmp

const (
	// IfNumberOid [Get] 网络接口的数目
	IfNumberOid = ".1.3.6.1.2.1.2.1.0"
	// IfDescOid [Walk] 网络接口信息描述
	IfDescOid = ".1.3.6.1.2.1.2.2.1.2"
	// IfTypeOid [Walk] 网络接口类型
	IfTypeOid = ".1.3.6.1.2.1.2.2.1.3"
	// IfMTUOid [Walk] 接口发送和接收的最大 IP 数据报
	IfMTUOid = ".1.3.6.1.2.1.2.2.1.4"
	// IfSpeedOid [Walk] 接口当前宽带 bps
	IfSpeedOid = ".1.3.6.1.2.1.2.2.1.5"
	// IfMacOid [Walk] 接口的物理地址
	IfMacOid = ".1.3.6.1.2.1.2.2.1.6"
	// IfStatusOid [Walk] 接口当前操作状态
	IfStatusOid = ".1.3.6.1.2.1.2.2.1.8"
	// IfInOctetOid [Walk] 接口收到的字节数
	IfInOctetOid = ".1.3.6.1.2.1.2.2.1.10"
	// IfOutOctetOid [Walk] 接口发送的字节数
	IfOutOctetOid = ".1.3.6.1.2.1.2.2.1.16"
	// IfInUcastPktsOid [Walk] 接口收到的数据包个数
	IfInUcastPktsOid = ".1.3.6.1.2.1.2.2.1.11"
	// IfOutUcastPktsOid [Walk] 接口发送的数据包个数
	IfOutUcastPktsOid = ".1.3.6.1.2.1.2.2.1.17"

	// OccupiedPortOid [Walk] 交换机占用的端口
	OccupiedPortOid = ".1.3.6.1.4.1.3320.127.1.4.2.1.4"
	// IndexRemoteIPOid [GetNext] 交换机某端口的远程 IP
	IndexRemoteIPOid = "1.3.6.1.4.1.3320.127.1.4.2.1.2."
	// IndexRemoteIDOid [GetNext] 交换机某端口的远程端口 ID
	IndexRemoteIDOid = "1.3.6.1.4.1.3320.127.1.4.1.1.7."
	// IndexRemoteDesOid [GetNext] 交换机某端口的远程端口描述
	IndexRemoteDesOid = "1.3.6.1.4.1.3320.127.1.4.1.1.8."
	// IndexLocalPortOid [GetNext] 交换机某端口的本地端口
	IndexLocalPortOid = "1.3.6.1.4.1.3320.127.1.4.1.1.2."
	// IndexLocalIDOid [Get] 交换机某端口的本地端口 ID
	IndexLocalIDOid = "1.3.6.1.4.1.3320.127.1.3.7.1.3."
	// IndexLocalDesOid [Get] 交换机某端口的本地端口描述
	IndexLocalDesOid = "1.3.6.1.4.1.3320.127.1.3.7.1.4."
)

var OidMap = map[string]string{
	"IfNumberOid":       ".1.3.6.1.2.1.2.1.0",
	"IfDescOid":         ".1.3.6.1.2.1.2.2.1.2",
	"IfTypeOid":         ".1.3.6.1.2.1.2.2.1.3",
	"IfMTUOid":          ".1.3.6.1.2.1.2.2.1.4",
	"IfSpeedOid":        ".1.3.6.1.2.1.2.2.1.5",
	"IfMacOid":          ".1.3.6.1.2.1.2.2.1.6",
	"IfStatusOid":       ".1.3.6.1.2.1.2.2.1.8",
	"IfInOctetOid":      ".1.3.6.1.2.1.2.2.1.10",
	"IfOutOctetOid":     ".1.3.6.1.2.1.2.2.1.16",
	"IfInUcastPktsOid":  ".1.3.6.1.2.1.2.2.1.11",
	"IfOutUcastPktsOid": ".1.3.6.1.2.1.2.2.1.17",
}
