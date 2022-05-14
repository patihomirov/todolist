package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

var database, _ = sql.Open("sqlite3", "./todobase.db")

func main() {
	//Инициализация БД
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, task TEXT, taskstatus TEXT, creationdate TEXT, closeddate TEXT)")
	statement.Exec()

	//Обработка внешних запросов
	router := mux.NewRouter()
	router.HandleFunc("/", sHead)
	router.HandleFunc("/create/{taskname}", sCreate)
	router.HandleFunc("/changestatus/{taskid}/{taskstatus}", sStatusChange)
	router.HandleFunc("/taskPrint", sPrint)
	router.HandleFunc("/clearall", sClearAll)
	router.HandleFunc("/cleardone", sClearDone)
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
	vars := mux.Vars(r)
	taskname := vars["taskname"]
	statement, _ := database.Prepare(`INSERT INTO tasks (task, taskstatus, creationdate, closeddate ) VALUES (?, "opened", ?, "")`)
	statement.Exec(taskname, time.Now().String())
	w.Write([]byte(`Tasks "` + taskname + `" Created`))
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
	printbuffer = fmt.Sprintf("ID \t Задача: \t Статуc: \t Дата открытия: \t Дата закрытия: \n")
	for rows.Next() {
		rows.Scan(&id, &task, &taskstatus, &creationdate, &closeddate)
		printbuffer = printbuffer + fmt.Sprintf("%d:\t %s \t %s \t %s \t %s \n", id, task, taskstatus, creationdate, closeddate)
	}
	return printbuffer
}

//Изменение статуса задачи
func sStatusChange(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskid := vars["taskid"]
	taskstatus := vars["taskstatus"]
	var taskclosedate string
	if taskstatus == "done" {
		taskclosedate = time.Now().String()
	} else {
		taskclosedate = ""
	}
	statement, _ := database.Prepare(`UPDATE tasks SET taskstatus = ?, closeddate = ? WHERE id = ?`)
	statement.Exec(taskstatus, taskclosedate, taskid)
	w.Write([]byte(`Tasks "` + taskid + `" change status to ` + taskstatus))
}

//Удаление из БД завершенных задач
func sClearDone(w http.ResponseWriter, r *http.Request) {
	statement, _ := database.Prepare(`DELETE FROM tasks WHERE taskstatus = ?`)
	statement.Exec("done")
	w.Write([]byte("Clear doned tasks done!"))
}

//Удаление из БД всех записей
func sClearAll(w http.ResponseWriter, r *http.Request) {
	statement, _ := database.Prepare(`DELETE FROM tasks`)
	statement.Exec()
	w.Write([]byte("Clear all tasks done!"))
}
