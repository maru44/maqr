package main

import (
	"flag"
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("need one or two args")
	}

	fileName := "qr.png"
	if len(args) > 1 {
		fileName = args[1]
	}

	qrcode.WriteFile(args[0], qrcode.Medium, 256, fileName)
	fmt.Println(fmt.Sprintf("%s is exported", fileName))
}
