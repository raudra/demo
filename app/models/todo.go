package model

type ToDo struct{
	Id  int `json: "id"`
	Name string `json:"name"`
}

func New(id int, val string) ToDo{
	return ToDo{Id: id,  Name: val}
}

func(todo ToDo) GetName() string{
	return todo.Name
}

