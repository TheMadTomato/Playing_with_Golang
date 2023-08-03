package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	/*file, err := os.Open("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)*/

	/*We use defer file.Close() right after opening the file to make sure the file is closed
	as soon as the function completes. Reading files is very common, so there’s a shorter
	way to do this:
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)*/

	/*To create a file, use the os.Create function. It takes the name of the file, creates it in
	the current working directory, and returns an os.File and possibly an error (if it was
	unable to create it for some reason). Here’s an example program:

	file, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString("test")*/

	/*To get the contents of a directory, we use the same os.Open function but give it a
	directory path instead of a file name. Then we call the Readdir method:

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}*/

	/*Often, we want to recursively walk a folder (read the folder’s contents, all the subfolders,
	all the sub-subfolders, etc.). To make this easier, there’s a Walk function provided
	in the path/filepath package:*/
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
