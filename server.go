package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

func readFile() (file []byte) {
	indexFileName := "./index.html"

	flag.Parse()

	args := flag.Args()

	log.Println(args)

	if len(args) > 0 {
		indexFileName = args[0]
	}

	log.Printf("index file loaded: %s", indexFileName)

	f, e := ioutil.ReadFile(indexFileName)

	if e != nil {
		log.Fatal(e.Error())
	}

	return f
}

func handler(file []byte) func(*fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("text/html")
		ctx.SetBody(file)
	}
}

func main() {

	port, _ := os.LookupEnv("PORT")

	if port == "" {
		port = "8080"
	}

	indexFile := readFile()

	log.Printf("server started on: %s", port)

	log.Fatal(fasthttp.ListenAndServe(":"+port, handler(indexFile)))
}
