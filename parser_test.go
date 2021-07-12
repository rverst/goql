package goql

import (
	"strings"
	"testing"
)

func TestParser_Parse(t *testing.T) {

	tests := []struct {
		name    string
		query   string
		want    *things
		wantErr bool
	}{
		{name: "Test 1", query: `title == "foo\\bar"`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `foo\bar`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test 2", query: `title == "foo\\bar" & bla != fasel`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `foo\bar`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}, {
			Link:       LNK_AND,
			Key:        "bla",
			Expression: `fasel`,
			ExprType:   LITERAL,
			Operator:   OP_NEQI,
		}}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := NewParser(strings.NewReader(tt.query))
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.want != nil {
				t.Errorf("Parse() got = \"%s\", want \"%s\"", got, tt.want)
			} else if got != nil && !got.Equals(tt.want) {
				t.Errorf("Parse() got = \"%s\", want \"%s\"", got, tt.want)
			}
		})
	}
}

func TestParser_ParseType(t *testing.T) {
	type fields struct {
		s   *Scanner
		buf struct {
			token   Token
			literal string
			n       int
		}
	}
	tests := []struct {
		name    string
		query  string
		want    *things
		wantErr bool
	}{
		{name: "Test string 1", query: `title == string`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `string`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test string 2", query: `title == "42"`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `42`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test string 3", query: `title == "42.1"`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `42.1`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test string 4", query: `title == "hello \"world\""`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `hello "world"`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test string 5", query: `title == "42\\1337"`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `42\1337`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test string 6", query: `title == "true"`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `true`,
			ExprType:   LITERAL,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test int 1", query: `title == 42`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `42`,
			ExprType:   INTEGER,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test int 2", query: `title == -1337`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `-1337`,
			ExprType:   INTEGER,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test float 1", query: `title == 3.14`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `3.14`,
			ExprType:   FLOAT,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test float 2", query: `title == -3.14`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `-3.14`,
			ExprType:   FLOAT,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test bool 1", query: `title == true`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `true`,
			ExprType:   BOOLEAN,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test bool 2", query: `title == false`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `false`,
			ExprType:   BOOLEAN,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test bool 3", query: `title == TRUE`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `TRUE`,
			ExprType:   BOOLEAN,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test bool 4", query: `title == FALSE`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `FALSE`,
			ExprType:   BOOLEAN,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test bool 4", query: `title == t`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `t`,
			ExprType:   BOOLEAN,
			Operator:   OP_EQI,
		}}}, wantErr: false},
		{name: "Test time 1", query: `title == '2006-01-02'`, want: &things{things: []*Thing{{
			Link:       EOF,
			Key:        "title",
			Expression: `2006-01-02`,
			ExprType:   TIME,
			Operator:   OP_EQI,
		}}}, wantErr: false},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewParser(strings.NewReader(tt.query))
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.want != nil {
				t.Errorf("Parse() got = \"%s\", want \"%s\"", got, tt.want)
			} else if got != nil && !got.Equals(tt.want) {
				t.Errorf("Parse() got = \"%s\", want \"%s\"", got, tt.want)
			}

		})
	}
}