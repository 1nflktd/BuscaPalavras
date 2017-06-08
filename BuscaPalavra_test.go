package main

import "testing"

func TestReadFile(t *testing.T) {
	words := readFile("pt-BR.dic")
	if (len(words) == 0) {
		t.Error("Nao carregou arquivo \"pt-BR.dic\" ", words)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("A funcao nao entrou em panico")
		}
	}()
	words = readFile("pt-BR2.dic")
}

type SearchWordsTests struct {
	search Search
	result []string
}

func TestSearchWords(t *testing.T) {
	words := readFile("pt-BR.dic")

	searchs := []SearchWordsTests{
		{ 
			Search{ 
				map[int]string{ 0: "h", 7: "e", 5: "q" },
				8,
			},
			[]string{ "henrique" },
		},
		{ 
			Search{ 
				map[int]string{ 1: "h", 7: "e", 5: "q" },
				8,
			},
			[]string{ "henrique", "chilique" },
		},
	}

	for _, search := range searchs {
		wordsFound := searchWords(&words, search.search)
		var wordsResult []string

		for _, w := range wordsFound {
			t.Log("w ", w)
			for _, s := range search.result {
				t.Log("s ", s)
				if w == s {
					wordsResult = append(wordsResult, s)
					break
				}
			}
		}

		if (len(wordsResult) != len(search.result)) {
			t.Error("So achou as palavras: ", wordsResult, " de ", search.result)
		}
	}

}
