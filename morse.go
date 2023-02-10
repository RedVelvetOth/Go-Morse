package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"toys/morse/morse"

	"github.com/f1bonacc1/glippy"
)

func main() {

	userCli := flag.Bool("d", false, "Demorse The Morse Code")
	flag.Parse()

	if *userCli {
		fmt.Println("What do you wanna Demorse: ")
		reader := bufio.NewReader(os.Stdin)
		giberish, _ := reader.ReadString('\n')
		memes := morse.Demorser(giberish)
		glippy.Set(memes)
		fmt.Printf("%s", memes)
	} else {
		fmt.Println("What do you wanna morse: ")
		reader := bufio.NewReader(os.Stdin)
		giberish, _ := reader.ReadString('\n')
		memes := morse.Morser(giberish)
		glippy.Set(memes)
		fmt.Printf("%s", memes)
	}

}
