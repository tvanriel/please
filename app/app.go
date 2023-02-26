package app

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PrintTodo(todoFilename string) error {
	file, err := os.OpenFile(todoFilename, os.O_RDONLY, 0644)

	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}

	var bf bytes.Buffer

	io.Copy(&bf, file)
	lines := strings.Split(strings.TrimRight(bf.String(), "\n\t \r"), "\n")

	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println(textStyle.Render("You have nothing left todo. Use "))
		fmt.Println(titleStyle.Render("$ please add <item>"))
		fmt.Println(textStyle.Render("to add something"))
		return nil
	}

	fmt.Println(titleStyle.Render("Heres your todo list!"))

	renderTodoList(lines)

	return nil

}

func Add(todoFilename string, item string) error {
	if len(strings.TrimSpace(item)) == 0 {
		fmt.Println(titleStyle.Render("What do you want me to write down for you?"))
		return nil
	}

	file, err := os.OpenFile(todoFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}
	fmt.Println(titleStyle.Render("Right away!"))

	defer file.Close()
	_, err = file.WriteString(" - [ ] " + item + "\n")
	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}

	fmt.Println(successStyle.Render("I've written it down on your todo list."))
	return nil
}

func Finish(filename string, toFinish string) error {
	idToFinish, err := strconv.Atoi(toFinish)

	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}
	fmt.Println(successStyle.Render("Good job!"))

	var bf bytes.Buffer

	io.Copy(&bf, file)
	lines := strings.Split(strings.TrimRight(bf.String(), "\n\t \r"), "\n")

	if idToFinish == 0 || idToFinish > len(lines) {
		fmt.Println(dangerStyle.Render("I don't see that item on your list!"))
		return err
	}
	lines[idToFinish-1] = strings.Replace(lines[idToFinish-1], " - [ ]", " - [x]", 1)
	file.Seek(0, 0)
	file.WriteString(strings.Join(lines, "\n"))

	return nil
}

func Delete(filename string, toFinish string) error {

	pos, err := strconv.Atoi(toFinish)

	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}

	file, err := os.OpenFile(filename, os.O_EXCL|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}

	var bf bytes.Buffer

	io.Copy(&bf, file)
	lines := strings.Split(strings.TrimRight(bf.String(), "\n\t \r"), "\n")

	if pos == 0 || pos > len(lines) {
		fmt.Println(dangerStyle.Render("I'm sorry but I don't see that item on your list!"))
		return err
	}

	lines = append(lines[:pos-1], lines[pos:]...)

	defer file.Close()
	file.Seek(0, 0)
	file, err = os.OpenFile(filename, os.O_EXCL|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println(dangerStyle.Render(err.Error()))
		return err
	}
	file.WriteString(strings.Join(lines, "\n"))
	fmt.Println(successStyle.Render("Item deleted!"))
	if len(lines) == 0 {
		fmt.Println(successStyle.Render("Wow! You're all done!"))
	} else {
		fmt.Println(textStyle.Render("Your first1 item is:"))
		fmt.Println(titleStyle.Render(lines[0]))
	}

	return nil
}
