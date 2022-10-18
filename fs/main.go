package main

import (
	"embed"
	"fmt"
)

//go:embed resources
var fs embed.FS

func main() {
	fmt.Printf("%+v\n", fs)
	dir, err := fs.ReadDir("resources/systems")
	if err != nil {
		panic(err)
	}
	for i, e := range dir {
		fmt.Println(i, e.Name())
	}
}

// GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags "-w -s" -o fs_test.exe
