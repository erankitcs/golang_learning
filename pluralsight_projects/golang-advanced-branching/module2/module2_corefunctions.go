package module2

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

var mainForStmt *ast.RangeStmt = nil
var forBlock, forWord *ast.RangeStmt = nil, nil
var ifBlock *ast.BlockStmt = nil

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

// Function for checking function
func checkFunc(funcName string) bool {

	f := readFile()
	foundFunc := false

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			if funcDecl.Name.String() == funcName {
				foundFunc = true
				break
			}
		}
	}

	return foundFunc

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

// function for checking declared variable within the blockStmt.
func checkVarDeclWithinFor(varName, varType string) bool {
	foundVar := false

	if mainForStmt == nil {
		return false
	}

	for _, decl := range mainForStmt.Body.List {
		if reflect.TypeOf(decl).String() == "*ast.DeclStmt" {
			s := decl.(*ast.DeclStmt).Decl.(*ast.GenDecl)
			if s.Tok.String() == "var" {
				for _, b := range s.Specs {
					a := b.(*ast.ValueSpec)
					if a.Names[0].String() == varName && a.Type.(*ast.Ident).String() == varType {
						foundVar = true
					}
				}
			}
		}
	}
	return foundVar
}

// Function for checking for statement
// This will be only called once as its the main for statement
// The block statement will be stored in a global variable
func checkMainForWithinFunc(funcName, key, value, x string) bool {
	foundForStmt, foundFunc := false, false

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
	for _, b := range funcBody {
		if reflect.TypeOf(b).String() == "*ast.RangeStmt" {
			mainForStmt = b.(*ast.RangeStmt)
			if mainForStmt.Key.(*ast.Ident).String() == key && mainForStmt.Value.(*ast.Ident).String() == value && mainForStmt.X.(*ast.SelectorExpr).Sel.String() == x {
				foundForStmt = true
			}
		}
	}
	return foundFunc && foundForStmt
}

func checkForStmt(mainForBlock *ast.RangeStmt, key, value, x string) bool {

	var foundForStmt = false

	if mainForBlock == nil {
		return false
	}

	for _, decl := range mainForBlock.Body.List {
		if reflect.TypeOf(decl).String() == "*ast.RangeStmt" {
			b := decl.(*ast.RangeStmt)
			if b.Key.(*ast.Ident).String() == key && b.Value.(*ast.Ident).String() == value && b.X.(*ast.SelectorExpr).Sel.String() == x {
				forBlock = b
				foundForStmt = true
			}
		}
	}
	return foundForStmt
}

func checkIfStmt(blck *ast.RangeStmt, leftInit, rightInit, cond string) bool {

	var foundIfStmt = false

	if blck == nil {
		return false
	}

	for _, b := range blck.Body.List {
		if reflect.TypeOf(b).String() == "*ast.IfStmt" {
			c := b.(*ast.IfStmt)
			if c.Init != nil {
				d := c.Init.(*ast.AssignStmt)
				e := d.Rhs[0].(*ast.CallExpr).Fun.(*ast.SelectorExpr).X
				if c.Cond != nil && reflect.TypeOf(c.Cond).String() == "*ast.BinaryExpr" {
					g := c.Cond.(*ast.BinaryExpr)
					str := fmt.Sprintf("%v(%v)%v%v", g.X.(*ast.CallExpr).Fun, g.X.(*ast.CallExpr).Args[0], g.Op, g.Y.(*ast.BasicLit).Value)
					if d.Lhs[0].(*ast.Ident).String() == leftInit && e.(*ast.Ident).String() == rightInit && str == cond {
						foundIfStmt = true
						ifBlock = c.Body
					}
				}
			}
		}
	}
	return foundIfStmt
}

func checkSetValues(blck *ast.BlockStmt, varName string) bool {
	var foundVar = false

	if blck == nil {
		return false
	}

	for _, b := range blck.List {
		switch reflect.TypeOf(b).String() {
		case "*ast.AssignStmt":
			s := b.(*ast.AssignStmt)
			str := fmt.Sprintf("%v%v%v", s.Lhs[0], s.Tok, s.Rhs[0].(*ast.BasicLit).Value)
			if str == varName {
				foundVar = true
			}

		case "*ast.IncDecStmt":
			s := b.(*ast.IncDecStmt)
			p := s.X.(*ast.SelectorExpr)
			str := fmt.Sprintf("%v.%v%v", p.X, p.Sel, s.Tok)
			if str == varName {
				foundVar = true
			}
		}
	}
	return foundVar
}

