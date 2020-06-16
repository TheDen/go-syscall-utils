package main

import (
	"fmt"
	"github.com/theden/system/syslib"
	"log"
)

func main() {
	filesystem := "/"
	inodes, err := syslib.Inodes(filesystem)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(inodes)
	data, err := syslib.Uptime()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
}
