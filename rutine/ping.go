package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"time"
)

func ScanPort(protocol, hostname string, port int) (bool, error) {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, time.Millisecond*100)
	if err != nil {
		return false, err
	}
	defer conn.Close()
	return true, nil
}

func main() {
	for i := 1; i <= runtime.NumCPU(); i++ {
		go func(min, max int) {
			for min <= max {
				if r, _ := ScanPort("ip4:1", "108.177.16.0/24", min); r == true {
					fmt.Println("localhost: ", min, " is open")
				}
				min++
			}
		}
	}()
}
