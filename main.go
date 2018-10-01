package main

/*		 TODO
handel md5, sha1, sha256, sha512
*/

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

func compHash(fileHash, argHash string) bool {
	return fileHash == argHash
}

func main() {
	if len(os.Args) == 3 {
		path := os.Args[1]
		argHash := strings.ToLower(os.Args[2])

		fileHash := hashFile(path)
		res := compHash(fileHash, argHash)
		fmt.Println(res)
	} else if len(os.Args) == 2 {
		path := os.Args[1]
		fileHash := hashFile(path)
		fmt.Println(fileHash)
	} else {
		fmt.Println("Incorrect Usage")
        fmt.Println("Usage: ./filehash file [checksum]")
	}
}
