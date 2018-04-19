package lesphina

import (
	"bytes"
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"regexp"
	"strings"

	item "github.com/straightdave/lesphina/item"
)

var (
	// counters
	nTotal, nImport, nConst, nType, nVar, nOther uint

	// regex to read in- and out-params of interfaces' methods
	rParams = regexp.MustCompile(`func\((.*?)\)(.*)`)
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

func parseSource(source string) (meta *Meta, err error) {
	defer func() {
		if r := recover(); r != nil {
			meta, err = nil, r.(error)
		}
	}()

	meta = &Meta{}
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
				// normally it has one and only one of such Specs
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

				if d.Methods.NumFields() > 0 {
					for _, m := range d.Methods.List {
						tmp := &item.InterfaceMethod{
							Name:    getNameFromIdents(m.Names),
							RawType: getNodeRawString(fset, m.Type),
						}

						getInterfaceMethodDetail(tmp)
						intf.Methods = append(intf.Methods, tmp)
					}
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

func getInterfaceMethodDetail(m *item.InterfaceMethod) {
	rawType := m.RawType // "func(.. ..) ...."

	tmp := rParams.FindStringSubmatch(rawType)

	// if matches, tmp must be 3 parts (whole matched string, group1, group2)
	if len(tmp) != 3 {
		return
	}

	group1 := tmp[1]
	group2 := strings.TrimSpace(tmp[2])
	group2 = strings.TrimPrefix(group2, "(")
	group2 = strings.TrimSuffix(group2, ")")

	inParams := getArgs(group1)
	outParams := getArgs(group2)

	for i := len(inParams) - 1; i >= 0; i-- {
		ele := &item.Element{
			Name:    inParams[i][0],
			RawType: inParams[i][1],
		}
		m.In = append(m.In, ele)
	}

	for i := len(outParams) - 1; i >= 0; i-- {
		ele := &item.Element{
			Name:    outParams[i][0],
			RawType: outParams[i][1],
		}
		m.Out = append(m.Out, ele)
	}
}

func getArgs(raw string) (res [][]string) {
	// raw is like: "a1, a2 t1, b t2", "t1, t2", "a1 t1" or just "t1"

	if raw == "" {
		return
	}

	parts := strings.Split(raw, ",")

	var lastType string

	for i := len(parts) - 1; i >= 0; i-- {
		var arg []string

		tmp := strings.TrimSpace(parts[i])
		innerParts := strings.Split(tmp, " ")
		if len(innerParts) == 2 {
			arg = append(arg, innerParts[0])
			arg = append(arg, innerParts[1])
			res = append(res, arg)

			lastType = innerParts[1]
		} else if len(innerParts) == 1 {
			if len(parts) == 1 {
				arg = append(arg, "")
				arg = append(arg, innerParts[0])
			} else {
				arg = append(arg, innerParts[0])
				arg = append(arg, lastType)
			}
			res = append(res, arg)
		}
	}

	return
}

// ---- special helpers

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

func Jsonify(obj interface{}) string {
	res, _ := json.MarshalIndent(obj, "", "    ")
	return string(res)
}
