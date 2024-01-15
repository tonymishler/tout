package main

import (
    "fmt" 
)

type Color bool

const (
    RED   Color = true
    BLACK Color = false
)

type Node struct {
    Value  string 
    Color  Color
    Left   *Node
    Right  *Node
    Parent *Node
}

type RedBlackTree struct {
    Root *Node
}

func (n *Node) isRed() bool {
    if n == nil {
        return false
    }
    return n.Color == RED
}

func (t *RedBlackTree) rotateLeft(x *Node) {
    y := x.Right
    x.Right = y.Left
    if y.Left != nil {
        y.Left.Parent = x
    }
    y.Parent = x.Parent
    if x.Parent == nil {
        t.Root = y
    } else if x == x.Parent.Left {
        x.Parent.Left = y
    } else {
        x.Parent.Right = y
    }
    y.Left = x
    x.Parent = y
}

func (t *RedBlackTree) rotateRight(x *Node) {
    y := x.Left
    x.Left = y.Right
    if y.Right != nil {
        y.Right.Parent = x
    }
    y.Parent = x.Parent
    if x.Parent == nil {
        t.Root = y
    } else if x == x.Parent.Right {
        x.Parent.Right = y
    } else {
        x.Parent.Left = y
    }
    y.Right = x
    x.Parent = y
}

func (t *RedBlackTree) insertFixup(z *Node) {
    for z.Parent.isRed() {
        if z.Parent == z.Parent.Parent.Left {
            y := z.Parent.Parent.Right
            if y.isRed() {
                z.Parent.Color = BLACK
                y.Color = BLACK
                z.Parent.Parent.Color = RED
                z = z.Parent.Parent
            } else {
                if z == z.Parent.Right {
                    z = z.Parent
                    t.rotateLeft(z)
                }
                z.Parent.Color = BLACK
                z.Parent.Parent.Color = RED
                t.rotateRight(z.Parent.Parent)
            }
        } else {
            y := z.Parent.Parent.Left
            if y.isRed() {
                z.Parent.Color = BLACK
                y.Color = BLACK
                z.Parent.Parent.Color = RED
                z = z.Parent.Parent
            } else {
                if z == z.Parent.Left {
                    z = z.Parent
                    t.rotateRight(z)
                }
                z.Parent.Color = BLACK
                z.Parent.Parent.Color = RED
                t.rotateLeft(z.Parent.Parent)
            }
        }
    }
    t.Root.Color = BLACK
}

func (t *RedBlackTree) Insert(value string) {
    newNode := &Node{Value: value, Color: RED}
    if t.Root == nil {
        t.Root = newNode
    } else {
        currentNode := t.Root
        var parentNode *Node

        for currentNode != nil {
            parentNode = currentNode
            if newNode.Value < currentNode.Value {
                currentNode = currentNode.Left
            } else {
                currentNode = currentNode.Right
            }
        }

        newNode.Parent = parentNode

        if newNode.Value < parentNode.Value {
            parentNode.Left = newNode
        } else {
            parentNode.Right = newNode
        }
    }

    t.insertFixup(newNode)
}

func (t *RedBlackTree) inOrderTraversal(node *Node) {
    if node != nil {
        t.inOrderTraversal(node.Left)
        fmt.Printf("%d ", node.Value)
        t.inOrderTraversal(node.Right)
    }
}

func (t *RedBlackTree) PrintInOrder() {
    t.inOrderTraversal(t.Root)
    fmt.Println()
}

