package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var path string
	fmt.Print("path:")
	fmt.Scanln(&path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	var file string
	fmt.Print("파일명:")
	fmt.Scanln(&file)

	bytes, err := ioutil.ReadFile(path + "/" + file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File contents: %s", bytes)

	files2, err := ioutil.ReadDir("./")
	if err != nil {
		panic(err)
	}

	for _, f := range files2 {
		fmt.Println(f.Name())
	}

	var file2_nm string
	fmt.Print("파일명:")
	fmt.Scanln(&file2_nm)

	file2, err := os.Open(file2_nm)
	if err != nil {
		panic(err)
	}

	file2_info, err := file2.Stat()
	if err != nil {
		panic(err)
	}

	file2_size := file2_info.Size()
	buffer := make([]byte, file2_size)

	bytesread, err := file2.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("bytes read: ", bytesread)
	fmt.Println("bytestream to string: ", string(buffer))

	defer file2.Close()

}
