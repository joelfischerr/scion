// +build !JFtesting

package main

import "fmt"

const jftesting = false

func testOptions() {
	fmt.Println("We are in not testing!")
}
