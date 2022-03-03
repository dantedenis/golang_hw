package packer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

func Pack(sig interface{}) (buf *bytes.Buffer, err error) {
	v := reflect.ValueOf(sig)
	if v.Kind() != reflect.Struct {
		err = errors.New("is not struct type")
		return
	}
	orderByte := binary.BigEndian
	buff := &bytes.Buffer{}
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Uint32, reflect.Uint64:
			s := v.Field(i).Uint()
			//_ = s
			err = binary.Write(buff, orderByte, byte(s))
			if err != nil {
				return
			}
		case reflect.String:
			s := v.Field(i).String()
			b := []byte(s)
			err = binary.Write(buff, orderByte, uint16(len(b)))
			if err != nil {
				return
			}
			err = binary.Write(buff, orderByte, b)
		case reflect.Slice:
			err = binary.Write(buff, orderByte, uint16(len(v.Field(i).Bytes())))
			err = binary.Write(buff, orderByte, v.Field(i).Bytes())
		default:
			fmt.Printf("undefined type field %v\n", v.Field(i).Type().Kind())
		}
	}
	buf = buff
	return
}
