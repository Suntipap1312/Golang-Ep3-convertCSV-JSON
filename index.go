package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Products struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}

func createList(data [][]string) []Products {
	// convert csv lines to array of structs
	var List []Products
	for i, line := range data {
		fmt.Println(i, line)
		if i > 0 { // omit header line
			var item Products // เตรียมภาชนะสำหรับบรรจุ
			for index, field := range line {
				if index == 0 {
					item.Name = field
				} else if index == 1 {
					var err error
					item.Quantity, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				} else if index == 2 {
					var err error
					item.Price, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				}
			}
			List = append(List, item)
		}
	}
	return List
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	List := createList(data)
	fmt.Printf("%+v\n", List)
	json, err := json.MarshalIndent(List, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}
