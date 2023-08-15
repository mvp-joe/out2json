package main

import (
	"encoding/json"
	"fmt"
	"log"
	"out2json"
)

func main() {
	input := out2json.ReadStdin()
	root, err := out2json.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	jsonText, err := json.Marshal(root)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonText))
}
