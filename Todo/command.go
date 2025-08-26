package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandFlags struct {
	Add  string
	Del  int
	Edit string
	Togg int
	List bool
}

func NewCommandsFlags() *CommandFlags {
	cf := CommandFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new chore to the list")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a chore from the list via index. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete a chore from the list by the index")
	flag.IntVar(&cf.Togg, "togg", -1, "Change the status of completion of the chore via index")
	flag.BoolVar(&cf.List, "ls", false, "List all the chores")

	flag.Parse()

	return &cf
}

func (cf *CommandFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])

	case cf.Togg != -1:
		todos.toggle(cf.Togg)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")
	}
}
