package main

import (
	"fmt"
	"strings"
	"sync"
)

type Vowel struct {
	Vowels []string
	Word   string
}

var vowels = []string{"a", "e", "i", "o", "u"}
var wg sync.WaitGroup

func main() {
	taskChan := make(chan string)
	workerNumber := 4
	resultChan := make(chan Vowel)

	words := []string{"Serendipity", "Whimsical", "Galactic", "Espresso", "Quantum", "Cacophony", "Peregrinate", "Nebulous", "Mellifluous", "Bibliopole",
		"Ineffable", "Zephyr", "Ephemeral", "Luminescent", "Aureate", "Quixotic", "Obfuscate", "Juxtapose", "Cynosure", "Effervescent",
		"Halcyon", "Petrichor", "Sonorous", "Labyrinthine", "Ethereal", "Bucolic", "Mellifluent", "Surreptitious", "Luminous", "Epiphany",
		"Panoply", "Sesquipedalian", "Peregrine", "Halcyon", "Ephemeral", "Quintessential", "Serendipitous", "Resplendent", "Efflorescence",
		"Enigmatic", "Evanescent", "Ineffable", "Mellifluous", "Vorfreude", "Saudade", "Sonder", "Crepuscule", "Aurora", "Effulgent",
		"Quasar", "Mystique", "Vivid", "Symbiosis", "Ebullient", "Penumbra", "Benevolent", "Murmuration", "Bivouac", "Phantasmagoria",
		"Quintessence", "Halcyon", "Ephemeral", "Nostalgia", "Luminescence", "Bibliophile", "Incandescent", "Obfuscation", "Gossamer",
		"Labyrinth", "Incantation", "Serenity", "Ineffable", "Quotidian", "Halcyon", "Ephemeral", "Surreptitious", "Aurora", "Peregrinate",
		"Luminescent", "Ethereal", "Cacophony", "Bucolic", "Quintessential", "Nebulous", "Mellifluous", "Aureate", "Halcyon", "Ephemeral", "Ufuk",
	}

	for i := 0; i < workerNumber; i++ {
		wg.Add(1)
		go findVowels(taskChan, resultChan)
	}

	go func() {
		for i := 0; i < len(words); i++ {
			fmt.Println(">>>>>2342342342")
			taskChan <- words[i]
		}
		fmt.Println("doneee")
		close(taskChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()
	fmt.Println("workde done")
	for result := range resultChan {
		fmt.Println(result)
	}

}

func findVowels(taskChan chan string, resultChan chan Vowel) {
	fmt.Println(">>>>>")
	defer wg.Done()
	for value := range taskChan {
		fmt.Println(">>>>>processing")
		foundedVowels := []string{}
		for _, vowel := range vowels {
			if strings.Contains(value, vowel) {
				foundedVowels = append(foundedVowels, vowel)
			}
		}
		vowel := Vowel{
			Vowels: foundedVowels,
			Word:   value,
		}
		fmt.Println("sdlkhjfjlksdfh", vowel)
		resultChan <- vowel
	}
}
