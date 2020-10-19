package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/olongfen/gormgen"
)

type config struct {
	input   string
	imports []gormgen.ImportPkg
	structs []string
}

var (
	cnf          config
	logName      string
	transformErr bool
)

func parseFlags() {

	var input, structs, imports string
	flag.StringVar(&structs, "structs", "", "[Required] The name of schema structs to generate structs for, comma seperated")
	flag.StringVar(&input, "input", "", "[Required] The name of the input file dir")
	flag.StringVar(&imports, "imports", "", "[Required] The name of the import  to import package")
	flag.StringVar(&logName, "logName", "", "[Option] The name of log db error")
	flag.BoolVar(&transformErr, "transformErr", false, "[Option] The name of transform db err")
	flag.Parse()

	if input == "" || structs == "" || len(imports) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	cnf = config{
		input:   input,
		structs: strings.Split(structs, ","),
	}
	s := strings.Split(imports, ",")
	for _, v := range s {
		cnf.imports = append(cnf.imports, gormgen.ImportPkg{
			Pkg: v,
		})
	}
}

func main() {
	parseFlags()

	p := gormgen.NewParser(cnf.input)

	gen := gormgen.NewGenerator(cnf.input).SetImportPkg(cnf.imports).SetLogName(logName)
	if transformErr {
		gen = gen.TransformError()
	}
	if err := gen.ParserAST(p, cnf.structs).Generate().Format().Flush(); err != nil {
		log.Fatalln(err)
	}

}
