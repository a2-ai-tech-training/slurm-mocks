// https://slurm.schedmd.com/sacct.html
package main

import (
	"bufio"
	"embed"
	"fmt"
	"github.com/a2-ai-tech-training/build-go-bin/internal/hasher"
	"log"
	"os"
)

//go:embed outputs/*
var f embed.FS

func main() {

	arguments := os.Args[1:]
	hash := hasher.Hasher(arguments)

	hash_path := fmt.Sprintf("outputs/%s.txt", hash)

	file, err := f.Open(hash_path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skips first line that contains command
	scanner.Scan()

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
