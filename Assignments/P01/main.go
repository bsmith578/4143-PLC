/*****************************************************************************
*  Author:           Brett Smith
*  Email:            brettsmith578@gmail.com
*  Label:            P01
*  Title:            Program 1 - Run a Go Program
*  Course:           CMPS 4143 - PLC
*  Semester:         Fall 2023
*
*  Description:
*        This program shows an example of implementing a Go Package called
*        "mascot". The mascot package contains a function "BestMascot" that
*        prints "Go Gopher" to the console. the function is called in a test
*        file, "mascot_test.go" and a main file "main.go". The test file is used
*        to make sure the functions are properly working before using them in
*        actual code
*
*  Usage:
*        BestMascot() is imported into a main.go file and a mascot_test.go file
*		 and it returns the string "Go Gopher"
*
*  Files:
*        mascot.go      : Go file that holds BestMascot() func
*        mascot_test.go : Go file to test proper functionality of BestMascot()
*        main.go        : main driver file that calls BestMascot()
*		 go.mod			: describes module's properties
*****************************************************************************/

package main

import (
	"fmt"

	"github.com/bsmith578/4143-PLC/tree/main/Assignments/P01/mascot"
)

func main() {
	fmt.Println(mascot.BestMascot())
}
