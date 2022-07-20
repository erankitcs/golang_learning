package module1

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"runtime"
	"strings"
)

//Function for reading the file
func readFile() *ast.File {

	_, currentFile, _, _ := runtime.Caller(1)

	src, err := ioutil.ReadFile(path.Join(path.Dir(currentFile), "../vehicle.go"))

	if err != nil {
		log.Fatal("This is the error", src)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)

	if err != nil {
		panic(err)
	}
	return f
}

// function for checking assiged variables within functions.
func checkVarWithinFunc(funcName, varName string) bool {
	foundVarName, foundFunc := false, false

	f := readFile()

	var funcBody []ast.Stmt

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			if funcDecl.Name.String() == funcName {
				funcBody = funcDecl.Body.List
				foundFunc = true
				break
			}
		}
	}

	if foundFunc {
		for _, b := range funcBody {
			if reflect.TypeOf(b).String() == "*ast.AssignStmt" {
				s := b.(*ast.AssignStmt)
				if s.Lhs[0].(*ast.Ident).String() == varName {
					foundVarName = true
				}
			}
		}

	}

	return foundFunc && foundVarName

}

// Function for checking map
func checkMap(mapName, mapKey, mapValue string) bool {

	foundMapName, foundMapKeyValue := false, false

	f := readFile()

	var findMap *ast.MapType

	//var vehicleResult map[string]feedbackResult

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.ValueSpec" {
					ValueSpec := spec.(*ast.ValueSpec)
					if reflect.TypeOf(ValueSpec.Type).String() == "*ast.MapType" {
						for _, name := range ValueSpec.Names {
							if name.String() == mapName {
								foundMapName = true
								findMap = ValueSpec.Type.(*ast.MapType)
								break
							}
						}
					}
				}
			}
		}
	}

	if foundMapName {
		//fmt.Println(findMap.Key, findMap.Value)
		if findMap.Key.(*ast.Ident).String() == mapKey && findMap.Value.(*ast.Ident).String() == mapValue {
			foundMapKeyValue = true
		}
	}

	return foundMapName && foundMapKeyValue
}

// Function for checking slice
func checkSlice(sliceName, sliceType string) bool {

	foundSliceName, foundSliceType := false, false

	f := readFile()
	var findSlice *ast.ArrayType
	//	var

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				//	fmt.Println(spec, reflect.TypeOf(spec).String())
				if reflect.TypeOf(spec).String() == "*ast.ValueSpec" {
					ValueSpec := spec.(*ast.ValueSpec)
					if reflect.TypeOf(ValueSpec.Type).String() == "*ast.ArrayType" {
						for _, arr := range ValueSpec.Names {
							if arr.Name == sliceName {
								foundSliceName = true
								findSlice = ValueSpec.Type.(*ast.ArrayType)
								break
							}
						}
					}
				}
			}
		}
	}

	if foundSliceName && findSlice.Elt.(*ast.Ident).String() == sliceType {
		foundSliceType = true

	}

	return foundSliceName && foundSliceType
}

// Function for checking said interface
func checkInterface(interfaceName string) (bool, bool) {

	var foundAInterface bool
	var foundVehicleInterface bool

	f := readFile()

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.TypeSpec" {
					typeSpec := spec.(*ast.TypeSpec)
					if reflect.TypeOf(typeSpec.Type).String() == "*ast.InterfaceType" {
						foundAInterface = true
						if typeSpec.Name.Name == interfaceName {
							foundVehicleInterface = true
						}
					}
				}
			}
		}
	}
	return foundAInterface, foundVehicleInterface
}

// Function for checking said struct
func checkStruct(structName string) (bool, bool) {

	var foundAStruct bool
	var foundNamedStruct bool

	f := readFile()

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.TypeSpec" {
					typeSpec := spec.(*ast.TypeSpec)
					if reflect.TypeOf(typeSpec.Type).String() == "*ast.StructType" {
						foundAStruct = true
						if typeSpec.Name.Name == structName {
							foundNamedStruct = true
						}
					}
				}
			}
		}
	}
	return foundAStruct, foundNamedStruct
}

// Function for checking struct fields
func checkStructProperties(structName, fieldName, fieldType string) bool {
	var targetStruct *ast.TypeSpec

	f := readFile()

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			for _, spec := range genDecl.Specs {
				if reflect.TypeOf(spec).String() == "*ast.TypeSpec" {
					typeSpec := spec.(*ast.TypeSpec)
					if reflect.TypeOf(typeSpec.Type).String() == "*ast.StructType" {
						if typeSpec.Name.Name == structName {
							targetStruct = typeSpec
							break
						}
					}
				}
			}
		}
	}

	if targetStruct == nil {
		return false
	}

	targetStructType := targetStruct.Type.(*ast.StructType)
	for _, field := range targetStructType.Fields.List {
		for _, name := range field.Names {
			if name.Name == fieldName {
				switch reflect.TypeOf(field.Type).String() {
				case "*ast.Ident":
					fType := field.Type.(*ast.Ident)
					return fType.Name == fieldType

				case "*ast.ArrayType":
					if !strings.Contains(fieldType, "[]") {
						return false
					}
					aType := field.Type.(*ast.ArrayType)
					elt := aType.Elt.(*ast.Ident)
					return elt.Name == strings.ReplaceAll(fieldType, "[]", "")
				}
			}
		}
	}
	return false
}
