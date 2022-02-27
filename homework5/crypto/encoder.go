package crypto

import (
	"errors"
	"homework5/sign"
	"homework5/sign/contract"
	"io/ioutil"
	"os"
	"strings"
)

type Encoder struct {
	hashSign   string
	fileSource string
	signature  contract.Signature
}

func NewEncoder(fileSource, fileHashString string) (enc *Encoder, err error) {
	hashString, err := ioutil.ReadFile(fileHashString)
	if err != nil {
		return
	}
	enc = &Encoder{fileSource: fileSource, hashSign: string(hashString)}
	return
}

func (enc *Encoder) EncryptSha256() error {
	file, err := os.Open(enc.fileSource)
	if err != nil {
		return err
	}
	defer func() {
		fError := file.Close()
		if fError != nil {
			err = fError
		}
	}()
	sig, err := sign.NewSignatureSha256FromFile(file, enc.hashSign)
	if err != nil {
		return err
	}
	enc.signature = sig
	return err
}
func (enc Encoder) SaveToFile(path string) error {
	err := ioutil.WriteFile(path, enc.signature.SignatureBytes(), 0644)
	return err
}

func (enc Encoder) GetSign() contract.Signature {
	return enc.signature
}

func Different(a, b []byte) error {
	signFirst, signSecond := strings.Split(string(a), ":"), strings.Split(string(b), ":")
	var result string
	if signFirst[0] != signSecond[0] {
		result += "date is modify\n"
	}
	if signFirst[1] != signSecond[1] {
		result += "size is modify\n"
	}
	if signFirst[2] != signSecond[2] {
		result += "name is modify\n"
	}
	if signFirst[3] != signSecond[3] {
		result += "header is modify\n"
	}
	if signFirst[4] != signSecond[4] {
		result += "data file is modify\n"
	}
	return errors.New(result)
}
