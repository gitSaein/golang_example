package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"
)

const FILE_PATH = "./tmp/todoList.json"

type Todos struct {
	Todos   []Todo `json:"todos"`
	LastIdx int    `json:"last_idx"`
}

type Todo struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
}

func IndexOf(todos []Todo, targetIdx int) int {
	for i, todo := range todos {
		if todo.ID == targetIdx {
			return i
		}
	}
	return -1
}

// 생성
// 2.1. a or add 입력하고, 추가할 데이터를 입력하고 엔터하면 저장되고 리스트(입력된 시간 역순)를 보여준다.
// 2.2. !exit 입력하면, 취소되고 home으로 돌아간다.
func createTodo(data Todos) Todos {
	var cmdInput string
	for {

		fmt.Println("[add] add item ================================")
		fmt.Print("Add item> ")
		fmt.Scanln(&cmdInput)

		if cmdInput == "!exit" {
			return data
		}

		data.LastIdx += 1
		data.Todos = append(data.Todos, Todo{ID: data.LastIdx, Content: cmdInput, IsDone: false, CreatedAt: time.Now()})
		return data

	}
}

// 완료
// 3.  c or complate 입력하면,  리스트를 불러오고, index를 입력하면 완료 표시한다.
// 3.1. !exit 입력하면, 취소되고 home으로 돌아간다.
func completeTodo(data Todos) (Todos, error) {
	var cmdInput string
	for {

		getTodoList(data, "complete")
		fmt.Print("complate item > ")
		fmt.Scanln(&cmdInput)

		if cmdInput == "!exit" {
			return data, nil
		}

		idx, err := strconv.Atoi(cmdInput)
		if err != nil {
			return data, err
		}
		targetIdx := IndexOf(data.Todos, idx)
		if targetIdx < 0 {
			return data, errors.New("out of range")
		}
		data.Todos[targetIdx].IsDone = !data.Todos[targetIdx].IsDone
		return data, nil
	}

}

// 삭제
// 4.  d or delete 입력하면,  리스트를 불러오고, index를 입력하면 완료 표시한다.
// 4.1. !exit 입력하면, 취소되고 home으로 돌아간다.
func deleteTodo(data Todos) (Todos, error) {
	var cmdInput string

	for {

		getTodoList(data, "delete")
		fmt.Print("delete item > ")
		fmt.Scanln(&cmdInput)

		if cmdInput == "!exit" {
			return data, nil
		}
		cmdInputToInt, err := strconv.Atoi(cmdInput)
		if err != nil {
			return data, err
		}

		targetIdx := IndexOf(data.Todos, cmdInputToInt)
		data.Todos = append(data.Todos[:targetIdx], data.Todos[targetIdx+1:]...)
		return data, nil
	}

}

// 리스트 불러오기
func getTodoList(data Todos, pagename string) {
	fmt.Printf("[%s] Welcome to TODO list ================================", pagename)
	fmt.Println()

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

// 리스트 불러오기
func readFile() Todos {

	data := Todos{}

	byteValueFile, err := ioutil.ReadFile(FILE_PATH)
	if os.IsNotExist(err) {
		fmt.Println(err.Error())

		err := ioutil.WriteFile(FILE_PATH, []byte(""), 0644)
		if err != nil {
			fmt.Printf("err: %s", err.Error())
			return data
		}
	} else if err != nil {
		fmt.Println(err.Error())

	}

	json.Unmarshal([]byte(byteValueFile), &data)

	return data

}

// 파일 쓰기
func writeTodoListFile(todos Todos) error {
	jsonTodos, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(FILE_PATH, []byte(jsonTodos), 0644)
	if err != nil {
		return err
	}
	return nil
}

// 명령어 확인하기
func checkCmdInput(input string, todos Todos) {
	var errorByService error
	switch {
	case input == "a" || input == "add":
		todoList := createTodo(todos)
		err := writeTodoListFile(todoList)
		if err != nil {
			errorByService = err
			break
		}
	case input == "c" || input == "complete":
		todoList, err := completeTodo(todos)
		if err != nil {
			errorByService = err
			break
		}
		err = writeTodoListFile(todoList)
		if err != nil {
			errorByService = err
			break
		}
	case input == "d" || input == "delete":
		todoList, err := deleteTodo(todos)
		if err != nil {
			errorByService = err
			break
		}
		writeTodoListFile(todoList)
		err = writeTodoListFile(todoList)
		if err != nil {
			errorByService = err
			break
		}
	default:
		errorByService = errors.New("cmd 입력 값 오류")
	}

	if errorByService != nil {
		fmt.Printf("err: %s", errorByService.Error())
		fmt.Println()
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
		data := readFile()
		getTodoList(data, "home")
		fmt.Print("Enter> ")
		fmt.Scanln(&cmdInput)

		checkCmdInput(cmdInput, data)

	}

}
