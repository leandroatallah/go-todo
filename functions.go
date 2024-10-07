package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"text/tabwriter"
)

func ListItems(list []ListItem, showAll bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)
	if !showAll {
		// Remove all completed items
		for index, item := range list {
			if item.isCompleted {
				list = slices.Delete(list, index, index+1)
			}
		}
	}

	if len(list) == 0 {
		fmt.Println("Não há itens pendentes.")
		return
	}

	header := "ID\tTask"
	if showAll {
		header = "ID\tTask\tDone"
	}
	fmt.Fprintln(w, header)

	for _, item := range list {
		var row string
		if showAll {
			row = fmt.Sprintf("%d\t%v\t%v", item.id, item.content, item.isCompleted)
		} else {
			row = fmt.Sprintf("%d\t%v", item.id, item.content)
		}
		fmt.Fprintln(w, row)
	}

	w.Flush()
}

func AddItem(list *[]ListItem, item string) {
	id := len(*list) + 1
	newItem := ListItem{id, item, false}
	*list = append(*list, newItem)
	WriteCsv(*list)
}

func DeleteItem(list *[]ListItem, index int) {
	nlist := *list
	*list = append(nlist[:index], nlist[index+1:]...)
	WriteCsv(*list)
}

func CompleteItem(list *[]ListItem, index int) {
	item := slices.IndexFunc(*list, func(item ListItem) bool {
		return item.id == index
	})

	if item == -1 {
		log.Fatalf("Index inválido.")
	}

	nlist := *list
	updatedItem := nlist[item]
	updatedItem.isCompleted = true
	nlist[item] = updatedItem
	*list = nlist

	WriteCsv(*list)
}
