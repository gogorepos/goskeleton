package main

import (
	"fmt"
	"github.com/gogorepos/goskeleton/snmp"
)

func main() {
	r, _ := snmp.GetSwitch("169.254.0.5")
	for _, ifUnit := range r["169.254.0.5"].IfUnit {
		fmt.Printf("%#v\n", ifUnit)
	}
}
