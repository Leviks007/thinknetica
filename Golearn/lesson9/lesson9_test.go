package lesson9

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestMaxAge(t *testing.T) {
	type args struct {
		user []AgeGet
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{user: []AgeGet{
				NewCustomer("ivan", 18),
				NewCustomer("john", 50),
				NewEmployee("alice", 78),
				NewEmployee("alice", 98),
			},
			},

			want: 98,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.args.user...); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

type bufferWriter struct {
	buf *bytes.Buffer
}

func (w *bufferWriter) Write(p []byte) (int, error) {
	return w.buf.Write(p)
}

func TestWriteStrings(t *testing.T) {
	tests := []struct {
		name     string
		writer   io.Writer
		input    []interface{}
		expected string
	}{
		{
			name:     "Test Case 1: Write to Buffer",
			writer:   &bufferWriter{buf: bytes.NewBuffer([]byte{})},
			input:    []interface{}{"Hello, world!", 1, 2, []string{"привет", "Добрый день!"}},
			expected: "Hello, world!12[привет Добрый день!]",
		},
		{
			name:     "Test Case 2: Write to File",
			writer:   createFileWriter("output.txt"),
			input:    []interface{}{"Hello, world!", 2, 3, []string{"привет", "Добрый день!"}},
			expected: "Hello, world!23[привет Добрый день!]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch w := tt.writer.(type) {
			case *bufferWriter:
				WriteStrings(tt.writer, tt.input...)
				if w.buf.String() != tt.expected {
					t.Errorf("Unexpected output. Got: %q, want: %q", w.buf.String(), tt.expected)
				}

			case *os.File:
				WriteStrings(tt.writer, tt.input...)
				fileContent, err := ioutil.ReadFile("output.txt")
				if err != nil {
					t.Fatalf("Failed to read file: %v", err)
				}
				if string(fileContent) != tt.expected {
					t.Errorf("Unexpected output in file. Got: %q, want: %q", string(fileContent), tt.expected)
				}

			}
		})
	}
}

func createFileWriter(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to create file: %v", err))
	}
	return file
}

func TestMaxAge_2_1(t *testing.T) {
	type args struct {
		users []interface{}
	}
	ivan := NewEmployee_2("Ivan", 18)
	john := NewEmployee_2("John", 50)
	alice := NewCustomer_2("Alice", 78)

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test Case 1",
			args: args{users: []interface{}{
				ivan,
				john,
				alice,
			},
			},

			want: alice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge_2(tt.args.users...); got != tt.want {
				t.Errorf("MaxAge_2_2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxAge_2(t *testing.T) {
	type args struct {
		users []interface{}
	}
	ivan := NewEmployee_2("Ivan", 18)
	john := NewEmployee_2("John", 50)
	alice := NewCustomer_2("Alice", 78)

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "Test Case 1",
			args: args{users: []interface{}{
				ivan,
				john,
				alice,
			},
			},

			want: alice,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge_2(tt.args.users...); got != tt.want {
				t.Errorf("MaxAge_2_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
