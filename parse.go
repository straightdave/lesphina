package lesphina

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"

	item "github.com/straightdave/lesphina/item"
)

var (
	// counters
	nTotal, nImport, nConst, nType, nVar, nOther uint
	nFunction                                    uint
)

type Meta struct {
	Functions []*item.Function
}

func parseSource(source string) (*Meta, error) {
	meta := &Meta{}

	fset := token.NewFileSet()
	ff, err := parser.ParseFile(fset, source, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// internal used map to distinguish different 'type' declarations
	// key is the position of keyword literal 'struct' or 'interface'
	// value is the name of such type
	posType := make(map[int]string)

	ast.Inspect(ff, func(n ast.Node) bool {

		switch d := n.(type) {

		case *ast.GenDecl:
			// GenDecl (general declarations): import, var, const, type

			nTotal++

			switch d.Tok {
			case token.IMPORT:
				nImport++
			case token.CONST:
				nConst++
			case token.VAR:
				nVar++
			case token.TYPE:
				// type declaration mainly has 3 kinds:
				// 'struct', 'interface' and 'alias'
				nType++

				// we can get names before we go further to 'struct' or 'interface' keywords
				// normally the Specs of such declaration is only one
				tName := d.Specs[0].(*ast.TypeSpec).Name.Name

				// position of literal 'type' word
				tPos := fset.Position(d.Pos()).Offset
				posType[tPos+len(tName)+6] = tName // 6 = 4 (len of keyword 'type') + two spaces

			default:
				// other GenDecl tokens (none in theory)
				nOther++
			}

		// case *ast.StructType:
		// 	// happens when parsing the 'struct' keywords
		// 	// and this happens normally after parsing its 'type' keyword

		// 	nStruct++

		// 	// this d is the literal 'struct' keywords
		// 	// check its type name if its position exists in temp pos map
		// 	tPos := fset.Position(d.Pos()).Offset
		// 	if tName, ok := posType[tPos]; ok {
		// 		// we get a struct

		// 		str := &Struct{
		// 			Name: tName,
		// 		}

		// 		// listing *ast.Field, the fields of such struct
		// 		for _, f := range d.Fields.List {
		// 			str.Fields = append(str.Fields, &Element{
		// 				Name: getNameFromIdents(f.Names),
		// 				Type: getNodeRawString(fset, f.Type),
		// 			})
		// 		}

		// 		_structs = append(_structs, str)
		// 	}

		// case *ast.InterfaceType:
		// 	// happens when parsing the 'interface' keywords
		// 	// and this happens normally after parsing its 'type' keyword

		// 	nInterface++

		// 	// this d is the literal 'interface' keywords
		// 	// check its type name if its position exists in temp pos map
		// 	tPos := fset.Position(d.Pos()).Offset
		// 	if tName, ok := posType[tPos]; ok {
		// 		// we get an interface

		// 		intf := &Interface{
		// 			Name: tName,
		// 		}

		// 		_interfaces = append(_interfaces, intf)
		// 	}

		case *ast.FuncDecl:
			// happens when parsing the 'func' keywords

			nFunction++

			fun := &item.Function{
				Name:    d.Name.Name,
				RawBody: getNodeRawString(fset, d.Body),
			}

			// receivers
			if d.Recv.NumFields() > 0 {
				for _, r := range d.Recv.List {
					recv := &item.Element{
						Name:    getNameFromIdents(r.Names),
						RawType: getNodeRawString(fset, r.Type),
					}

					// if ex, ok := r.Type.(*ast.StarExpr); ok {
					// 	// is pointer
					// 	recv.IsPointer = true
					// 	recv.Type = getNodeRawString(fset, ex.X)
					// } else {
					// 	recv.Type = getNodeRawString(fset, r.Type)
					// }

					fun.Recv = append(fun.Recv, recv)
				}
			}

			// in params
			if d.Type.Params.NumFields() > 0 {
				for _, p := range d.Type.Params.List {
					fun.In = append(fun.In, &item.Element{
						Name:    getNameFromIdents(p.Names),
						RawType: getNodeRawString(fset, p.Type),
					})
				}
			}

			// out params
			if d.Type.Results.NumFields() > 0 {
				for _, r := range d.Type.Results.List {
					fun.Out = append(fun.Out, &item.Element{
						Name:    getNameFromIdents(r.Names),
						RawType: getNodeRawString(fset, r.Type),
					})
				}
			}

			meta.Functions = append(meta.Functions, fun)

		}

		return true
	})

	return meta, nil
}

func getNodeRawString(fset *token.FileSet, node ast.Node) string {
	var tmp bytes.Buffer
	printer.Fprint(&tmp, fset, node)
	return string(tmp.Bytes())
}

func getNameFromIdents(idents []*ast.Ident) (res string) {
	if len(idents) > 0 {
		res = idents[0].Name
	}
	return
}
