// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"time"

// 	"github.com/jlaffaye/ftp"
// )

// func main() {
// 	//폴더 조회 및 생성
// 	// var folder string
// 	// var fileName string
// 	// var fileContent string

// 	//FTP 연결
// 	c, err := ftp.Dial("127.0.0.1:21", ftp.DialWithTimeout(5*time.Second))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//FTP 서버 인증
// 	err = c.Login("tester", "tester")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = c.MakeDir("test3")

// 	if err != nil {
// 		panic(err)
// 	}

// 	list, err := c.NameList("test3")

// 	for _, line := range list {
// 		fmt.Println(line)
// 	}

// 	// 버퍼를 생성해서 파일 저장
// 	data := bytes.NewBufferString("Hello World")
// 	err = c.Stor("test-file.txt", data)
// 	if err != nil {
// 		log.Fatal(err)
// 		panic(err)
// 	}

// 	// 파일 읽어오기
// 	r, err := c.Retr("test-file.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 		panic(err)
// 	}
// 	defer r.Close()

// 	buf, err := ioutil.ReadAll(r)
// 	println(string(buf))

// 	//FTP 연결 해제
// 	if err := c.Quit(); err != nil {
// 		log.Fatal(err)
// 	}
// }
