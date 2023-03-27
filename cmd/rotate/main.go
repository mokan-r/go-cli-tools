package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Tar struct {
	filePaths []string
	outputDir string
}

func ParseFiles() (Tar, error) {
	args := Tar{}
	flag.StringVar(&(args.outputDir), `a`, ``, `count characters in file`)
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("you must specify path to at least one file")
	}

	args.filePaths = append(args.filePaths, flag.Args()...)

	return args, nil
}

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	header.Name = filename
	header.ModTime = time.Now()
	header.AccessTime = time.Now()
	header.ChangeTime = time.Now()

	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}

func createArchive(fileName string, dir string, c chan error) {
	info, _ := os.Stat(fileName)
	if !strings.HasSuffix(fileName, `.log`) {
		c <- fmt.Errorf(`file must be a .log extension`)
		return
	}

	archiveName := strings.TrimSuffix(fileName, `.log`)

	if dir != `` {
		archiveName = filepath.Join(dir, filepath.Base(archiveName))
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			c <- fmt.Errorf(`couldn't create directory to archive place`)
		}
	}

	outFile, err := os.Create(archiveName + strconv.FormatInt(info.ModTime().Unix(), 10) + `.tar.gz`)
	if err != nil {
		c <- fmt.Errorf("error creating archive for file: %s", fileName)
	}
	tw := tar.NewWriter(outFile)
	defer tw.Close()

	c <- addToArchive(tw, fileName)
}

func main() {
	args, err := ParseFiles()
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan error, len(args.filePaths))
	for _, path := range args.filePaths {
		go createArchive(path, args.outputDir, c)
	}

	for i := 0; i < len(args.filePaths); i++ {
		if e := <-c; e != nil {
			fmt.Println(e)
		}
	}
	close(c)
}
