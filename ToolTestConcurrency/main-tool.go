package main

import (
	"flag"
	"fmt"
	"sync"
)

//---------------------------------------------------------------------------
// variables
//---------------------------------------------------------------------------
var server string
var port int
var requests int
var help bool
var wg sync.WaitGroup

func init() {
	flag.StringVar(&server, "server", "http://127.0.0.1", "specify server address")
	flag.IntVar(&port, "port", 9000, "specify server port")
	flag.IntVar(&requests, "r", 1000, "specify requisition numbers")
	flag.BoolVar(&help, "h", false, "command-line list")
	flag.BoolVar(&help, "help", false, "command-line list")
	flag.Parse()

	if help {
		flag.PrintDefaults()
	}
}

func main() {
	fmt.Println("Iniciando...")
	for i := 0; i < requests; i++ {
		wg.Add(1)
		go setupRequest()
	}
	wg.Wait()
	fmt.Println("Finalizado!")
	fmt.Scanln()
}

func setupRequest() {
	var url string = fmt.Sprintf("%s:%d/other/hello", server, port)
	api := new(ApiUtil)
	fmt.Println(api.apiGetRequest(url))
	wg.Done()
}
