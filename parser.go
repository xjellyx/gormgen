package gormgen

import (
	"github.com/jinzhu/gorm"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
)

// The Parser is used to parse a directory and expose information about the structs defined in the files of this directory.
type Parser struct {
	dir         string
	pkg         *build.Package
	parsedFiles []*ast.File
}

// NewParser create a new parser instance.
func NewParser(dir string) *Parser {
	return &Parser{
		dir: dir,
	}
}

// getPackage parse dir get go file and package
func (p *Parser) getPackage() {
	pkg, err := build.Default.ImportDir(p.dir, build.ImportComment)
	if err != nil {
		log.Fatalf("cannot process directory %s: %s", p.dir, err)
	}
	p.pkg = pkg

}

// parseGoFiles parse go file
func (p *Parser) parseGoFiles() {
	var parsedFiles []*ast.File
	fs := token.NewFileSet()
	for _, file := range p.pkg.GoFiles {
		file = p.dir + "/" + file
		parsedFile, err := parser.ParseFile(fs, file, nil, 0)
		if err != nil {
			log.Fatalf("parsing package: %s: %s\n", file, err)
		}
		parsedFiles = append(parsedFiles, parsedFile)
	}
	p.parsedFiles = parsedFiles
}

// parseTypes parse type of struct
func (p *Parser) parseTypes(file *ast.File) (ret []structConfig) {
	ast.Inspect(file, func(n ast.Node) bool {
		decl, ok := n.(*ast.GenDecl)
		if !ok || decl.Tok != token.TYPE {
			return true
		}
		for _, spec := range decl.Specs {
			var (
				data structConfig
			)
			typeSpec, _ok := spec.(*ast.TypeSpec)
			if !_ok {
				continue
			}
			// We only care about struct declaration (for now)
			var structType *ast.StructType
			if structType, ok = typeSpec.Type.(*ast.StructType); !ok {
				continue
			}
			data.StructName = typeSpec.Name.Name
			for _, v := range structType.Fields.List {
				var (
					field fieldConfig
				)
				// type is select expr, and sel is Model , else continue
				if _v, _ok := v.Type.(*ast.SelectorExpr); _ok {
					if _v.Sel.Name == "Model" {
						field.FieldName = "ID"
						field.FieldType = "uint"
						field.ColumnName = gorm.ToDBName("ID")
						data.Fields = append(data.Fields, field)
					}
					continue
				}

				// type is ident, get field type
				if t, _ok := v.Type.(*ast.Ident); _ok {
					field.FieldType = t.String()
				}
				// get file name
				if len(v.Names) > 0 {
					field.FieldName = v.Names[0].String()
					field.ColumnName = gorm.ToDBName(field.FieldName)
				}
				// get field tag
				if v.Tag != nil {
					field.Tag = v.Tag.Value
				}
				data.Fields = append(data.Fields, field)
			}
			ret = append(ret, data)
		}
		return true
	})
	return
}

// Parse should be called before any type querying for the parser. It takes the directory to be parsed and extracts all the structs defined in this directory.
func (p *Parser) Parse() (ret []structConfig) {
	var (
		data []structConfig
	)
	p.getPackage()
	p.parseGoFiles()
	for _, f := range p.parsedFiles {
		data = append(data, p.parseTypes(f)...)
	}
	return data
}
