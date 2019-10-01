package main


import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func hashFile(path string) string {
	file, err := os.Open(path)
    if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal()
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func printUseage() {
    fmt.Println("Incorrect Usage")
    fmt.Println("Usage: ./filehash file [checksum]")
}

func compHash(h1, h2 string) {
    if h1 == h2{
        fmt.Println("Checksums match!")
    } else {
        fmt.Println("Checksums DO NOT match!")
    }
}

func main() {
	if len(os.Args) == 3 {
		path := os.Args[1]
		argHash := strings.ToLower(os.Args[2])

		fileHash := hashFile(path)
		compHash(fileHash, argHash)
	} else if len(os.Args) == 2 {
		path := os.Args[1]
		fileHash := hashFile(path)
		fmt.Println(fileHash)
	} else {
		printUseage()
	}
}
