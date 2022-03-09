package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"log"
	"os"
	"reflect"
	"strings"
)

type tpl struct {
	FieldName string
}

var (
	intTpl = template.Must(template.New("intTpl").Parse(`
		// {{.FieldName}}
		var {{.FieldName}}Raw uint64
		binary.Read(r, binary.BigEndian, &{{.FieldName}}Raw)
		in.{{.FieldName}} = uint64({{.FieldName}}Raw)
`))
	strTpl = template.Must(template.New("strTpl").Parse(`
	// {{.FieldName}}
	var {{.FieldName}}LenRaw uint8
	binary.Read(r, binary.BigEndian, &{{.FieldName}}LenRaw)
	{{.FieldName}}Raw := make([]byte, {{.FieldName}}LenRaw)
	binary.Read(r, binary.BigEndian, &{{.FieldName}}Raw)
	in.{{.FieldName}} = string({{.FieldName}}Raw)
`))
	slByteTpl = template.Must(template.New("slByteTpl").Parse(`
	// {{.FieldName}}
	var {{.FieldName}}LenRaw uint8
	binary.Read(r, binary.BigEndian, &{{.FieldName}}LenRaw)
	{{.FieldName}}Raw := make([]byte, {{.FieldName}}LenRaw)
	binary.Read(r, binary.BigEndian, &{{.FieldName}}Raw)
	in.{{.FieldName}} = {{.FieldName}}Raw
`))
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	out, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintln(out, "package", node.Name.Name)
	fmt.Fprintln(out)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintln(out, `import "encoding/binary"`)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintln(out, `import "bytes"`)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(out)

	for _, f := range node.Decls {
		g, ok := f.(*ast.GenDecl)
		if !ok {
			fmt.Printf("SKIP %T is not ast.GeDecl\n", f)
			continue
		}
		for _, spec := range g.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				fmt.Printf("SKIP %T is not ast.TypeSpec\n", spec)
				continue
			}
			currStruct, ok := currType.Type.(*ast.StructType)
			_ = currStruct
			if !ok {
				fmt.Printf("SKIP %T is not ast.StructType\n", currStruct)
				continue
			}
			if g.Doc == nil {
				fmt.Printf("SKIP struct %v doesn't have comments\n", currType.Name.Name)
				continue
			}
			needCodegen := false
			for _, comment := range g.Doc.List {
				needCodegen = needCodegen || strings.HasPrefix(comment.Text, "// cgen: binpack")
			}
			if !needCodegen {
				fmt.Printf("SKIP struct %v doesn't have cgen mark\n", currType.Name.Name)
				continue
			}

			fmt.Printf(" process struct %s\n", currType.Name.Name)
			fmt.Printf("\tstart generating unpeck method\n")

			fmt.Fprintln(out, "func (in *"+currType.Name.Name+") Unpack(data []byte) error {")
			fmt.Fprintln(out, "\tr := bytes.NewReader(data)")

			for _, field := range currStruct.Fields.List {
				if field.Tag != nil {
					tag := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
					if tag.Get("cgen") == "-" {
						fmt.Printf("property %v skiped mark\n", field.Names[0].Name)
						continue
					}
				}

				fieldName := field.Names[0].Name
				var fieldType string
				if _, err := field.Type.(*ast.Ident); err {
					fieldType = field.Type.(*ast.Ident).Name
				} else {
					fieldType = "[]"
				}

				fmt.Printf("\tgenerating code for field %s.%s\n", fieldType, fieldName)
				switch fieldType {
				case "uint64":
					intTpl.Execute(out, tpl{fieldName})
				case "string":
					strTpl.Execute(out, tpl{fieldName})
				case "[]":
					slByteTpl.Execute(out, tpl{fieldName})
				}

			}

			fmt.Fprintln(out, "\treturn nil")
			fmt.Fprintln(out, "}")
		}
	}
}
