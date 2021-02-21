package main

import "fmt"

//Interfaces: a set of methods and a type used to define the behaviour of a type

type coordinates interface {
	xaxis() int
	yaxis() int
}

type point2D struct {
	X int
	Y int
}

func (s point2D) xaxis() int{
	return s.X
}

func (s point2D) yaxis() int  {
	return s.Y
}

func findCoordinates(a coordinates)  {
	fmt.Println("X:", a.xaxis(), "Y:", a.yaxis())
}

type coordinate int

func (s coordinate)  xaxis() int{
	return int(s)

}

func (s coordinate) yaxis() int {
	return 0
}