package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type activity struct{
	Name string
	Start string
	End string
}

func parseConfig()  {
	file, _ := os.Open("./conf/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file);
	conf := activity{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(conf.Name);
}

func jsonEncode()  {
	type Rgb struct {
		R, G, B int
	}

	type ColorGroup struct {
		ID     int `json:",string"`
		Name   string
		Colors []string
		Rgb `json:"rgb"`
	}
	group := ColorGroup{
		1,
		"Reds",
		[]string{"Crimson", "Red", "Ruby", "Maroon"},
		Rgb{50, 50, 50},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}