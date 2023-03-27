package utils

import (
	"flag"
	"log"
)

type Args struct {
	SymlinkFlag   bool
	DirectoryFlag bool
	FileFlag      bool
	Extension     string
	Paths         []string
}

func ParseFindFlags() Args {
	args := Args{}
	flag.BoolVar(&(args.FileFlag), `f`, false, `find files`)
	flag.BoolVar(&(args.DirectoryFlag), `d`, false, `find directories`)
	flag.BoolVar(&(args.SymlinkFlag), `sl`, false, `find symlinks`)
	flag.StringVar(&(args.Extension), `ext`, ``, `find symlinks`)
	flag.Parse()

	if args.Extension != `` && !args.FileFlag {
		log.Fatal(`flag -ext can be specified only when flag -f is present`)
	}

	if !args.DirectoryFlag && !args.FileFlag && !args.SymlinkFlag {
		args = Args{true, true, true, ``, []string{}}
	}

	args.Paths = append(args.Paths, flag.Args()...)

	if len(args.Paths) == 0 {
		args.Paths = []string{`./`}
	}

	return args
}
