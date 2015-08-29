package main

import "fmt"

import "./core"

func main() {
	fmt.Println("Game goes here")

	fmt.Println(core.NewGame().Positions())
}
