package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Search struct {
	letters map[int]string
	lenght int
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if (err != nil) {
		panic(err)
	}
	defer file.Close()

	var words []string;
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordSplit := strings.Split(scanner.Text(), "/")
		if (len(wordSplit) == 0) {
			continue
		}
		words = append(words, strings.TrimSpace(strings.ToLower(wordSplit[0])))
	}

	return words
}

func readInput() Search {
	var search Search
	search.letters = make(map[int]string)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Quantidade de letras da palavra buscada")
	l, _ := reader.ReadString('\n')
	lenght, _ := strconv.Atoi(strings.TrimSpace(l))
	search.lenght = lenght

	fmt.Println("Digite as posicoes e as letras, ie. 1-b 3-c 4-h")
	letters, _ := reader.ReadString('\n')

	lettersSplit := strings.Split(letters, " ")
	for _, w := range lettersSplit {
		wSplit := strings.Split(w, "-")
		if (len(wSplit) != 2) {
			continue
		}

		index, _ := strconv.Atoi(wSplit[0])
		search.letters[index] = wSplit[1]
	}

	return search
}

func searchWords(words *[]string, search Search) []string {
	var wordsFound []string;

	WORD_LOOP:
	for _, word := range *words {
		if (len(word) != search.lenght) {
			continue
		}

		letters := strings.Split(word, "")
		for index, letter := range search.letters {
			if len(letters) <= index || letters[index] != letter {
				continue WORD_LOOP
			}
		}

		wordsFound = append(wordsFound, word)
	}

	return wordsFound;
}

func main() {
	words := readFile("pt-BR.dic")
	search := readInput()

	wordsFound := searchWords(&words, search)

	fmt.Println("Palavras achadas:")
	for _, w := range wordsFound {
		fmt.Println(w)
	}
}
