package module3

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"path"
	"reflect"
	"runtime"
)

var ifBlock, funcBlock, methodBlock *ast.BlockStmt = nil, nil, nil
var mainForStmt, forBlock *ast.RangeStmt = nil, nil

// ------------------------------------- Compute functions -------------------------------

//Function for reading the file
func readFile() *ast.File {

	_, currentFile, _, _ := runtime.Caller(1)

	src, err := ioutil.ReadFile(path.Join(path.Dir(currentFile), "../vehicle.go"))

	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)

	if err != nil {
		panic(err)
	}
	return f
}

// Task 1
func checkFunc(funcName, paramName string) bool {

	f := readFile()
	foundFunc := false

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			if funcDecl.Name.String() == funcName {
				par := funcDecl.Type.Params.List
				if len(par) != 1 {
					return false
				}
				if par[0].Names[0].String() == paramName {
					foundFunc = true
					funcBlock = funcDecl.Body
				}

				break
			}
		}
	}
	return foundFunc
}

// Task 2
func checkAssignedValue(blck *ast.BlockStmt, varName string) bool {

	var foundVar = false

	if blck == nil {
		return false
	}

	for _, b := range blck.List {
		switch reflect.TypeOf(b).String() {
		case "*ast.AssignStmt":
			s := b.(*ast.AssignStmt)
			str := fmt.Sprintf("%v%v%v", s.Lhs[0], s.Tok, s.Rhs[0].(*ast.Ident).String())
			if str == varName {
				foundVar = true
			}
		}
	}
	return foundVar
}

// Task 3
func checkForStmt(blck *ast.BlockStmt, key, value, x string) bool {

	var foundForStmt = false

	if blck == nil {
		return false
	}

	for _, decl := range blck.List {
		if reflect.TypeOf(decl).String() == "*ast.RangeStmt" {
			b := decl.(*ast.RangeStmt)
			if b.Key.(*ast.Ident).String() == key && b.Value.(*ast.Ident).String() == value && b.X.(*ast.Ident).String() == x {
				forBlock = b
				foundForStmt = true
			}
		}
	}
	return foundForStmt
}

// Task 4 and 6
func checkIfStmt(blck *ast.BlockStmt, cond string) bool {

	var foundIfStmt = false

	if blck == nil {
		return false
	}

	for _, b := range blck.List {
		if reflect.TypeOf(b).String() == "*ast.IfStmt" {
			c := b.(*ast.IfStmt)
			if c.Cond != nil && reflect.TypeOf(c.Cond).String() == "*ast.BinaryExpr" {
				g := c.Cond.(*ast.BinaryExpr)
				str := fmt.Sprintf("%v%v%v", g.X, g.Op, g.Y)
				if str == cond {
					foundIfStmt = true
					ifBlock = c.Body
					break
				}
			} else if c.Cond != nil && reflect.TypeOf(c.Cond).String() == "*ast.UnaryExpr" {
				g := c.Cond.(*ast.UnaryExpr)
				str := fmt.Sprintf("%v%v", g.Op, g.X)
				if str == cond {
					foundIfStmt = true
					ifBlock = c.Body
					break
				}
			}
		}
	}
	return foundIfStmt
}

// Task 5 and 6
func checkStmts(blck *ast.BlockStmt, stmt string) bool {

	var foundStmt = false

	if blck == nil {
		return false
	}

	for _, b := range blck.List {
		if reflect.TypeOf(b).String() == "*ast.ExprStmt" { // ExprStmt
			c := b.(*ast.ExprStmt).X.(*ast.CallExpr).Fun.(*ast.SelectorExpr)
			str := fmt.Sprintf("%v.%v", c.X, c.Sel)
			if str == stmt {
				foundStmt = true
				break
			}

		} else if reflect.TypeOf(b).String() == "*ast.AssignStmt" { // AssignStmt
			s := b.(*ast.AssignStmt)
			str := fmt.Sprintf("%v%v%v", s.Lhs[0], s.Tok, s.Rhs[0].(*ast.Ident))
			if str == stmt {
				foundStmt = true
				break
			}
		}
	}
	return foundStmt
}

// Task 7 8 9
func checkMethod(methodName, name string) bool {

	f := readFile()
	foundMethod := false
	methodBlock = nil
	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			if funcDecl.Name.String() == methodName && len(funcDecl.Recv.List) == 1 {
				m := funcDecl.Recv.List[0]
				s := m.Type.(*ast.StarExpr)
				str := fmt.Sprintf("%v *%v", m.Names[0], s.X)
				if str == name {
					foundMethod = true
					methodBlock = funcDecl.Body
					break
				}
			}
		}
	}
	return foundMethod
}


// Task 10
func checkForWithinMain(funcName, key, value, x string) bool {
	foundFor := false

	f := readFile()
	var funcBody []ast.Stmt

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			if funcDecl.Name.String() == funcName {
				funcBody = funcDecl.Body.List
				break
			}
		}
	}
	for _, b := range funcBody {
		if reflect.TypeOf(b).String() == "*ast.RangeStmt" {
			mainForStmt = b.(*ast.RangeStmt)
			if mainForStmt.Key.(*ast.Ident).String() == key && mainForStmt.Value.(*ast.Ident).String() == value && mainForStmt.X.(*ast.Ident).String() == x {
				foundFor = true
				break
			}
		}
	}
	return foundFor
}

// Task 11: Check Switch for type
func checkSwitchType(blck *ast.RangeStmt, exp string) bool {
	var foundSwitch = false

	if blck == nil {
		return false
	}

	for _, b := range blck.Body.List {
		if reflect.TypeOf(b).String() == "*ast.TypeSwitchStmt" {
			s := b.(*ast.TypeSwitchStmt).Assign.(*ast.AssignStmt)
			t := s.Rhs[0].(*ast.TypeAssertExpr)
			str := fmt.Sprintf("%v%v%v.(type)", s.Lhs[0], s.Tok, t.X)
			if str == exp && len(b.(*ast.TypeSwitchStmt).Body.List) == 4 {
				foundSwitch = true
				break
			}
		}
	}
	return foundSwitch
}
