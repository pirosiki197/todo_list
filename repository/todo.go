package repository

type Todo struct {
	ID     int
	Task   string
	Status string
}

func (repo *Repository) CreateTodo(todo Todo) (id int, err error) {
	result, err := repo.db.Exec("INSERT INTO `todos` (`task`, `status`) VALUES (?, ?)", todo.Task, todo.Status)
	if err != nil {
		return 0, err
	}
	idi64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(idi64), nil
}

func (repo *Repository) GetTodo(id int) (Todo, error) {
	row := repo.db.QueryRow("SELECT `id`, `task`, `status` FROM `todos`")
	var todo Todo
	err := row.Scan(&todo.ID, &todo.Task, &todo.Status)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (repo *Repository) UpdateTodo(id int, todo Todo) error {
	_, err := repo.db.Exec("UPDATE `todos` SET `task` = ?, `status` = ? WHERE `id` = ?", todo.Task, todo.Status, id)
	return err
}
