package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	list = flags.Bool("l", false, "List the stored jumps")
	add  = flags.String("a", "", "Name of list item to add")
	del = flags.String("d", "", "Name of list item to delete")
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flags.Parse()
	file, err := os.Open("~/.jumprc")
	check(err)

	links := make(map[string]string)
	check(json.NewDecoder(file).Decode(&links))
	check(file.Close())

	if *list {
		for key, value := range links {
			fmt.Printf("\t%s:\t%s\n", key, value)
		}
	}

	if *add != "" {
		links[*add] = os.Args[1]
		fmt.Print("TODO: actually store file")
	}

}
