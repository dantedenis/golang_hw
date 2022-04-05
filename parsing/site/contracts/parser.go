package contracts

import (
	"io"
	"parsing/site"
)

type Parser interface {
	ListSite(response io.Reader) []site.RawData
}
