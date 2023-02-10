package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"toys/morse/morse"
)

func main() {

	userCli := flag.Bool("d", false, "Demorse The Morse Code")
	flag.Parse()

	if *userCli {
		fmt.Println("What do you wanna Demorse: ")
		reader := bufio.NewReader(os.Stdin)
		giberish, _ := reader.ReadString('\n')
		fmt.Printf("%s", morse.Demorser(giberish))
	} else {
		fmt.Println("What do you wanna morse: ")
		reader := bufio.NewReader(os.Stdin)
		giberish, _ := reader.ReadString('\n')
		fmt.Printf("%s", morse.Morser(giberish))
	}

}
