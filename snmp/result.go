package snmp

import (
	"strconv"
	"strings"

	"github.com/gosnmp/gosnmp"
)

type Result struct {
	gosnmp.SnmpPDU
}

func (r Result) String() string {
	if n, ok := r.Value.([]byte); ok {
		return string(n)
	}
	if n, ok := r.Value.(int); ok {
		return strconv.Itoa(n)
	}
	return ""
}

func (r Result) Int() int {
	return int(gosnmp.ToBigInt(r.Value).Int64())
}

func (r Result) Uint() uint {
	if n, ok := r.Value.(uint); ok {
		return n
	}
	return 0
}

func (r Result) MacString() string {
	builder := strings.Builder{}
	if v, ok := r.Value.([]uint8); ok {
		// 不是一个规范的 mac 地址，返回空字符串
		if len(v) != 6 {
			return ""
		}
		for _, u := range v {
			// 将十进制数字转为大写形式的十六进制的字符串
			n := strings.ToUpper(strconv.FormatUint(uint64(u), 16))
			// 如果字符串只有一位，补零
			if len(n) < 2 {
				builder.WriteString("0" + n + ":")
			} else {
				builder.WriteString(n + ":")
			}
		}
	}
	// 去除最后一个 ":" 冒号
	return builder.String()[:builder.Len()-1]
}
