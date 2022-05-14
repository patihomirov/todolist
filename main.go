package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

func main() {
	database, _ := sql.Open("sqlite3", "./todobase.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, task TEXT, taskstatus TEXT, creationdate TEXT, closeddate TEXT)")
	statement.Exec()

	//Запись задачи в БД
	statement, _ = database.Prepare(`INSERT INTO tasks (task, taskstatus, creationdate ) VALUES (?, "opened", datetime())`)
	//Тестируем - создаем 10 задач
	for i := 0; i <= 10; i++ {
		statement.Exec("task №" + strconv.Itoa(i))
	}

	//Удаление завершенных задач
	statement, _ = database.Prepare("DELETE FROM tasks WHERE taskstatus == (?)")
	statement.Exec("done")

	//Вывод на печать задач
	rows, _ := database.Query("SELECT id, task, taskstatus, creationdate, closeddate FROM tasks")
	var id int
	var task string
	var taskstatus string
	var creationdate string
	var closeddate string

	for rows.Next() {
		rows.Scan(&id, &task, &taskstatus, &creationdate, &closeddate)
		fmt.Printf("%d: %s %s %s %s\n", id, task, taskstatus, creationdate, closeddate)
	}

}
