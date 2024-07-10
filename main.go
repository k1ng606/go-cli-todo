package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Description string
	Done        bool
}

type ItemCollection struct {
	Items []Item
}

func NewItemCollection() *ItemCollection {
	return &ItemCollection{
		Items: []Item{},
	}
}

func (ic *ItemCollection) PrintAllItems() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("description|done")
	for i, item := range ic.Items {
		fmt.Printf("%d. %s|%t\n", i+1, strings.ReplaceAll(item.Description, "\n", ""), item.Done)
	}
	fmt.Println("---------------------------------------------------")
}

func (ic *ItemCollection) MarkItemDone(index int) {
	if index > 0 && index <= len(ic.Items) {
		ic.Items[index-1].Done = true
	} else {
		fmt.Println("Invalid index!")
	}
}

func (ic *ItemCollection) AddItem(item Item) {
	ic.Items = append(ic.Items, item)
}

func NewItem(description string) Item {
	return Item{
		Description: description,
		Done:        false,
	}
}

func main() {
	itemCollection := NewItemCollection()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		if parts[0] == "complete" {
			if len(parts) != 2 {
				fmt.Println("Invalid input!")
			} else {
				index, err := strconv.Atoi(parts[1])
				if err != nil {
					fmt.Println("Invalid index!")
				} else {
					itemCollection.MarkItemDone(index)
				}
			}
		} else {
			newItem := NewItem(input)
			itemCollection.AddItem(newItem)
		}
		itemCollection.PrintAllItems()
	}
}
