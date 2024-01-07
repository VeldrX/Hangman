package hangman

import (
	hangman "command-line-arguments/home/mathisvassy/Devellopement/hangclone/hangman/hangmanadvanced.go"
	"fmt"
	"os"
)

type HangManData struct {
	Word         []rune // Word composed of '_', ex: H_ll_
	WordToFind   []rune // Final word chosen by the program at the beginning. It is the word to find
	Attempts     int    // Number of attempts left
	WrongLetters []rune
	Asccitype    string
}

func MakeStruct() *HangManData {

	Myhangman := new(HangManData)

	fileword := "words1.txt"
	Myhangman.WordToFind = hangman.GetWord(fileword, Myhangman)
	Myhangman.Word = hangman.GetInitLetters(Myhangman.WordToFind)
	Myhangman.Attempts = 10

	return Myhangman
}
func NotRunGame() {

	argu := os.Args
	if len(argu) < 2 {
		fmt.Println("il manque un argument")
		os.Exit(1)
	}

	Myhangman := MakeStruct()
	fmt.Println("!!! WELCOME  TO  THE  HANGED-MAN  GAME !!!")
	fmt.Println()
	Display(Myhangman.Word)

}

func Inputconv(imput string) []rune {

	// for Myhangman.Attempts > 0

	imput_in_rune := []rune(imput)
	return imput_in_rune
}

func Checkletter(imput_in_rune []rune, imputcorect bool, Myhangman *HangManData) {
	imput_in_rune, imputcorect = VerifyImput(imput_in_rune)
	var letterused bool
	if imputcorect {

		letterused = Boxletters(imput_in_rune, *Myhangman)
		if !letterused {
			AddLetter(Myhangman, imput_in_rune)

			compteur_egal := 0
			for i := 0; i < len(Myhangman.Word); i++ {
				if Myhangman.Word[i] == Myhangman.WordToFind[i] {
					compteur_egal += 1
				}
			}
			if compteur_egal == len(Myhangman.WordToFind) && Myhangman.Attempts > 0 {
				fmt.Println("BRAVO vous avez Gagné")
				os.Exit(0)
			}
		}
	}
}
