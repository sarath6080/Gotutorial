package main

import "fmt"

func main(){
	var a int //variable declaration
	fmt.Println("the age is :",a)
	a = 29
	fmt.Println("the age is :",a) 
	a=56
	fmt.Println("the age is :",a) 

	// type inference
	var age = 29
	fmt.Println("the age is :",age)

	// Multiple var declartion
	var weight,height = 100,50
	fmt.Println("the weight is", weight ,"the height is", height)

	var (
	age1 = 90
	name = "sarath"
	height2 int
	)
	fmt.Println("the name is",name , "the age is ",age1 ,"and height", height2)

	//short ahnd description
	name1,agei := "nav",40
	fmt.Println("the name ",name1 ,"the age is",agei)

}
