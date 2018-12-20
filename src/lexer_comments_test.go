package main

import (
	"strings"
	"testing"
)

func TestLexerSingleLineComments(t *testing.T) {
	str := `Rectangle {
    width: 100.0// Test
    height: 10 //--// A-ha-ha
    color: "red"
}/////////`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "width" {
		t.Errorf("%s is not IDENTIFIER(width) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE_LITERAL || lex != "100.0" {
		t.Errorf("%s is not DOUBLE_LITERAL(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "height" {
		t.Errorf("%s is not IDENTIFIER(height) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != INT_LITERAL || lex != "10" {
		t.Errorf("%s is not INT_LITERAL(10) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != STRING_LITERAL || lex != "red" {
		t.Errorf("%s is not STRING_LITERAL(red) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not RIGHT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}

func TestLexerMultilineComments(t *testing.T) {
	str := `Rectangle {
    width: 100.0/* Test
    height: 10 */
    color: "red"/**/
}`
	r := strings.NewReader(str)
	s := VmlScanner(r)
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "Rectangle" {
		t.Errorf("%s is not IDENTIFIER(Rectangle) token", lex)
	}
	if token, lex := s.Scan(); token != LEFT_CURLY_BRACKET {
		t.Errorf("%s is not LEFT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "width" {
		t.Errorf("%s is not IDENTIFIER(width) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != DOUBLE_LITERAL || lex != "100.0" {
		t.Errorf("%s is not DOUBLE_LITERAL(100.0) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != IDENTIFIER || lex != "color" {
		t.Errorf("%s is not IDENTIFIER(color) token", lex)
	}
	if token, lex := s.Scan(); token != COLON {
		t.Errorf("%s is not COLON token", lex)
	}
	if token, lex := s.Scan(); token != STRING_LITERAL || lex != "red" {
		t.Errorf("%s is not STRING_LITERAL(red) token", lex)
	}
	if token, lex := s.Scan(); token != NEW_LINE {
		t.Errorf("%s is not NEW_LINE token", lex)
	}
	if token, lex := s.Scan(); token != RIGHT_CURLY_BRACKET {
		t.Errorf("%s is not RIGHT_CURLY_BRACKET token", lex)
	}
	if token, lex := s.Scan(); token != EOF || lex != "" {
		t.Errorf("%s is not EOF token", lex)
	}
}
