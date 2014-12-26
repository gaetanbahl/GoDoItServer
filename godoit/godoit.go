package godoit

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const todoFile string = "./godoit.db"

type entry struct {
	id      int
	state   string
	created string
	name    string
}

func listItems() []entry {

	var itemList []entry
	text := readFile()
	lines := strings.Split(text, "\n")

	for _, elem := range lines[:len(lines)-1] {
		fmt.Println(elem)
		slicedElem := strings.Split(elem, "|")
		id, _ := strconv.Atoi(slicedElem[0])
		I := entry{id, slicedElem[1], slicedElem[2], slicedElem[3]}
		fmt.Println(I.id)
		itemList = append(itemList, I)
	}

	return itemList
}

func createItem(state string, name string) {

	datetime := time.Now()
	Year, Month, Day := datetime.Date()
	date := strings.Join([]string{string(Day), string(Month), string(Year)}, " ")

	var listItem []entry
	listItem = listItems()
	I := entry{len(listItem), state, date, name}
	listItem = append(listItem, I)

	writeFile(listItem)

}

func deleteItem(id int) {

	itemList := listItems()

	var newList []entry

	for _, item := range itemList {
		if item.id < id {
			newList = append(newList, item)
		} else if item.id > id {
			newList = append(newList, entry{item.id - 1, item.state, item.created, item.name})
		}
	}

	writeFile(newList)

}

func markItemAs(id int, state string) {

	itemList := listItems()

	itemList[id] = entry{id, state, itemList[id].created, itemList[id].name}

}

func Initialize() {

}

func readFile() string {
	bs, err := ioutil.ReadFile(todoFile)

	if err != nil {
		panic(err)
	}
	return string(bs)
}

func stringFromItems(itemList []entry, sep string) string {

	s := string("")
	var lines []string

	for _, elem := range itemList {
		s = strings.Join([]string{string(elem.id), elem.state, elem.created, elem.name}, sep)
		lines = append(lines, s)
	}

	return strings.Join(lines, "\n")

}

func writeFile(itemList []entry) {

	f, err := os.Create(todoFile)

	check(err)
	defer f.Close()
	f.WriteString(stringFromItems(itemList, "|"))
	f.Sync()

}

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func Retrieve() string {

	return stringFromItems(listItems(), " ")

}
