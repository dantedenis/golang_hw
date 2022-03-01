package sign

import (
	"github.com/golang/mock/gomock"
	"homework5/mock"
	"os"
	"testing"
)

func TestNewSignatureSha256FromFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSign := mock.NewMockSignature(ctrl)
	file, _ := os.Open("source.txt")
	defer file.Close()
	sign1, _ := NewSignatureSha256FromFile(file, "13123131313131")
	//sign2, _ := NewSignatureSha256FromFile(file, "3123131231231231231")
	mockSign.EXPECT().SignatureBytes().Return(SignatureSha256{signature: sign1.SignatureBytes()})
}
