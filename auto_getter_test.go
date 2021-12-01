package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDecorator(t *testing.T) {
	var c ProxyGetter
	c = WrapWithTimeDecorator(c, 360)
	c = WrapWithThresholdDecorator(c, 80)
	go func() {
		for {
			// for c.LenOfProxies() < 90{

			// }
			p, _ := c.GetProxy()
			fmt.Printf("c.EraseProxy(p): %v\n", c.EraseProxy(p))
			time.Sleep(time.Millisecond * 500)
		}
	}()
	select {}
}
