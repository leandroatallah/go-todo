package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func ReadCsv() ([]byte, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		os.Create(filePath)
	}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func ParseCsvToListItem(data []byte) []ListItem {
	// Parse
	reader := csv.NewReader(bytes.NewReader(data))

	// Process
	result := []ListItem{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Houve um erro a carregar os dados.", err)
			break
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal("Houve um erro a carregar os dados.", err)
		}
		content := record[1]
		isCompleted, err := strconv.ParseBool(record[2])
		if err != nil {
			log.Fatal("Houve um erro a carregar os dados.", err)
		}

		result = append(result, ListItem{id, content, isCompleted})
	}

	return result
}

func WriteCsv(list []ListItem) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w := csv.NewWriter(f)

	records := make([][]string, len(list))

	for index, item := range list {
		id := strconv.Itoa(item.id)
		isCompleted := strconv.FormatBool(item.isCompleted)
		records[index] = []string{id, item.content, isCompleted}
	}

	for _, record := range records {
		w.Write(record)
	}

	w.Flush()
}