func checkForWithinIf(blck *ast.BlockStmt, key, value, x string) bool {
	var foundFor = false

	if blck == nil {
		return false
	}

	for _, b := range blck.List {
		if reflect.TypeOf(b).String() == "*ast.RangeStmt" {
			f := b.(*ast.RangeStmt)
			if f.Key.(*ast.Ident).String() == key && f.Value.(*ast.Ident).String() == value && f.X.(*ast.Ident).String() == x {
				foundFor = true
				forWord = f
			}
		}
	}
	return foundFor
}

func checkSwitchCalRating(blck *ast.RangeStmt, leftInit, rightInit, exp string) bool {
	var foundSwitch = false

	if blck == nil {
		return false
	}
	for _, b := range blck.Body.List {
		if reflect.TypeOf(b).String() == "*ast.SwitchStmt" {
			s := b.(*ast.SwitchStmt)
			if s.Init == nil || s.Tag == nil {
				return false
			}
			i := s.Init.(*ast.AssignStmt)
			j := i.Rhs[0].(*ast.CallExpr).Fun.(*ast.SelectorExpr)
			if i.Lhs[0].(*ast.Ident).String() == leftInit && j.X.(*ast.Ident).String() == rightInit && s.Tag.(*ast.Ident).String() == exp {
				if len(s.Body.List) == 4 { // 4 case statements
					foundSwitch = true
				}
			}
		}
	}
	return foundSwitch
}

func checkSwitchAddFeedback(blck *ast.BlockStmt) bool {
	var foundSwitch = false

	if blck == nil {
		return false
	}

	for _, b := range blck.List {
		if reflect.TypeOf(b).String() == "*ast.SwitchStmt" {
			s := b.(*ast.SwitchStmt)
			if s.Init != nil || s.Tag != nil || len(s.Body.List) != 3 {
				return false
			}
			foundSwitch = true
		}
	}
	return foundSwitch
}

func checkAppendRating(blck *ast.BlockStmt, varName string) bool {
	var foundVar = false

	if blck == nil {
		return false
	}

	for _, b := range blck.List {
		switch reflect.TypeOf(b).String() {
		case "*ast.AssignStmt":
			s := b.(*ast.AssignStmt)
			p := s.Lhs[0].(*ast.IndexExpr)
			q := p.Index.(*ast.SelectorExpr)
			str := fmt.Sprintf("%v[%v.%v]%v%v", p.X, q.X, q.Sel, s.Tok, s.Rhs[0])
			if str == varName {
				foundVar = true
			}
		}
	}
	return foundVar
}

// Function for checking generateRating() within main
func checkFuncGenerateRating(funcName string) bool {

	f := readFile()
	foundFunc := false

	for _, decl := range f.Decls {
		if reflect.TypeOf(decl).String() == "*ast.FuncDecl" {
			funcDecl := decl.(*ast.FuncDecl)
			if funcDecl.Name.String() == "main" {
				for _, b := range funcDecl.Body.List {
					if reflect.TypeOf(b).String() == "*ast.ExprStmt" {
						e := b.(*ast.ExprStmt).X.(*ast.CallExpr)
						if e.Fun.(*ast.Ident).String() == funcName {
							foundFunc = true
						}
					}
				}
			}
		}
	}
	return foundFunc
}

// Function for checking import statements
func checkImports(pkgName string) bool {

	f := readFile()
	foundPkg := false

	for _, decl := range f.Decls {

		if reflect.TypeOf(decl).String() == "*ast.GenDecl" {
			genDecl := decl.(*ast.GenDecl)
			if genDecl.Tok == token.IMPORT {
				for _, b := range genDecl.Specs {
					c := b.(*ast.ImportSpec)
					if c.Path.Value == pkgName {
						foundPkg = true
					}
				}
			}
		}
	}
	return foundPkg
}
