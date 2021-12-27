package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(*testing.T) {
	t1 := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-t1.C:
			fmt.Println(time.Now().Unix())
			t1.Reset(time.Second * 2)
		}

	}
}
