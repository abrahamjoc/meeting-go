package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	run := func (name string, s time.Duration) {
		time.Sleep(s * time.Second)
		fmt.Println(time.Now().String()[:19], name, "executed")
	}

	go run("T0", 0)
	go run("G1", 3)
	go run("G2", 2)

	bufio.NewScanner(os.Stdin).Scan()
}
