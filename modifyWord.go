package hangmanpackage

import (
	"fmt"
	"strings"
)

func ModifyWord(wordtofind, actualword, letter string, attempts int) (string, int) {
	fmt.Println("ModifyWord called.")
	// Vérifie si la lettre est présente dans le mot
	if strings.Contains(wordtofind, letter) {
		fmt.Println("Letter is in the word.")
		// Met à jour actualword en ajoutant la lettre à la position correspondante
		for i, r := range wordtofind {
			if r == rune(letter[0]) {
				actualword = actualword[:i] + letter + actualword[i+1:]
			}
		}
		fmt.Println(attempts)
		fmt.Println(actualword)
		return actualword, attempts
	}
	attempts -= 1
	fmt.Println("Letter is not in the word.")
	fmt.Println(attempts)
	fmt.Println(actualword)
	return actualword, attempts
}
