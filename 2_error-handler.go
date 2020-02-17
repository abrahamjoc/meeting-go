package main

import (
	"encoding/json"
	"fmt"
	"github.com/kr/pretty"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("json file error:", err)
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read file error:", err)
		os.Exit(1)
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("json marshalling error:", err)
		os.Exit(1)
	}

	_, _ = pretty.Println(jsonData)
}
