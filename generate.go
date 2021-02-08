package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var count int
	flag.IntVar(&count, "n", 261, "how many return values to generate")
	flag.Parse()
	if count < 2 {
		fmt.Fprintln(os.Stderr, "Error: -n cannot be below 2")
		os.Exit(1)
	}

	out(count, template[0], "_, ")
	out(count, template[1], "int, ")
	out(count, template[2], "0, ")
	fmt.Print(template[3])
}

func out(count int, prefix, rep string) {
	fmt.Print(prefix)
	for i := 2; i < count; i++ {
		fmt.Print(rep)
	}
}

var template = []string{
	`package main

import (
	"fmt"
)

func main() {
	val, `, `last := ohgodwhy()

	fmt.Println(val, last)
}

func ohgodwhy() (int, `, `int) {
	return 2, `, `12
}
`}
