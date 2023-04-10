package filter

import (
	"fmt"
	"time"
)

var FilterChan = make(chan map[string]interface{}, 10000)

func Start(num int) {
	for {
		select {
		case s := <-FilterChan:
			fmt.Println(num, s)
		default:
			time.Sleep(time.Second)
		}
	}
}
