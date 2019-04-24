package lesphina

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"regexp"
	"runtime/debug"
	"strings"
)

var (
	// regex to read in- and out-params of interfaces' methods
	rParams = regexp.MustCompile(`func\((.*?)\)(.*)`)

	// regex to parse map type
	rMap = regexp.MustCompile(`map\[(.*?)\](.*)`)
)

type Meta struct {
	NumImport    int `json:"num_import"`
	NumConst     int `json:"num_const"`
	NumVar       int `json:"num_var"` // package level Var
	NumStruct    int `json:"num_struct"`
	NumInterface int `json:"num_interface"`
	NumFunction  int `json:"num_function"`

	Imports    []*Import    `json:"imports"`
	Consts     []*Const     `json:"consts"`
	Vars       []*Var       `json:"vars"`
	Structs    []*Struct    `json:"structs"`
	Interfaces []*Interface `json:"interfaces"`
	Functions  []*Function  `json:"functions"`
}

func (m *Meta) Json() string {
	j, _ := json.MarshalIndent(m, "", "    ")
	return string(j)
}

func parseSource(source string) (meta *Meta, err error) {
	defer func() {
		if r := recover(); r != nil {
			meta, err = nil, fmt.Errorf("lesphina parsing panic: %v", r.(error))
			fmt.Println(string(debug.Stack()))
		}
	}()

	meta = &Meta{}
	fset := token.NewFileSet()
	ff, err := parser.ParseFile(fset, source, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// an internal-use map to distinguish different 'type' declarations
	// key is the position of keyword literal 'struct' or 'interface'
	// value is the name of such type
	posType := make(map[int]string)

	ast.Inspect(ff, func(n ast.Node) bool {

		switch d := n.(type) {
		case *ast.GenDecl:
			// GenDecl (general declarations): import, var, const, type

			switch d.Tok {
			case token.IMPORT:
				// in case of `import ( ... )`:
				// there're several paths in one single `import` declaration.
				// this works for `import ...` case, too.
				// NOTE: this Specs pattern is similar for CONST or VAR, etc.
				for _, imp := range d.Specs {
					o := imp.(*ast.ImportSpec)
					i := &Import{}
					if o.Name != nil {
						i.Alias = o.Name.Name
					}
					if o.Path != nil {
						i.Name = o.Path.Value
					}
					meta.Imports = append(meta.Imports, i)
					meta.NumImport++
				}

			case token.CONST:
				for _, imp := range d.Specs {
					o := imp.(*ast.ValueSpec)

					var cc []*Const

					// `const aa, bb (type) = .., ..`
					// NOTE: number of values should match the number of names here;
					// and a const must have a value
					lenN := len(o.Names)
					lenV := len(o.Values)

					for i := 0; i < minInt(lenN, lenV); i++ {
						c := &Const{
							Name:    o.Names[i].Name,
							RawType: getNodeRawString(fset, o.Type),
						}
						if i < lenV {
							c.RawValue = getNodeRawString(fset, o.Values[i])
						}
						cc = append(cc, c)
					}

					meta.Consts = append(meta.Consts, cc...)
					meta.NumConst += minInt(lenN, lenV)
				}

			case token.VAR:
				for _, imp := range d.Specs {
					o := imp.(*ast.ValueSpec)

					var vv []*Var
					lenN := len(o.Names)
					lenV := len(o.Values)

					for i := 0; i < lenN; i++ {
						v := &Var{
							Name:    o.Names[i].Name,
							RawType: getNodeRawString(fset, o.Type),
						}
						// if number of values is less than numbers of names,
						// the exceeding names have no value.
						if i < lenV {
							v.RawValue = getNodeRawString(fset, o.Values[i])
							v.IsFunc = strings.HasPrefix(v.RawValue, "func(")
						}
						vv = append(vv, v)
					}

					meta.Vars = append(meta.Vars, vv...)
					meta.NumVar += minInt(lenN, lenV)
				}

			case token.TYPE:
				// there're mainly 3 kinds of 'TYPE' declaration:
				// 'struct', 'interface' and alias

				// we can get names before we go further to 'struct' or 'interface' keywords
				// normally it has one and only one of such Specs
				tName := d.Specs[0].(*ast.TypeSpec).Name.Name

				// position of literal 'interface', 'struct' or types (for alias) keyword
				tPos := fset.Position(d.Pos()).Offset
				posType[tPos+len(tName)+6] = tName // 6 => 4 (len of keyword 'type') + 2 spaces
			}

		case *ast.StructType:
			// struct declaration

			meta.NumStruct++

			// this d is the literal 'struct' keywords
			// check its type name if its position exists in temp pos map
			tPos := fset.Position(d.Pos()).Offset
			if tName, ok := posType[tPos]; ok {
				// we get a struct

				str := &Struct{
					Name:    tName,
					RawBody: getNodeRawString(fset, d),
				}

				// listing *ast.Field, the fields of such struct
				for _, f := range d.Fields.List {
					ele := &Element{
						Name:    getNameFromIdents(f.Names),
						RawType: getNodeRawString(fset, f.Type),
					}
					if f.Tag != nil {
						ele.RawTag = f.Tag.Value
					}
					parseEle(ele)
					str.Fields = append(str.Fields, ele)
				}

				meta.Structs = append(meta.Structs, str)
			}

		case *ast.InterfaceType:
			// interface declaration

			meta.NumInterface++

			// this d is the literal 'interface' keywords
			// check its type name if its position exists in temp pos map
			tPos := fset.Position(d.Pos()).Offset
			if tName, ok := posType[tPos]; ok {
				// we get an interface

				intf := &Interface{
					Name:    tName,
					RawBody: getNodeRawString(fset, d),
				}

				if d.Methods.NumFields() > 0 {
					for _, m := range d.Methods.List {
						tmp := &InterfaceMethod{
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
			// function declaration

			meta.NumFunction++

			fun := &Function{
				Name:    d.Name.Name,
				RawBody: getNodeRawString(fset, d.Body),
			}

			// receivers
			if d.Recv.NumFields() > 0 {
				for _, r := range d.Recv.List {
					recv := &Element{
						Name:    getNameFromIdents(r.Names),
						RawType: getNodeRawString(fset, r.Type),
					}
					parseEle(recv)
					fun.Recv = append(fun.Recv, recv)
				}
			}

			// in params
			if d.Type.Params.NumFields() > 0 {
				for _, p := range d.Type.Params.List {
					ele := &Element{
						Name:    getNameFromIdents(p.Names),
						RawType: getNodeRawString(fset, p.Type),
					}
					parseEle(ele)
					fun.In = append(fun.In, ele)
				}
			}

			// out params
			if d.Type.Results.NumFields() > 0 {
				for _, r := range d.Type.Results.List {
					ele := &Element{
						Name:    getNameFromIdents(r.Names),
						RawType: getNodeRawString(fset, r.Type),
					}
					parseEle(ele)
					fun.Out = append(fun.Out, ele)
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

func getInterfaceMethodDetail(m *InterfaceMethod) {
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
		ele := &Element{
			Name:    inParams[i][0],
			RawType: inParams[i][1],
		}
		parseEle(ele)
		m.In = append(m.In, ele)
	}

	for i := len(outParams) - 1; i >= 0; i-- {
		ele := &Element{
			Name:    outParams[i][0],
			RawType: outParams[i][1],
		}
		parseEle(ele)
		m.Out = append(m.Out, ele)
	}
}

func parseEle(ele *Element) {
	// parse deeper of element

	if ele.RawType == "" {
		return
	}

	ele.BaseType = ele.RawType

	if strings.HasPrefix(ele.BaseType, "...") {
		ele.IsVariadic = true
		ele.BaseType = strings.TrimLeft(ele.BaseType, "...")
	}

	if strings.HasPrefix(ele.BaseType, "*") {
		ele.IsPointer = true
		ele.BaseType = strings.TrimLeft(ele.RawType, "*")
	}

	if strings.HasPrefix(ele.BaseType, "[]") {
		ele.IsSlice = true
		ele.BaseType = strings.TrimLeft(ele.BaseType, "[]")
	}

	if strings.HasPrefix(ele.BaseType, "map") {
		ele.IsMap = true

		matches := rMap.FindStringSubmatch(ele.BaseType)
		if len(matches) != 3 {
			return
		}

		ele.KeyType = matches[1]
		ele.ValueType = matches[2]
	}
}

func getArgs(raw string) (res [][]string) {
	// structuring go pattern argument list
	// raw is like: "a1, a2 t1, b t2", "t1, t2 (only types)", "a1 t1" or just "t1"
	// no such: "a1 t1, t2" (cannot omit any name except all names are omitted)
	// and cannot omit the last type

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
			if lastType == "" {
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

func minInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
