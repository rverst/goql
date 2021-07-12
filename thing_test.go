package goql

import (
	"strings"
	"testing"
	"time"
)

func TestThings_CheckMap(t1 *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name    string
		query  	string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "Test 1",
			query:   `title == "foo BAR"`,
			args:    args{
				m: map[string]interface{}{
					"title": "foo bar",
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test 2",
			query:   `title === "foo BAR"`,
			args:    args{
				m: map[string]interface{}{
					"title": "foo bar",
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test 3",
			query:   `title == "foo BAR" && not disabled == t`,
			args:    args{
				m: map[string]interface{}{
					"title": "foo bar",
					"disabled": false,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test 4",
			query:   `title == "foo BAR" && not disabled == t`,
			args:    args{
				m: map[string]interface{}{
					"title": "foo bar",
					"disabled": true,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test 5",
			query:   `title == "foo BAR" && age > 42`,
			args:    args{
				m: map[string]interface{}{
					"title": "foo bar",
					"age": 43,
				},
			},
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test 6",
			query:   `title == "foo BAR" && age > 42`,
			args:    args{
				m: map[string]interface{}{
					"title": "foo bar",
					"age": 39,
				},
			},
			want:    false,
			wantErr: false,
		},
		{
			name:    "Test 7",
			query:   `title == "foo BAR" && date > '2008-01-02'`,
			args:    args{
				m: map[string]interface{}{
					"title": "foo bar",
					"date": time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			p := NewParser(strings.NewReader(tt.query))
			t, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t1.Errorf("CheckMap() parser error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if t == nil {
				t1.Errorf("CheckMap() parser error")
				return
			}

			t.AddDateFormat("2006-01-02")

			got, err := t.CheckMap(tt.args.m)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CheckMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("CheckMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_things_CheckStruct(t1 *testing.T) {

	type TestStruct struct {
		Title string
		Age int
	}

	type args struct {
		s interface{}
	}
	tests := []struct {
		name    string
		query   string
		args    interface{}
		want    bool
		wantErr bool
	}{
		{ name: "Test 1", query: `Title == "foo bar" & Age < 43`, args: &TestStruct{
			Title: "foo bar",
			Age:   42,
		}, want: true,wantErr: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			p := NewParser(strings.NewReader(tt.query))
			t, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t1.Errorf("CheckStruct() parser error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if t == nil {
				t1.Errorf("CheckStruct() parser error")
				return
			}

			t.AddDateFormat("2006-01-02")

			got, err := t.CheckStruct(tt.args)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CheckStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t1.Errorf("CheckStruct() got = %v, want %v", got, tt.want)
			}
		})
	}
}