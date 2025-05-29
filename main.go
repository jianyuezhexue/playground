package main

import (
	"fmt"
	"time"
)

func main() {
	lastMaxCreateTime := time.Now().Format("2006-01-02 00:00:00")
	fmt.Println(lastMaxCreateTime)
}
