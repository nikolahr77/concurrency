package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var wg sync.WaitGroup
var closed bool

//ReadFilesInDirectory reads all the files names
// in a given directory and sends them to FileContent
func ReadFilesInDirectory(ch chan string, dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, f := range files {
		//fmt.Println(f.Name())
		ch <- f.Name()
	}
	close(ch)
	return nil
}

//FileContent opens the files sent by
// ReadFilesInDirectory, displays the content and sends the files to
// DeleteFiles
func FileContent(ch chan string, send chan string, dir string) error {
	defer wg.Done()
	for i := range ch {
		b, err := ioutil.ReadFile(dir + "/" + i)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		send <- i
	}
	return nil
}

//DeleteFiles deletes the files that were already opened by FileContent
func DeleteFiles(ch chan string, dir string) error {
	defer wg.Done()
	for i := range ch {
		var err = os.Remove(dir + "/" + i)
		if err != nil {
			return err
		}
		fmt.Println("File Deleted")
	}
	return nil
}

func main() {
	ch := make(chan string)
	reciv := make(chan string)
	dir := "./readdelete/files"
	wg.Add(2)
	go ReadFilesInDirectory(ch, dir)
	go FileContent(ch, reciv, dir)
	go FileContent(ch, reciv, dir)
	go DeleteFiles(reciv, dir)
	wg.Wait()
	wg.Add(1)
	close(reciv)
	wg.Wait()
}
