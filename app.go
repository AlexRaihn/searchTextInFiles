package main

import (
	"fmt"
	"os"
	"strings"
)

var Search = ""
var searchRune = []rune(Search)
var countWord int
var files []os.FileInfo

func main() {
	resultFiles, err := readDirectory()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	files = resultFiles
	installSearchWord()
}

func installSearchWord() {
	fmt.Println("Какое слово нужно подсчитать в файлах?")
	fmt.Scanln(&Search)

	if len(Search) > 0 {
		fmt.Println("Идёт поиск, пожалуйста подождите...")
		readFileInFor(0, len(files))
	} else {
		installSearchWord()
	}
}

func readDirectory() ([]os.FileInfo, error) {
	// Открываем текущую директорию
	dir, err := os.Open("./files")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer dir.Close()

	// Получаем список файлов и папок
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return files, err
	}
	return files, nil
}

func readFileInFor(index int, len int) {
	if index < len {
		readFile(index)
		index = index + 1
		readFileInFor(index, len)
	} else {
		fmt.Println("Задача выполнена, найдено совпадений: ", countWord)
	}
}

func readFile(index int) {
	file, err := os.ReadFile("./files/" + files[index].Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	refString := string(file)

	words := strings.Fields(refString)

	for i := 0; i < len(words); i++ {
		if words[i] == Search {
			countWord = countWord + 1
		}
	}
}
