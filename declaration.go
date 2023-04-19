package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		var text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println(string(text[0]))
	}
}
