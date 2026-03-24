package main

import (
	"fmt"

	"github.com/duke-git/lancet/v2/strutil"
)

func main() {
	reversed := strutil.Reverse("Hello")
	fmt.Println(reversed)
}
