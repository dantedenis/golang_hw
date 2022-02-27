package contract

import "time"

type Signature interface {
	Date() time.Time
	Size() string
	Name() string
	SignatureBytes() []byte
	Equal(string) (bool, error)
}
