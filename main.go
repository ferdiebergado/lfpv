package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
)

func main() {
	f := os.Args[1]
	w := os.Args[2]
	h := os.Args[3]
	ext := filepath.Ext(f)

	fmt.Println(f, w, h)

	m := mime.TypeByExtension(ext)
	fmt.Println(m)
}
