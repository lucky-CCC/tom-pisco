//go:generate goversioninfo -icon=favicon32.ico -manifest=main.manifest

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", sayBye)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v1 httpServer"))
}

// GOOS=windows CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags "-w -s -H windowsgui" -o one1.exe
