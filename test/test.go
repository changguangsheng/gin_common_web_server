package main

import (
	"flag"
	"fmt"
)

func main() {
	worPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bot")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Println("word:", *worPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

}
