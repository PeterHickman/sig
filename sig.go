package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	ep "github.com/PeterHickman/expand_path"
	"io"
	"os"
	"path/filepath"
)

var root string
var number_of_bytes int64

func init() {
	var n = flag.Int64("size", 16, "Number of bytes to read")

	flag.Parse()

	if len(flag.Args()) == 1 {
		root = flag.Arg(0)
	} else {
		fmt.Println("Must provide the directory to start the scan from")
		os.Exit(1)
	}

	root, _ = ep.ExpandPath(root)

	number_of_bytes = *n

	if number_of_bytes < 1 {
		fmt.Printf("Number of bytes must be a positive integer, not [%d]\n", number_of_bytes)
		os.Exit(2)
	}
}

func process(file string) {
	f, err := os.Open(file)
	if err != nil {
		return
	}

	fi, _ := f.Stat()
	if fi.Size() < number_of_bytes {
		fmt.Printf("too_small_%d %s\n", fi.Size(), file)
		return
	}

	defer f.Close()

	header := make([]byte, number_of_bytes)
	_, err = io.ReadFull(f, header[:])
	if err != nil {
		return
	}

	fmt.Printf("%s %s\n", hex.EncodeToString(header), file)
}

func main() {
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			// Weird filesystem errors
			if info == nil {
				os.Stderr.WriteString(fmt.Sprintf("%s\n %s\n\n", path, err))
				return nil
			}

			if !info.IsDir() {
				process(path)
			}

			return nil
		})

	if err != nil {
		fmt.Println(err)
		os.Exit(6)
	}
}
