package snmp

import (
	"strconv"

	"github.com/gosnmp/gosnmp"
)

type Result struct {
	Data gosnmp.SnmpPDU
}

func (r Result) String() string {
	switch r.Data.Type {
	case gosnmp.Integer:
		return strconv.Itoa(r.Data.Value.(int))
	}
	return string(r.Data.Value.([]byte))
}

func (r Result) Int() int {
	return r.Data.Value.(int)
}

func (r Result) Uint() uint {
	return r.Data.Value.(uint)
}

func (r Result) UintSlice() []uint {
	s := r.Data.Value.([]uint8)
	result := make([]uint, 0, len(s))
	for _, u := range s {
		result = append(result, uint(u))
	}
	return result
}
