package data

type Todo struct {
	ID string `json:"id"`
	Item string  `json:"item"`
	Completed bool `json:"completed"`
}

var Todos = []Todo{
	{ ID: "1", Item: "Wash", Completed: false },
	{ ID: "2", Item: "Code", Completed: false },
	{ ID: "3", Item: "Read", Completed: false },
}
