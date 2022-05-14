package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

var database, _ = sql.Open("sqlite3", "./todobase.db")

func main() {
	//Инициализация БД

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, task TEXT, taskstatus TEXT, creationdate TEXT, closeddate TEXT)")
	statement.Exec()

	//Обработка внешних запросов
	router := mux.NewRouter()
	router.HandleFunc("/", sHead)
	router.HandleFunc("/create", sCreate)
	router.HandleFunc("/taskPrint", sPrint)
	//	router.HandleFunc("/clear",sClear)
	//	router.HandleFunc("/list",sList)
	//	router.HandleFunc("/feedback/{feedback}",sFeedback)
	log.Fatal(http.ListenAndServe(":3000", router))

	//Удаление завершенных задач
	//statement, _ = database.Prepare("DELETE FROM tasks WHERE taskstatus == (?)")
	//statement.Exec("done")
}

func sHead(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Добро пожаловать!`))
}

//Создание задач
func sCreate(w http.ResponseWriter, r *http.Request) {
	statement, _ := database.Prepare(`INSERT INTO tasks (task, taskstatus, creationdate ) VALUES (?, "opened", datetime())`)
	//Тестируем - создаем 10 задач
	for i := 1; i <= 10; i++ {
		statement.Exec("task №" + strconv.Itoa(i))
	}
	w.Write([]byte("Tasks Created"))
}

func sPrint(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(taskPrint()))
}

//Вывод на печать задач
func taskPrint() string {
	var printbuffer string
	rows, _ := database.Query("SELECT id, task, taskstatus, creationdate, closeddate FROM tasks")
	var id int
	var task string
	var taskstatus string
	var creationdate string
	var closeddate string
	for rows.Next() {
		rows.Scan(&id, &task, &taskstatus, &creationdate, &closeddate)
		//printbuffer = append(printbuffer, fmt.Sprintf("%d: %s %s %s %s\n", id, task, taskstatus, creationdate, closeddate))
		printbuffer = printbuffer + fmt.Sprintf("%d: %s %s %s %s\n", id, task, taskstatus, creationdate, closeddate)
	}
	return printbuffer
	//w.Write([]byte("Tasks Created"))
}
