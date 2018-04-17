package lesphina

import (
	"bytes"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"

	item "github.com/straightdave/lesphina/item"
)

var (
	// counters
	nTotal, nImport, nConst, nType, nVar, nOther uint
)

type Meta struct {
	NumImport    uint `json:"num_import"`
	NumVar       uint `json:"num_var"` // package level Var
	NumStruct    uint `json:"num_struct"`
	NumInterface uint `json:"num_interface"`
	NumFunction  uint `json:"num_function"`

	Imports    []*item.Import    `json:"imports"`
	Vars       []*item.Var       `json:"vars"`
	Structs    []*item.Struct    `json:"structs"`
	Interfaces []*item.Interface `json:"interfaces"`
	Functions  []*item.Function  `json:"functions"`
}

func (m *Meta) Json() string {
	res, _ := json.MarshalIndent(m, "", "    ")
	return string(res)
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
				// there're mainly 3 kinds of 'type' declaration:
				// 'struct', 'interface' and alias
				nType++

				// we can get names before we go further to 'struct' or 'interface' keywords
				// normally the Specs of such declaration is one and only one
				tName := d.Specs[0].(*ast.TypeSpec).Name.Name

				// position of literal 'interface', 'struct' or types (for alias) keyword
				tPos := fset.Position(d.Pos()).Offset
				posType[tPos+len(tName)+6] = tName // 6 = 4 (len of keyword 'type') + two spaces

			default:
				// other GenDecl tokens (none in theory)
				nOther++
			}

		case *ast.StructType:
			// happens when parsing the 'struct' keywords
			// and this happens normally after parsing its 'type' keyword

			meta.NumStruct++

			// this d is the literal 'struct' keywords
			// check its type name if its position exists in temp pos map
			tPos := fset.Position(d.Pos()).Offset
			if tName, ok := posType[tPos]; ok {
				// we get a struct

				str := &item.Struct{
					Name:    tName,
					RawBody: getNodeRawString(fset, d),
				}

				// listing *ast.Field, the fields of such struct
				for _, f := range d.Fields.List {
					str.Fields = append(str.Fields, &item.Element{
						Name:    getNameFromIdents(f.Names),
						RawType: getNodeRawString(fset, f.Type),
					})
				}

				meta.Structs = append(meta.Structs, str)
			}

		case *ast.InterfaceType:
			// happens when parsing the 'interface' keywords
			// and this happens normally after parsing its 'type' keyword

			meta.NumInterface++

			// this d is the literal 'interface' keywords
			// check its type name if its position exists in temp pos map
			tPos := fset.Position(d.Pos()).Offset
			if tName, ok := posType[tPos]; ok {
				// we get an interface

				intf := &item.Interface{
					Name:    tName,
					RawBody: getNodeRawString(fset, d),
				}

				meta.Interfaces = append(meta.Interfaces, intf)
			}

		case *ast.FuncDecl:
			// happens when parsing the 'func' keywords

			meta.NumFunction++

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

// special helpers

func (les *Lesphina) MethodsOfStruct(s *item.Struct) []*item.Function {
	var res []*item.Function
	for _, f := range les.Meta.Functions {
		// for now only consider single receiver
		if len(f.Recv) == 1 {
			// for now, using 'contains' to tell (can cover pointer condition);
			// but may not work for maps or slices
			if strings.Contains(f.Recv[0].RawType, s.Name) {
				res = append(res, f)
			}
		}
	}
	return res
}
