package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/souravdey425/cyoa"
)

func main() {
	fileName := flag.String("file", "gopher.json", "its contain a json file")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		fmt.Println(err)
	}
	d := json.NewDecoder(f)
	var story cyoa.Story
	d.Decode(&story)
	fmt.Printf("%+v", story)
}
