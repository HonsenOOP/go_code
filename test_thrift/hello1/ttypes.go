// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package hello1

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Pair struct {
	Value string `thrift:"value,1,required" json:"value"`
	Key   string `thrift:"key,2,required" json:"key"`
}

func NewPair() *Pair {
	return &Pair{}
}

func (p *Pair) GetKey() string {
	return p.Key
}

func (p *Pair) GetValue() string {
	return p.Value
}
func (p *Pair) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Pair) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Key = v
	}
	return nil
}

func (p *Pair) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *Pair) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Pair"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Pair) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:value: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Value)); err != nil {
		return fmt.Errorf("%T.value (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:value: %s", p, err)
	}
	return err
}

func (p *Pair) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("key", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:key: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Key)); err != nil {
		return fmt.Errorf("%T.key (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:key: %s", p, err)
	}
	return err
}

func (p *Pair) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Pair(%+v)", *p)
}
