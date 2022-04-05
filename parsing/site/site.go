package site

import (
	"net/http"
	"parsing/site/contracts"
	"time"
)

type Site struct {
	url    string
	parser contracts.Parser
}

func NewSite(url string, parser contracts.Parser) *Site {
	return &Site{url: url, parser: parser}
}

func (site *Site) parse() (res []RawData, err error) {
	client := http.Client{Timeout: time.Second * 5}
	resp, err := client.Get(site.url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	res = site.parser.ListSite(resp.Body)
}

type RawData struct {
	Data      time.Time
	Url       string
	Ip        string
	Who       string
	CountDown int64
}
