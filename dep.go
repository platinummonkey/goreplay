package goreplay

import (
	"io"
)

var CloseCh chan int = make(chan int)

func Finalize() {
	for _, p := range Plugins.All {
		if cp, ok := p.(io.Closer); ok {
			cp.Close()
		}
	}
}
