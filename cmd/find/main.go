package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mokan-r/go-cli-tools/internal/utils"
)

func printFiles(path string, args *utils.Args) {
	info, _ := os.Lstat(path)

	if info != nil && info.Mode()&os.ModeSymlink == os.ModeSymlink {
		if linkTo, err := filepath.EvalSymlinks(path); err == nil {
			fmt.Println(path + ` -> ` + linkTo)
		} else {
			fmt.Println(path + ` -> ` + `[broken]`)
		}
	} else if info != nil && info.IsDir() && args.DirectoryFlag {
		fmt.Println(path)
	} else if args.FileFlag {
		if (args.Extension != `` && filepath.Ext(path) == (`.`+args.Extension)) || args.Extension == `` {
			fmt.Println(path)
		}
	}
}

func Find(args *utils.Args) {
	for _, path := range args.Paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf(`path does not exists: %s`, path)
		}
		fmt.Println(path)
		recursiveFind(path, args)
	}
}

func recursiveFind(path string, args *utils.Args) {
	files, err := os.ReadDir(path)
	if err == nil {
		for _, f := range files {
			printFiles(filepath.Join(path, f.Name()), args)
			if f.IsDir() {
				recursiveFind(filepath.Join(path, f.Name()), args)
			}
		}
	}
}

func main() {
	args := utils.ParseFindFlags()
	Find(&args)
}
