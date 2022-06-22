package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   uint   `json:"age"`
}

func Perform(args Arguments, writer io.Writer) error {
	if _, ok := args["operation"]; !ok || args["operation"] == "" {
		return errors.New("-operation flag has to be specified")
	}

	if _, ok := args["fileName"]; !ok || args["fileName"] == "" {
		return errors.New("-fileName flag has to be specified")
	}

	file, err := os.OpenFile(args["fileName"], os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	switch args["operation"] {
	case "list":
		list(file, writer)
		break
	case "add":

		break
	case "findById":

		break
	case "remove":

		break
	default:
		return fmt.Errorf("Operation %s not allowed!", args["operation"])
	}

	return nil
}

func list(file *os.File, writer io.Writer) {
	b, _ := ioutil.ReadAll(file)
	writer.Write(b)
	file.Seek(0, 0)
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
	id := flag.String("id", "", "item ID for finding")
	operation := flag.String("operation", "",
		"list, add, findById, remove")
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
