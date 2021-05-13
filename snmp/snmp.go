package snmp

import (
	"errors"
	"log"
	"time"

	"github.com/gosnmp/gosnmp"
)

type SNMP struct {
	snmp *gosnmp.GoSNMP
}

var NoneErr = errors.New("not result")

func NewSNMP(ip string) *SNMP {
	snmp := &gosnmp.GoSNMP{
		Target:             ip,
		Port:               161,
		Transport:          "udp",
		Community:          "public",
		Version:            gosnmp.Version2c,
		Timeout:            time.Second,
		Retries:            0,
		ExponentialTimeout: true,
		MaxOids:            gosnmp.MaxOids,
	}
	return &SNMP{snmp: snmp}
}

func (s *SNMP) Logger(logger *log.Logger) {
	s.snmp.Logger = logger
}

func (s *SNMP) Connect() error {
	return s.snmp.Connect()
}

func (s *SNMP) Close() error {
	return s.snmp.Conn.Close()
}

func (s *SNMP) Get(oid string) (Result, error) {
	oids := []string{oid}
	p, err := s.snmp.Get(oids)
	if err != nil {
		return Result{}, err
	}
	return getFirstResultFromPacket(p)
}

func (s *SNMP) GetAll(oids []string) ([]Result, error) {
	p, err := s.snmp.Get(oids)
	if err != nil {
		return nil, err
	}
	return getAllResultFromPacket(p)
}

func (s *SNMP) GetNext(oid string) (Result, error) {
	oids := []string{oid}
	p, err := s.snmp.GetNext(oids)
	if err != nil {
		return Result{}, err
	}
	return getFirstResultFromPacket(p)
}

func (s *SNMP) GetNextAll(oids []string) ([]Result, error) {
	p, err := s.snmp.GetNext(oids)
	if err != nil {
		return nil, err
	}
	return getAllResultFromPacket(p)
}

func (s *SNMP) Walk(oid string) ([]Result, error) {
	p, err := s.snmp.WalkAll(oid)
	if err != nil {
		return nil, err
	}
	var result []Result
	for _, pdu := range p {
		result = append(result, Result{Data: pdu})
	}
	return result, nil
}

// getFirstResultFromPacket 根据 <packet> 获取第一条结果
func getFirstResultFromPacket(packet *gosnmp.SnmpPacket) (Result, error) {
	result := Result{}
	if len(packet.Variables) == 0 {
		return result, NoneErr
	}
	result = Result{Data: packet.Variables[0]}
	return result, nil
}

// getAllResultFromPacket 根据 <packet> 获取所有结果
func getAllResultFromPacket(packet *gosnmp.SnmpPacket) ([]Result, error) {
	var result []Result
	for _, pdu := range packet.Variables {
		result = append(result, Result{Data: pdu})
	}
	return result, nil
}
