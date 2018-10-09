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

func checkFile(path string) bool {
    file, err := os.Open(path)
    sErr := fmt.Sprintf("%x", err)
    fmt.Println(sErr)
    
    if strings.Contains(sErr, "no such file or directory") {
        return false
    } else if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    return true
}

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
	if len(os.Args) == 3{
        arg1 := os.Args[1]
        arg2 := os.Args[2]
        exArg1 := checkFile(arg1)
        exArg2 := checkFile(arg2)
        fmt.Println()
        
		if exArg1 && exArg2 {
			compHash(hashFile(arg1), hashFile(arg2))
		} else if exArg1 {
			compHash(hashFile(arg1), arg2)
		} else if exArg2 {
			compHash(arg1, hashFile(arg2))
		} else {
			compHash(arg1, arg2)
		}
	} else if len(os.Args) == 2{
		path := os.Args[1]
		fileHash := hashFile(path)
		fmt.Println(fileHash)
	} else {
		printUseage()
	}
}
