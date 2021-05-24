package client

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

type Ipssc struct {
	Header string      `json:"cmd_header"`
	Body   interface{} `json:"cmd_body"`
}

func (i Ipssc) Address() string {
	return "169.254.0.253:6000"
}

func (i Ipssc) Command() []byte {
	command, _ := json.Marshal(i)
	return command
}

func IpsscSend(header string, body interface{}) (gjson.Result, error) {
	r, err := Send(Ipssc{
		Header: header,
		Body:   body,
	})
	if err != nil {
		return gjson.Result{}, err
	}
	return gjson.GetBytes(r, "cmd_body"), nil
}
