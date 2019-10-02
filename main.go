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
		fmt.Printf("Unable to find the file \"%s\". Please try again.\n", path)
		os.Exit(1)
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func printUseage() {
    fmt.Println("Incorrect Usage")
	fmt.Println("Run \"filehash help\" for more info")
}

func compHash(h1, h2 string) {
    if h1 == h2{
        fmt.Println("Checksums match!")
    } else {
        fmt.Println("Checksums DO NOT match!")
    }
}
func printHelp() {
	fmt.Println("Generate a sha256 checksum for a specified file or compare a file against another checksum.")
	fmt.Println("Usage: filehash file_name [checksum]")
	fmt.Println("The first file_name parameter is required.")
	fmt.Println("You may optionally add a sha256 checksum to compare against as second parameter.")
}
func main() {
	if len(os.Args) == 3 {
		path := os.Args[1]
		argHash := strings.ToLower(os.Args[2])

		fileHash := hashFile(path)
		compHash(fileHash, argHash)
	} else if len(os.Args) == 2 {

		if strings.ToLower(os.Args[1]) == "help"{
			printHelp()
		} else {
		path := os.Args[1]
		fileHash := hashFile(path)
		fmt.Println(fileHash)
		}
	} else {
		printUseage()
	}
}
