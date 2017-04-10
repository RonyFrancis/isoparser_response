package iso8583

import (
	"errors"
	"fmt"
	"reflect"
)

// Parser for ISO 8583 messages
type Parser struct {
	messages  map[string]reflect.Type
	MtiEncode int
}

// Register MTI
func (p *Parser) Register(mti string, tpl interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Critical error:" + fmt.Sprint(r))
		}
	}()

	if len(mti) != 4 {
		return errors.New("MTI must be a 4 digit numeric field")
	}
	v := reflect.ValueOf(tpl)
	// TODO do more check
	if p.messages == nil {
		p.messages = make(map[string]reflect.Type)
	}
	p.messages[mti] = reflect.Indirect(v).Type()
	fmt.Println("this is p.messages[mti]")
	fmt.Println(reflect.Indirect(v).Type())
	return nil
}

func decodeMti(raw []byte, encode int) (string, error) {
	mtiLen := 4
	if encode == BCD {
		mtiLen = 2
	}
	if len(raw) < mtiLen {
		return "", errors.New("bad MTI raw data")
	}

	var mti string
	switch encode {
	case ASCII:
		mti = string(raw[:mtiLen])
	case BCD:
		mti = string(bcd2Ascii(raw[:mtiLen]))
	default:
		return "", errors.New("invalid encode type")
	}
	return mti, nil
}

//Parse MTI
func (p *Parser) Parse(raw []byte) (ret *Message, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Critical error:" + fmt.Sprint(r))
			ret = nil
		}
	}()
	// fmt.Println("this is raw\n")
	// fmt.Println(string(raw))
	// fmt.Println(reflect.ValueOf(p.MtiEncode))
	// fmt.Println("this is raw\n")
	mti, err := decodeMti(raw, p.MtiEncode)
	if err != nil {
		return nil, err
	}

	tp, ok := p.messages[mti]
	if !ok {
		return nil, errors.New("no template registered for MTI: " + mti)
	}
	// fmt.Println(reflect.TypeOf(tp))
	tpl := reflect.New(tp)
	// fmt.Println(reflect.ValueOf(tp.NumField()))
	initStruct(tp, tpl)
	msg := NewMessage(mti, tpl.Interface())
	msg.MtiEncode = p.MtiEncode
	// fmt.Printf("%+T\n", msg.Data)
	// fmt.Println(reflect.ValueOf(msg.Data))
	// s := reflect.ValueOf(msg.Data).Elem()
	// typeOfT := s.Type()

	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("%d: %s %s = %v\n", i,
	// 		typeOfT.Field(i).Name, f.Type(), f.Interface())
	// }
	// fmt.Println("hello world\n")
	// fmt.Println(reflect.TypeOf(msg.Load(raw)))
	// fmt.Println("hello world\n")
	err = errors.New("Critical error:    zsdkjsdkjfsdk")
	return msg, err
}
func initStruct(tp reflect.Type, val reflect.Value) {
	for i := 0; i < tp.NumField(); i++ {
		field := reflect.Indirect(val).Field(i)
		fieldType := tp.Field(i)
		//fmt.Println(reflect.ValueOf(fieldType))
		switch fieldType.Type.Kind() {
		case reflect.Ptr: // only initialize Ptr fields
			fieldValue := reflect.New(fieldType.Type.Elem())
			field.Set(fieldValue)
		}
	}
}
