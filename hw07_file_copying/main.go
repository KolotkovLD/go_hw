package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//var (
//	from, to      string
//	limit, offset int64
//)
//
//func init() {
//	flag.StringVar(&from, "from", "", "file to read from")
//	flag.StringVar(&to, "to", "", "file to write to")
//	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
//	flag.Int64Var(&offset, "offset", 0, "offset in input file")
//}

func main() {
	flag.Parse()
	from := flag.String("from", "", "Path to source file")
	to := flag.String("to", "", "Path to destination file")
	offset := flag.Int64("offset", 0, "Offset in source file")
	limit := flag.Int64("limit", 0, "Number of bytes to copy (0 means entire file)")

	if *from == "" || *to == "" {
		fmt.Println("Both -from and to arguments are requred")
		flag.Usage()
		os.Exit(1)
	}
	err := CopyFile(*from, *to, *offset, *limit)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
	}
	fmt.Printf("File %s copied to %s\n", *from, *to)

}
