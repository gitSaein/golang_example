// package main

// import (
// 	"bytes"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/jlaffaye/ftp"
// )

// func main() {
// 	http.Handle("/", new(testHandler))

// 	http.ListenAndServe(":5000", nil)
// }

// type testHandler struct {
// 	http.Handler
// }

// func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	str := "Your Request Path is " + req.URL.Path
// 	ftpService(req.URL.Path)
// 	w.Write([]byte(str))
// }

// func ftpService(path string) {
// 	log.Fatal("start ftp connected ...")
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
// 	// curDir, err := c.List("test")

// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// log.Fatal(curDir)

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