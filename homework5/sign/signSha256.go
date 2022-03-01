package sign

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"homework5/sign/contract"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const separator = ">>>>>>>>>>sign<<<<<<<<<<"

type SignatureSha256 struct {
	date      time.Time
	name      string
	size      string
	signature []byte
}

func (sig SignatureSha256) encrypt(text string) (b []byte, err error) {
	sha := sha256.New()
	_, err = sha.Write([]byte(text))
	if err != nil {
		return
	}
	return sha.Sum(nil), nil
}

func NewSignatureSha256FromFile(file *os.File, hashString string) (sign SignatureSha256, err error) {
	stat, err := file.Stat()
	if err != nil {
		return
	}
	sign.size = strconv.FormatInt(stat.Size(), 10)
	sign.date = stat.ModTime()
	sign.name = path.Base(file.Name())
	fileData := make([]byte, stat.Size())
	countBytes, err := file.Read(fileData)
	if err != nil {
		return
	} else if countBytes == 0 {
		err = errors.New("file is empty")
		return
	}
	data := string(fileData) + hashString
	sign.signature, err = sign.encrypt(data)
	if err != nil {
		return
	}
	return
}

func New(date time.Time, name string, size string, sign []byte) *SignatureSha256 {
	return &SignatureSha256{date: date, name: name, size: size, signature: sign}
}

func NewSignatureSha256(sigSha contract.Signature) *SignatureSha256 {
	return &SignatureSha256{date: sigSha.Date(), size: sigSha.Size(), name: sigSha.Name(), signature: sigSha.SignatureBytes()}
}

func (sig SignatureSha256) headerString() string {
	return strings.Join([]string{sig.Date().Format("2006-01-02 15-04-05"), sig.size, sig.name, separator}, ":")
}

func (sig SignatureSha256) SignatureBytes() []byte {
	result := bytes.NewBufferString(sig.headerString())
	result.WriteString(":")
	result.Write(sig.signature)
	return result.Bytes()
}

func (sig SignatureSha256) Date() time.Time {
	return sig.date
}

func (sig SignatureSha256) Size() string {
	return sig.size
}

func (sig SignatureSha256) Name() string {
	return sig.name
}

func different(a, b []byte) error {
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

func (sig SignatureSha256) Equal(s string) (bool, error) {
	file, err := os.Open(s)
	if err != nil {
		return false, err
	}

	defer func() {
		fError := file.Close()
		if fError != nil {
			err = fError
		}
	}()
	stat, err := file.Stat()
	if err != nil {
		return false, err
	}
	fileData := make([]byte, stat.Size())
	countBytes, err := file.Read(fileData)
	if err != nil {
		return false, err
	} else if countBytes == 0 {
		return false, errors.New("file is empty")
	}
	if bytes.Equal(sig.SignatureBytes(), fileData) == true {
		return true, nil
	} else {
		return false, different(sig.SignatureBytes(), fileData)
	}
}
