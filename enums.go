package main

type CommandIndex int

const (
	List CommandIndex = iota
	Add
	Delete
	Complete
)

var CommandName = map[CommandIndex]string{
	List:     "list",
	Add:      "add",
	Delete:   "delete",
	Complete: "complete",
}
