package packer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func PackTime(time interface{}) (buf *bytes.Buffer, err error) {
	v := reflect.ValueOf(time)
	orderByte := binary.BigEndian
	buff := &bytes.Buffer{}
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Uint64:
			s := v.Field(i).Uint()
			err = binary.Write(buff, orderByte, byte(s))
			if err != nil {
				return
			}
		case reflect.Int64:
			s := v.Field(i).Int()
			err = binary.Write(buff, orderByte, byte(s))
			if err != nil {
				return
			}
		default:
			fmt.Printf("(PackTime)undefined type field %v\n", v.Field(i).Type().Kind())
		}
	}
	buf = buff
	return
}

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
		case reflect.Struct:
			if v.Field(i).Type() == reflect.TypeOf(time.Time{}) {
				b, er := PackTime(v.Field(i))
				if er != nil {
					return
				}
				err = binary.Write(buff, orderByte, b)
			}
		default:
			fmt.Printf("(PACK)undefined type field %v\n", v.Field(i).Type().Kind())
		}
	}
	buf = buff
	return
}
