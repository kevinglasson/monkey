package parser

import (
	"fmt"
	"testing"

	"github.com/kevinglasson/monkey/ast"
	"github.com/kevinglasson/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 260991;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram return nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf(
			"program.Statements does not contain 3 statements. got %d",
			len(program.Statements),
		)
	}

	// Create an slice of expected identifiers.
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	// For each expected identifier check that it is as expected in the
	// programs statement slice.
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	// If the parser did have errors then log them and fail!
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser errors: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	// Check the token's literal value is "let"
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	// Assert the Type of this Statement is a LetStatement
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement, got=%T", s)
		return false
	}

	fmt.Printf("%+v", letStmt)

	// Test that the Value of the Name (Identifier) of this Statement is as
	// expected.
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	// Test that the
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf(
			"letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name,
			letStmt.Name.TokenLiteral(),
		)
		return false
	}
	return true
}
