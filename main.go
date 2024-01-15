package main

import (
    "fmt" 
)

func main() {
    tree := RedBlackTree{}
    tree.Insert(20)
    tree.Insert(30)
    tree.Insert(40)
	tree.Insert(10)
    tree.Insert(15)


    fmt.Println("Red-Black Tree created with elements")
	tree.PrintInOrder()
}