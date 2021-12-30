package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Todos struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

func createTodo() {
	fmt.Println("[add] add item ================================")

}

func completeTodo() {
	fmt.Println("[complate] Welcome to TODO list ================================")

}

func deleteTodo() {
	fmt.Println("[delete] Welcome to TODO list ================================")

}

func getTodoList(data Todos) {
	fmt.Println("[home] Welcome to TODO list ================================")

	lenOfData := len(data.Todos)

	if lenOfData == 0 {
		fmt.Println("Empty Task ..")
		fmt.Println("Please add task.(a or add)")
	} else {

		for i := 0; i < lenOfData; i++ {
			if data.Todos[i].IsDone {
				fmt.Printf("%d. (%s) ", i+1, "V")
			} else {
				fmt.Printf("%d. (%s) ", i+1, " ")
			}
			fmt.Print(data.Todos[i].Content)
		}
	}
	fmt.Println("===================================================")
}

func openTodoListFile() Todos {

	data := Todos{}
	filepath := "./todoList.json"

	// jsonFile, err := os.Open(filepath)
	byteValueFile, err := ioutil.ReadFile(filepath)
	if os.IsNotExist(err) {
		fmt.Println(err.Error())

		err := ioutil.WriteFile(filepath, []byte(""), 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if err != nil {
		fmt.Println(err.Error())

	}

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValueFile), &data)
	fmt.Println(data)

	return data

}

func checkCmdInput(input string) {
	switch input {
	case "a":
	case "add":
		createTodo()
		break
	case "c":
	case "complete":
		completeTodo()
		break
	case "d":
	case "delete":
		deleteTodo()
		break
	default:
		break
	}
}

func main() {
	var isEnd bool
	// var isGoHome bool
	isEnd = true
	// isGoHome = false

	// 1. todo list가 있는지 없는지 확인해서 있으면 리스트 불러오기
	for isEnd {
		var cmdInput string
		data := openTodoListFile()
		getTodoList(data)
		// 2.1. a or add 입력하고, 추가할 데이터를 입력하고 엔터하면 저장되고 리스트(입력된 시간 역순)를 보여준다.
		fmt.Print("Enter> ")
		fmt.Scanln(&cmdInput)
		checkCmdInput(cmdInput)
		// 2.2. !exit 입력하면, 취소되고 home으로 돌아간다.
		// 3.  c or complate 입력하면,  리스트를 불러오고, index를 입력하면 완료 표시한다.
		// 3.1. !exit 입력하면, 취소되고 home으로 돌아간다.
		// 4.  d or delete 입력하면,  리스트를 불러오고, index를 입력하면 완료 표시한다.
		// 4.1. !exit 입력하면, 취소되고 home으로 돌아간다.
	}

}
