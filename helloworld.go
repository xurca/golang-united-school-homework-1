package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	helloWorld := helloworld()

	fmt.Println(helloWorld)
}

func helloworld() string {
	return emoji.Sprint("Hello :world_map:!")
}
