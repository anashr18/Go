package main

import (
	"fmt"

	"github.com/anashr18/toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.GenerateRandomString(10)
	fmt.Println(s)
}
