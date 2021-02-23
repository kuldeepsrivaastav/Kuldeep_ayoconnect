package main

import (
	Binary_Tree_Assignment "Kuldeep-GoLangCodingExcersie-GolangSostwareEngineer/Binary-Tree-Assignment"
	House_Robber "Kuldeep-GoLangCodingExcersie-GolangSostwareEngineer/House-Robber"
	"fmt"
)


func main()  {

	fmt.Println("Binary Search Tree")
	var t Binary_Tree_Assignment.Tree

	t.Insert(10)
	t.Insert(20)
	t.Insert(40)
	t.Insert(60)
	t.Insert(30)
	t.Insert(50)

	fmt.Printf("In Order: ")
	Binary_Tree_Assignment.PrintInOrder(t.Root)
	fmt.Println()
	fmt.Printf("Pre Order: ")
	Binary_Tree_Assignment.PrintPreOrder(t.Root)
	fmt.Println()
	fmt.Printf("Post Order: ")
	Binary_Tree_Assignment.PrintPostOrder(t.Root)
	fmt.Println()

	fmt.Println("---------------------------------------------")
	fmt.Println("House Robber")
	slc:=[]int{1,2,3,1}
	fmt.Println(House_Robber.Rob(slc))
}