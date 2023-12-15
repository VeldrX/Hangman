package hangman

import (
	"fmt"
	"os"
)

// type HangManData struct {
// 	Word         []rune // Word composed of '_', ex: H_ll_
// 	WordToFind   []rune // Final word chosen by the program at the beginning. It is the word to find
// 	Attempts     int    // Number of attempts left
// 	WrongLetters []rune
// 	Asccitype    string
// }

func Hangstartstop() {
	imputcorect := true
	Myhangman := new(HangManData)
	var saved int
	var savefile string
	argu := os.Args
	if len(argu) < 2 {
		fmt.Println("il manque un argument")
		os.Exit(1)
	}
	fileword := argu[1]
	if len(argu) < 4 {
		fmt.Println("il manque un argument")
		os.Exit(1)
	}
	if argu[2] == "--startWith" {
		savefile = argu[3]
		fmt.Println("dddd")
		fmt.Println(savefile)
		saved = 1
	} else if argu[2] == "--letterFile" {
		Myhangman.Asccitype = argu[3]
	} else {
		fmt.Println("mauvais argument")
		os.Exit(1)
	}
	if saved == 0 {
		Myhangman.WordToFind = GetWord(fileword, Myhangman)
		Myhangman.Word = GetInitLetters(Myhangman.WordToFind)
		Myhangman.Attempts = 10
	}

	fmt.Println("!!! WELCOME  TO  THE  HANGED-MAN  GAME !!!")
	fmt.Println()

	if len(argu) == 4 && argu[2] == "--startWith" {
		var erp error
		*Myhangman, erp = LoadFromFile(savefile)
		if erp != nil {
			fmt.Println("error")
			os.Exit(1)
		}
	}

	Asciiart(Myhangman.Word, Myhangman.Asccitype)

	var imput string

	for Myhangman.Attempts > 0 {
		fmt.Println("")
		fmt.Print("Type a Letter : ")
		fmt.Scan(&imput)

		if imput == "stop" || imput == "STOP" {
			Myhangman.SaveToFile("save.txt")
			os.Exit(0)
		} else {
			imput_in_rune := []rune(imput)

			imput_in_rune, imputcorect = VerifyImput(imput_in_rune)
			var letterused bool
			if imputcorect {
				attemptbefor := Myhangman.Attempts
				letterused = Boxletters(imput_in_rune, *Myhangman)
				if !letterused {
					AddLetter(Myhangman, imput_in_rune)
					Asciiart(Myhangman.Word, Myhangman.Asccitype)
					if attemptbefor != Myhangman.Attempts {
						fmt.Println("")
						fmt.Println("")

						Josedisplay(*Myhangman)
						Display(Myhangman.WrongLetters)
					}
					fmt.Println()
					if Myhangman.Attempts > 0 {
						fmt.Println(Myhangman.Attempts)
					} else {
						fmt.Println(0)
						fmt.Println()
						fmt.Println()
						Asciiart(Myhangman.WordToFind, Myhangman.Asccitype)
						fmt.Println("Vous avez perdu")
					}

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
				} else {
					fmt.Println("Une lettre est déjà utilisez")
				}
			} else {
				fmt.Println()
				fmt.Println("Erreur il faut des lettres")
			}
		}
	}
}
