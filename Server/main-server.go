package main

import (
	"flag"
	"fmt"
	"net/http"
)

//---------------------------------------------------------------------------
// variables
//---------------------------------------------------------------------------
var port int
var help bool

func init() {
	flag.IntVar(&port, "port", 9000, "specify server execution port")
	flag.BoolVar(&help, "h", false, "command-line list")
	flag.BoolVar(&help, "help", false, "command-line list")
	flag.Parse()

	if help {
		flag.PrintDefaults()
	}
}

func startServer() {
	a := &App{
		OtherHandler: new(OtherHandler),
	}

	fmt.Println("Executando...")
	http.ListenAndServe(fmt.Sprintf(":%d", port), a)
}

func main() {
	startServer()
}
