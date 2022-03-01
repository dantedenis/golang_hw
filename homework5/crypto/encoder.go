package crypto

import (
	"homework5/sign"
	"homework5/sign/contract"
	"io/ioutil"
	"os"
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
