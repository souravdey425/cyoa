package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/souravdey425/cyoa"
)

func main() {
	port := flag.Int("port", 3000, "port to start the server")
	fileName := flag.String("file", "gopher.json", "its contain a json file")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
	}
	story, err := cyoa.JsonStory(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", story)
	h := cyoa.NewHandler(story)
	fmt.Printf("starting the server the port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
