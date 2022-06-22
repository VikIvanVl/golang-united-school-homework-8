package main

import (
	"encoding/json"
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

const perm = 0755
const fileNameArgs = "fileName"
const itemArgs = "item"
const idArgs = "id"
const operationArgs = "operation"

func Perform(args Arguments, writer io.Writer) error {
	if _, ok := args[operationArgs]; !ok || args[operationArgs] == "" {
		return errors.New("-operation flag has to be specified")
	}

	if _, ok := args[fileNameArgs]; !ok || args[fileNameArgs] == "" {
		return errors.New("-fileName flag has to be specified")
	}

	file, err := os.OpenFile(args[fileNameArgs], os.O_RDWR|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	switch args[operationArgs] {
	case "list":
		list(file, writer)
		break
	case "add":
		if _, ok := args[itemArgs]; !ok || args[itemArgs] == "" {
			return errors.New("-item flag has to be specified")
		}

		err = add(args[itemArgs], file, writer)
		if err != nil {
			_, err := writer.Write([]byte(err.Error()))
			if err != nil {
				return err
			}
		}
		break
	case "findById":
		if _, ok := args[idArgs]; !ok || args[idArgs] == "" {
			return errors.New("-id flag has to be specified")
		}

		user, err := findById(args[idArgs], file)
		if err != nil {
			return err
		}

		if user == nil {
			_, err := writer.Write([]byte(""))
			if err != nil {
				return err
			}
		} else {
			b, _ := json.Marshal(user)
			_, err := writer.Write(b)
			if err != nil {
				return err
			}
		}
		break
	case "remove":
		if _, ok := args[idArgs]; !ok || args[idArgs] == "" {
			return errors.New("-id flag has to be specified")
		}

		user, err := findById(args[idArgs], file)
		if err != nil {
			return err
		}
		if user == nil {
			_, err := writer.Write([]byte(fmt.Errorf("Item with id %s not found", args[idArgs]).Error()))
			if err != nil {
				return err
			}
		} else {
			err := remove(user, file)
			if err != nil {
				return err
			}
		}
		break
	default:
		return fmt.Errorf("Operation %s not allowed!", args[operationArgs])
	}

	return nil
}

func list(file *os.File, writer io.Writer) {
	b, _ := ioutil.ReadAll(file)
	_, err := writer.Write(b)
	if err != nil {
		return
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return
	}
}

func remove(user *User, file *os.File) error {
	defer func(file *os.File, offset int64, whence int) {
		_, err := file.Seek(offset, whence)
		if err != nil {

		}
	}(file, 0, 0)
	byteVal, _ := ioutil.ReadAll(file)

	var users []User
	_ = json.Unmarshal(byteVal, &users)
	for k, u := range users {
		if u.Id == user.Id {
			users = append(users[:k], users[k+1:]...)
		}
	}

	list, _ := json.Marshal(users)
	err := file.Truncate(0)
	if err != nil {
		return err
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.Write(list)
	if err != nil {
		return err
	}

	return nil
}

func add(item string, file *os.File, writer io.Writer) error {
	user := &User{}
	err := json.Unmarshal([]byte(item), user)
	if err != nil {
		return err
	}

	u, err := findById(user.Id, file)
	if u != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("Item with id %s already exists", user.Id)
	}

	byteVal, _ := ioutil.ReadAll(file)
	var users []User
	_ = json.Unmarshal(byteVal, &users)

	users = append(users, *user)
	byteUsers, _ := json.Marshal(users)

	_, err = file.Write(byteUsers)
	if err != nil {
		return err
	}

	list(file, writer)

	return nil
}

func findById(id string, file *os.File) (*User, error) {
	var usersVal []User

	defer func(file *os.File, offset int64, whence int) {
		_, err := file.Seek(offset, whence)
		if err != nil {

		}
	}(file, 0, 0)

	byteVal, _ := ioutil.ReadAll(file)
	err := json.Unmarshal(byteVal, &usersVal)
	if err != nil {
		return nil, err
	}

	for _, user := range usersVal {
		if user.Id == id {
			return &user, nil
		}
	}

	return nil, nil
}

func parseArgs() Arguments {
	id := flag.String("id", "", "id for finding")
	operation := flag.String(operationArgs, "",
		"list, add, findById, remove")
	item := flag.String(itemArgs, "", "json example: { id: \"1\",\n    email: \"test@test.com\",\n    age: 31\n}")
	fileName := flag.String("fileName", "", "json filename")

	flag.Parse()

	return Arguments{
		idArgs:        *id,
		operationArgs: *operation,
		itemArgs:      *item,
		fileNameArgs:  *fileName,
	}
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
