// https://slurm.schedmd.com/squeue.html
package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
)

//go:embed outputs/*
var f embed.FS

func main() {
	// this functions args will be different most every time because its specific to a file path.
	// the possible outcomes are
	// 		1. successful .sh script which results in:
	//			Submitted batch job X
	//		2. file exists but is empty:
	//			Exit status 1
	//			sbatch: error: Batch script is empty
	//		3. file doesn't exist:
	//			Exit status: 1
	//			sbatch: error: Unable to open file <path>
	//
	// Going to just map everything to successful run.

	//arguments := os.Args[1:]
	//hash := hasher.Hasher(arguments)
	//hash_path := fmt.Sprintf("outputs/%s.txt", hash)
	hash_path := "outputs/f70aa748f40c53d1eafa1fdd2e917cfa.txt"
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
