package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Response struct {
	contents   string
	is         bool
	error_code int
}

type File struct {
	name string
	path string
}

func main() {

	var path string

	isPath := false
	for isPath == false {
		fmt.Print("파일이 있는 폴더 경로를 입력해주세요.:")
		fmt.Scanln(&path)

		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Print("경로를 잘못 입력하였습니다. 다시 입력해주세요.")
			continue
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}
		break

	}

	isfile := false
	for isfile == false {

		var file string
		fmt.Print("파일명:")
		fmt.Scanln(&file)

		response := readFileBySize(path, file)

		if !response.is {
			fmt.Print("파일이 없습니다. 파일명을 다시 입력해주세요.")
			continue
		}
		fmt.Println("===== file contents =====")
		fmt.Printf("%s", response.contents)
		fmt.Println("=========================")
		break
	}

}

func readFileBySize(path string, filename string) Response {

	file, err := os.Open(path + "/" + filename)
	if err != nil {
		log.Fatal(err)
		return Response{"file not found", false, 404}
	}

	// 1. 미리 버퍼 사이즈를 알아야하는 경우
	file_info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	file2_size := file_info.Size()
	buffer := make([]byte, file2_size)

	bytesread, err := file.Read(buffer)
	if err != nil {
		panic(err)
	}

	return Response{string(buffer[:bytesread]), true, 200}

}
func splitString() {
	// slice := strings.Split(string(buffer), ":")
	// for _, str := range slice {
	// 	fmt.Println(str)
	// }
}
