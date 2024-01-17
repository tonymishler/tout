package main

import (
	"fmt"
)

func main() {
	tree := RedBlackTree{}
	tree.Insert("fizz")
	tree.Insert("buzz")
	tree.Insert("alpha")
	tree.Insert("bravo")
	tree.Insert("charlie")

	fmt.Println("Red-Black Tree created with elements")
	tree.PrintInOrder()
	tree.PrintTree()
}
