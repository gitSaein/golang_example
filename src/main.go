package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

const FILE_PATH = "./todoList.json"

type Todos struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

func sortByFieldName(todos Todos) {

}

func checkWantExit() {

}

func createTodo(data Todos) Todos {
	var cmdInput string
	var isExit = false

	for !isExit {

		fmt.Println("[add] add item ================================")
		fmt.Print("Add item> ")
		fmt.Scanln(&cmdInput)

		if cmdInput == "!exit" {
			return data
		}

		lenOfData := len(data.Todos)
		data.Todos = append(data.Todos, Todo{ID: lenOfData + 1, Content: cmdInput, IsDone: false, CreatedAt: time.Now()})
	}

	return data
}

func completeTodo(data Todos) Todos {
	fmt.Println("[complate] Welcome to TODO list ================================")

	return data

}

func deleteTodo(data Todos) Todos {
	fmt.Println("[delete] Welcome to TODO list ================================")

	return data

}

func getTodoList(data Todos) {
	fmt.Println("[home] Welcome to TODO list ================================")

	lenOfData := len(data.Todos)

	if lenOfData == 0 {
		fmt.Println("Empty Task ..")
		fmt.Println("Please add task.(a or add)")
	} else {
		sort.Slice(data.Todos, func(i, j int) bool {
			return data.Todos[i].CreatedAt.After(data.Todos[j].CreatedAt)
		})

		for i := 0; i < lenOfData; i++ {
			if data.Todos[i].IsDone {
				fmt.Printf("%d. (%s) ", data.Todos[i].ID, "V")
			} else {
				fmt.Printf("%d. (%s) ", data.Todos[i].ID, " ")
			}
			fmt.Print(data.Todos[i].Content)
			fmt.Println("")
		}
	}
	fmt.Println("===================================================")
}

func readTodoListFile() Todos {

	data := Todos{}

	byteValueFile, err := ioutil.ReadFile(FILE_PATH)
	if os.IsNotExist(err) {
		fmt.Println(err.Error())

		err := ioutil.WriteFile(FILE_PATH, []byte(""), 0644)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if err != nil {
		fmt.Println(err.Error())

	}

	json.Unmarshal([]byte(byteValueFile), &data)
	fmt.Println(data)

	return data

}

func writeTodoListFile(todos Todos) (Todos, error) {
	jsonTodos, err := json.Marshal(todos)
	if err != nil {
		return todos, err
	}
	err = ioutil.WriteFile(FILE_PATH, []byte(jsonTodos), 0644)

	return todos, err
}

func checkCmdInput(input string, todos Todos) {
	switch {
	case input == "a" || input == "add":
		todoList := createTodo(todos)
		writeTodoListFile(todoList)
	case input == "c" || input == "complete":
		todoList := completeTodo(todos)
		writeTodoListFile(todoList)
	case input == "d" || input == "delete":
		todoList := deleteTodo(todos)
		writeTodoListFile(todoList)
	default:
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
		data := readTodoListFile()
		getTodoList(data)
		// 2.1. a or add 입력하고, 추가할 데이터를 입력하고 엔터하면 저장되고 리스트(입력된 시간 역순)를 보여준다.
		fmt.Print("Enter> ")
		fmt.Scanln(&cmdInput)
		checkCmdInput(cmdInput, data)
		// 2.2. !exit 입력하면, 취소되고 home으로 돌아간다.
		// 3.  c or complate 입력하면,  리스트를 불러오고, index를 입력하면 완료 표시한다.
		// 3.1. !exit 입력하면, 취소되고 home으로 돌아간다.
		// 4.  d or delete 입력하면,  리스트를 불러오고, index를 입력하면 완료 표시한다.
		// 4.1. !exit 입력하면, 취소되고 home으로 돌아간다.
	}

}
