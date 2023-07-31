package main

import (
	"container/list"
	"fmt"
)

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	for e := x.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
	/*The zero value for a List is an empty list (a *List can also be created using
	list.New). Values are appended to the list using PushBack. We loop over each item in
	the list by getting the first element, and following all the links until we reach nil.*/

}
