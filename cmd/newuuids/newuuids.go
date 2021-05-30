package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/motevets/newuuids/pkg/uuidbump"
)

const usage = `
newuuids

Reads line-by-line from STDIN, replaces UUIDs with new random UUIDs, and prints result to STDOUT.
If it encounters a UUID that its replaced before, it will replace that UUID with the same UUID is used before.const

Example:

    cat /path/to/file/with/uuids | newuuids
`

func main() {
	if(len(os.Args[1:]) > 0) {
		fmt.Println(usage)
		os.Exit(0);
	}

	uuidRotator := uuidbump.New()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		out := uuidRotator.BumpUuids(scanner.Text())
		fmt.Println(out)
	}
}
