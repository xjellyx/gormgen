package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/olongfen/gormgen"
)

type config struct {
	input      string
	importPkgs []gormgen.ImportPkg
	structs    []string
}

var (
	cnf          config
	logName      string
	transformErr bool
)

func parseFlags() {
	var inputDir, structs, importPkgs string
	flag.StringVar(&structs, "structs", "", "[Required] The name of schema structs to generate structs for, comma seperated")
	flag.StringVar(&inputDir, "inputDir", "", "[Required] The name of the inputDir file")
	flag.StringVar(&importPkgs, "importPkgs", "", "[Required] The name of the importPkgs to import")
	flag.StringVar(&logName, "logName", "", "[Option] The name of log db error")
	flag.BoolVar(&transformErr, "transformErr", false, "[Option] The name of transform db err")
	flag.Parse()

	if inputDir == "" || structs == "" || len(importPkgs) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	cnf = config{
		input:   inputDir,
		structs: strings.Split(structs, ","),
	}
	s := strings.Split(importPkgs, ",")
	for _, v := range s {
		cnf.importPkgs = append(cnf.importPkgs, gormgen.ImportPkg{
			Pkg: v,
		})
	}
}

func main() {
	parseFlags()

	p := gormgen.NewParser(cnf.input)

	gen := gormgen.NewGenerator(cnf.input).SetImportPkg(cnf.importPkgs).SetLogName(logName)
	if transformErr {
		gen = gen.TransformError()
	}
	if err := gen.ParserAST(p, cnf.structs).Generate().Format().Flush(); err != nil {
		log.Fatalln(err)
	}

}
