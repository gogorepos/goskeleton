package mac

import (
	"strconv"
	"strings"
)

type Mac []uint

func (m Mac) String() string {
	if len(m) != 6 {
		return ""
	}
	builder := strings.Builder{}
	for _, u := range m {
		n := strconv.FormatUint(uint64(u), 16)
		builder.WriteString(n + ".")
	}
	return builder.String()[:builder.Len()-1]
}
