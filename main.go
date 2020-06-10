package main

import (
	"github.com/Herlitzd/granule/pkg"
	"log"
	"os"
)

/*
master --------------------------------------------------------- 1.1.3
						develop--------------------------------------------- 1.1.4
												feature/123----------------------------- 1.1.10
*/
func main() {
	path := os.Args[1]
	out := pkg.ParseConfig(&path)
	log.Printf("\n%s", out.String())

}
