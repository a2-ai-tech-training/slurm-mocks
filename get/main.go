package main

import (
	"flag"
	"fmt"
	"os"
	//"github.com/spf13/cobra"
)

/*
	"crypto/md5"
	"encoding/hex"
	"fmt"
	//"io"
	"log"
	"os"
	"os/exec"
	"strings"*/

func main() {
	args := os.Args
	var a string
	var s string
	var o string
	fmt.Printf("args: %s\n\n", args)
	envO := os.Getenv("OUTPUT_DIR")
	envS := os.Getenv("SCENARIO")
	fmt.Println("$OUTPUT_DIR: ", envO)
	fmt.Printf("$SCENARIO: %s\n\n", envS)

	flag.StringVar(&a, "a", a, "The command to get")
	flag.StringVar(&o, "o", o, "Output Dir")
	flag.StringVar(&s, "s", s, "The user scenario")
	flag.Parse()
	otherArgs := flag.Args()

	fmt.Printf("flag a contents: %s\n", a)
	fmt.Printf("flag o contents: %s\n", o)
	fmt.Printf("flag s contents: %s\n", s)
	fmt.Printf("Other Args %s\n", otherArgs)

	//fmt.Printf("Before a: %s\nBefore o: %s\nBefore s: %s\n\n", a, o, s)

	if o == "" {
		if envO == "" {
			o = "outputs"
		} else {
			o = envO
			fmt.Println("Using env $OUTPUT_DIR:", o)
		}
	}

	fmt.Printf("Using output directory: %s\n", o)

	if s == "" {
		if envS == "" {
			s = "default"

		} else {
			s = envS
			fmt.Println("Using env $SCENARIO:", s)
		}
	}
	fmt.Printf("Using scenario: %s\n", s)

}

/*
func get(args []string) {

	//args := os.Args[1:]
	folder := os.Args[1]
	arguments := os.Args[2:]
	argString := strings.Join(args, " ")

	hasher := md5.New()

	for _, v := range arguments {
		hasher.Write([]byte(v))
	}

	hash := hex.EncodeToString(hasher.Sum(nil))

	name := fmt.Sprintf("%s/outputs/%s.txt", folder, hash)

	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}

	fmt.Printf("%s", b)

	wut2write := fmt.Sprintf("%s\n%s", argString, b)

	length, err := file.WriteString(wut2write)
	if err != nil {
		panic(err)
	}

	if length == 0 {
		fmt.Printf("No bytes written")
	}

	defer file.Close()

}*/
