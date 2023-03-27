package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

type wordCount struct {
	filePaths     []string
	charCountFlag bool
	wordCountFlag bool
	lineCountFlag bool
}

func ParseCountFlags() wordCount {
	args := wordCount{}
	flag.BoolVar(&(args.charCountFlag), `m`, false, `count characters in file`)
	flag.BoolVar(&(args.wordCountFlag), `w`, false, `count words in file`)
	flag.BoolVar(&(args.lineCountFlag), `l`, false, `count lines in file`)
	flag.Parse()

	if (args.charCountFlag && args.wordCountFlag) ||
		(args.charCountFlag && args.lineCountFlag) ||
		(args.wordCountFlag && args.lineCountFlag) {
		log.Fatal("only one flag can be used at a time")
	}

	if !args.charCountFlag && !args.wordCountFlag && !args.lineCountFlag {
		args.wordCountFlag = true
	}

	if flag.NArg() == 0 {
		log.Fatal("you should specify path to file")
	}

	args.filePaths = append(args.filePaths, flag.Args()...)

	return args
}

func ReadFile(args *wordCount, path string, c chan string) {
	count := 0
	file, err := os.Open(path)
	if err != nil {
		c <- `myWc: ` + err.Error()
		return
	} else if info, _ := file.Stat(); info.IsDir() {
		c <- `myWc: ` + path + `: is a directory`
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if args.charCountFlag {
			count += utf8.RuneCountInString(scanner.Text())
		} else if args.wordCountFlag {
			count += len(strings.Fields(scanner.Text()))
		} else if args.lineCountFlag {
			count++
		}
	}
	file.Close()

	c <- strconv.Itoa(count) + ` ` + path
}

func main() {
	args := ParseCountFlags()
	c := make(chan string, len(args.filePaths))

	for _, path := range args.filePaths {
		go ReadFile(&args, path, c)
	}

	for i := 0; i < len(args.filePaths); i++ {
		fmt.Println(<-c)
	}
	close(c)
}
