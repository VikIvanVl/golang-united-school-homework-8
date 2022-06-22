package main

import (
	"flag"
	"io"
	"os"
)

type Arguments map[string]string

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}

func Perform(args Arguments, writer io.Writer) error {

	return nil
}

func list(file *os.File, writer io.Writer) {
}

func remove(user *User, file *os.File) error {

	return nil
}

func add(item string, file *os.File, writer io.Writer) error {
	return nil
}

func find(id string, file *os.File) (*User, error) {

	return nil, nil
}

func parseArgs() Arguments {
	id := flag.String("id", "", "finding")
	operation := flag.String("operation", "", "list, findById, remove")
	item := flag.String("item", "", "json example: { id: \"1\",\n    email: \"test@test.com\",\n    age: 31\n}")
	fileName := flag.String("fileName", "", "json filename")

	flag.Parse()

	return Arguments{
		"id":        *id,
		"operation": *operation,
		"item":      *item,
		"fileName":  *fileName,
	}
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
