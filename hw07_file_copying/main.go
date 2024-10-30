package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	from := flag.String("from", "", "Path to source file")
	to := flag.String("to", "", "Path to destination file")
	offset := flag.Int64("offset", 0, "Offset in source file")
	limit := flag.Int64("limit", 0, "Number of bytes to copy (0 means entire file)")

	if *from == "" || *to == "" {
		fmt.Println("Both -from and to arguments are required")
		flag.Usage()
		os.Exit(1)
	}
	err := CopyFile(*from, *to, *offset, *limit)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
	fmt.Printf("File %s copied to %s\n", *from, *to)
}
