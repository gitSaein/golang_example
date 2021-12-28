package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// var path string
	// fmt.Print("path:")
	// fmt.Scanln(&path)

	// files, err := ioutil.ReadDir(path)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, f := range files {
	// 	fmt.Println(f.Name())
	// }

	// var file string
	// fmt.Print("파일명:")
	// fmt.Scanln(&file)

	// bytes, err := ioutil.ReadFile(path + "/" + file)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("File contents: %s", bytes)
	// fmt.Println("")

	// files2, err := ioutil.ReadDir("./")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("==================================")

	// for _, f := range files2 {
	// 	fmt.Println(f.Name())
	// }
	// fmt.Println("==================================")

	var file2_nm string
	fmt.Print("파일명:")
	fmt.Scanln(&file2_nm)

	file2, err := os.Open(file2_nm)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	// 1. 미리 버퍼 사이즈를 알아야하는 경우
	// file2_info, err := file2.Stat()
	// if err != nil {
	// 	panic(err)
	// }
	// file2_size := file2_info.Size()
	// buffer := make([]byte, file2_size)

	// bytesread, err := file2.Read(buffer)
	// if err != nil {
	// 	panic(err)
	// }

	//2. 100 버퍼씩 읽는 경우
	BUFFER_SIZE := 100
	buffer := make([]byte, BUFFER_SIZE)

	for {
		bytesread, err := file2.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println("bytes read: ", bytesread)
		fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}

	fmt.Println("==================================")

	// 문자열 분리 기능
	// slice := strings.Split(string(buffer), ":")
	// for _, str := range slice {
	// 	fmt.Println(str)
	// }

}

func sliceFunc() {

}

func remoteOpenFunc() {

}
