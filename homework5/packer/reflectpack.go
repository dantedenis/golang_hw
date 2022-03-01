package packer

import (
	"bytes"
	"encoding/binary"
	"errors"
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

		}
	}
}
