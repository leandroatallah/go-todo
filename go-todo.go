package main

import (
	"log"
	"os"
	"strconv"
)

const filePath = "./teste.csv"

func main() {
	data, err := ReadCsv()
	if err != nil {
		log.Fatalf("Houve um erro ao carregar os dados.")
	}

	list := ParseCsvToListItem(data)
	args := os.Args[1:]

	if len(args) == 0 {
		ListItems(list, false)
		return
	}

	switch args[0] {
	case CommandName[List]:
		showAll := len(args) > 1 && args[1] == "--all"
		ListItems(list, showAll)
	case CommandName[Add]:
		if len(args) < 2 {
			log.Fatalf("O comando add precisa de um conteúdo.")
		}
		AddItem(&list, args[1])
		ListItems(list, false)
	case CommandName[Delete]:
		if len(args) < 2 {
			log.Fatalf("O comando delete precisa de um index.")
		}

		index, ok := strconv.Atoi(args[1])
		if ok != nil || index < 1 {
			log.Fatalf("O index não é válido.")
		}

		DeleteItem(&list, index-1)
		ListItems(list, false)
	case CommandName[Complete]:
		if len(args) < 2 {
			log.Fatalf("O comando complete precisa de um index.")
		}

		index, ok := strconv.Atoi(args[1])
		if ok != nil || index < 1 {
			log.Fatalf("O index não é válido.")
		}
		CompleteItem(&list, index)
		ListItems(list, false)
	}
}
