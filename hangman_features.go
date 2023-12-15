package hangman

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
)

func GetWord(file string, man *HangManData) []rune {

	choosen_word := ""
	//-------------------------------------------------------//
	//---------- Creat a tab to stock every words -----------//
	//-------------------------------------------------------//
	fichtxt := file
	fichLecture, err := os.Open(fichtxt)
	if err != nil { //traiter une possible absence de fichier(une erreur)
		log.Printf("Error: %v", err)
		os.Exit(1)
	}

	fichScan := bufio.NewScanner(fichLecture)
	fichScan.Split(bufio.ScanLines)

	// ----------Création d'un tableau contenant tout les mots du fichier (.txt)----------//
	var tabwrds []string
	for fichScan.Scan() {
		tabwrds = append(tabwrds, fichScan.Text())
	}
	fichLecture.Close()

	//------------------random-------------------------//

	max := len(tabwrds) - 1
	random_number := rand.Intn(max)
	choosen_word = tabwrds[random_number]

	//word in tab//

	choosen_word_inrune := []rune(choosen_word)
	for i := 0; i < len(choosen_word_inrune); i++ {
		man.Word = append(man.Word, choosen_word_inrune[i])
	}
	return man.Word
}

func GetInitLetters(word []rune) []rune {

	lenword := len(word)
	base_display := []rune{}
	number_to_reveal := (lenword / 2) - 1
	verifletter := []int{}

	for i := 0; i < lenword; i++ {
		base_display = append(base_display, 95)
	}

	for number_to_reveal > 0 {

		max := len(word) - 1
		random_number := rand.Intn(max)

		checkletterusing := false

		//verif//

		for m := 0; m < len(verifletter); m++ {
			if verifletter[m] == random_number {
				checkletterusing = true
				break
			}
		}

		verifletter = append(verifletter, random_number)
		if !checkletterusing {
			number_to_reveal = number_to_reveal - 1
		}
		for j := 0; j < lenword; j++ {

			if word[random_number] == word[j] {
				if !checkletterusing {

					base_display[j] = word[random_number]
				}
			}
		}

	}

	return base_display
}

func AddLetter(hang *HangManData, letr []rune) ([]rune, int) {
	count := 0

	for j := 0; j < len(letr); j++ {
		letter_exist := false
		for i := 0; i < len(hang.WordToFind); i++ {
			if hang.WordToFind[i] == letr[j] {
				hang.Word[i] = letr[j]
				letter_exist = true
			}
		}
		if !letter_exist {
			count += 1
			hang.WrongLetters = append(hang.WrongLetters, letr[j])
		}
	}
	hang.Attempts = hang.Attempts - count

	return hang.Word, hang.Attempts
}

func Display(word []rune) {
	for i := 0; i < len(word); i++ {
		fmt.Print(string(word[i]))
	}
}

func VerifyImput(imput []rune) ([]rune, bool) {

	onlymin := true
	for m := 0; m < len(imput); m++ {
		if imput[m] >= 'A' && imput[m] <= 'Z' {
			imput[m] += 32
		}
		if imput[m] < 'A' || (imput[m] > 'Z' && imput[m] < 'a') || imput[m] > 'z' {
			onlymin = false
		}

	}
	return imput, onlymin
}

func Josedisplay(hang HangManData) {
	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()

	var joseascii []string
	scanner := bufio.NewScanner(file)
	joseascii = append(joseascii, "\n\n\n\n\n\n========")
	currentjose := ""
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		if len(line) == 0 {

			joseascii = append(joseascii, currentjose)
			currentjose = ""
		} else {
			currentjose += line + "\n"
		}

	}

	desiredIndex := 10 - hang.Attempts
	if desiredIndex >= 0 && desiredIndex < len(joseascii) {
		fmt.Print(joseascii[desiredIndex])
	} else {
		fmt.Println("Invalid josé index.")
	}
}

func Boxletters(imput []rune, hang HangManData) bool {
	letterused := false

	for i := 0; i < len(hang.WrongLetters); i++ {
		for j := 0; j < len(imput); j++ {
			if imput[j] == hang.WrongLetters[i] {
				letterused = true
				return letterused
			}
		}
	}

	return letterused
}

func Asciiart(word []rune, displaytype string) {

	file, err := os.Open(displaytype)
	if err != nil {
		fmt.Println("Error file:", err)

	}
	defer file.Close()

	var asciitab []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		asciitab = append(asciitab, scanner.Text())

	}

	for i := 0; i < 9; i++ {
		for j := 0; j < len(word); j++ {
			fmt.Print(asciitab[((int(word[j])-32)*9)+i])
		}
		fmt.Print("\n")
	}

}

func (hang *HangManData) SaveToFile(filename string) error {
	data, err := json.Marshal(hang)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func LoadFromFile(filename string) (HangManData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return HangManData{}, err
	}
	var hang HangManData
	if err := json.Unmarshal(data, &hang); err != nil {
		return HangManData{}, err
	}
	return hang, nil
}
