todolist - проект простой программы планировщика дел.
Программа разрабатывается как учебный проект для отработки навыков:
- работы с базой данных;
- работы с http-запросами и созданием простого web-интерфейса.
Программа находится в стадии разработки основного функционала.

Реализованные запросы:
http://localhost:3000/ - экран приветствия
http://localhost:3000/create/имя задачи - вызывает создание задачи
http://localhost:3000/changestatus/id задачи/статус задачи - изменяет статус задачи, можно установить любой
http://localhost:3000/taskPrint - вызывает вывод перечня всех задач
http://localhost:3000/cleardone - вызывает удаление задач со статусом "done"
http://localhost:3000/clearall - вызывает удаления всех задач