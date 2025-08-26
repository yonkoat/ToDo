package main

func main() {
	todos := Todos{}
	Storage := NewStorage[Todos]("todos.json")
	Storage.Load(&todos)
	CommandFlags := NewCommandsFlags()
	CommandFlags.Execute(&todos)
	Storage.Save(todos)

}
