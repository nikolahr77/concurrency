package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		root  string
		files []string
		err   error
	)

	root := "/home/manigandan/golang/samples"
	// filepath.Walk
	//files, err = FilePathWalkDir(root)
	//if err != nil {
	//	panic(err)
	//}
	// ioutil.ReadDir
	files, err = IOReadDir(root)
	if err != nil {
		panic(err)
	}
	//os.File.Readdir
	files, err = OSReadDir(root)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}