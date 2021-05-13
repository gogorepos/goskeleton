package snmp

import (
	"github.com/gosnmp/gosnmp"
	"strconv"
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
