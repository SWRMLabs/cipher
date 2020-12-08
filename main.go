package main

import (
	"flag"
	"fmt"
	"github.com/StreamSpace/cipher/lib"
)

// Command arguments
var (
	key    = flag.String("key", "", "Cipher key")
	file   = flag.String("file", "", "location of file to be encrypted")
)

func main() {
	flag.Parse()
	if *key == "" || *file == "" {
		fmt.Println("cipher key or file flag is blank")
		return
	}
	filename, err := lib.Encrypt(*key, *file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypted file generated at "+filename)
}