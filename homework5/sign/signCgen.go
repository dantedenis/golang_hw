package sign

import "encoding/binary"
import "bytes"

func (in *SignatureSha256) Unpack(data []byte) error {
	r := bytes.NewReader(data)

	// name
	var nameLenRaw uint8
	binary.Read(r, binary.BigEndian, &nameLenRaw)
	nameRaw := make([]byte, nameLenRaw)
	binary.Read(r, binary.BigEndian, &nameRaw)
	in.name = string(nameRaw)

	// size
	var sizeLenRaw uint8
	binary.Read(r, binary.BigEndian, &sizeLenRaw)
	sizeRaw := make([]byte, sizeLenRaw)
	binary.Read(r, binary.BigEndian, &sizeRaw)
	in.size = string(sizeRaw)

	// signature
	var signatureLenRaw uint8
	binary.Read(r, binary.BigEndian, &signatureLenRaw)
	signatureRaw := make([]byte, signatureLenRaw)
	binary.Read(r, binary.BigEndian, &signatureRaw)
	in.signature = signatureRaw
	return nil
}
