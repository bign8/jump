package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
)

var (
	target string
	goals  []string
)

func main() {
	path := filepath.Join(build.Default.GOPATH, "src")
	target = os.Args[len(os.Args)-1]
	if err := filepath.Walk(path, walk); err != nil {
		panic(err)
	}
	// TODO: get shortest in list or present new UI
	fmt.Println(goals[0])
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil || !info.IsDir() {
		return err
	}
	base := filepath.Base(path)
	if base == "testdata" || base == "internal" {
		return filepath.SkipDir
	}
	if base == target {
		goals = append(goals, path)
	}
	return nil
}
