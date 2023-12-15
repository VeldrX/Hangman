Hangman Game in Go,

Welcome to the Hangman game made in Golang! This program is a simple school project of the Hangman game, written in the Gollang programming language.

Instructions
Make sure you have Go installed on your system. If not, you can download it here.

Clone this repository to your local machine.

    
    git clone https://ytrack.learn.ynov.com/git/vmathis/Hangmans.git
    


Navigate to the game directory.

    
    cd main
    


Run the game:

    In the main.go file ,you can change the game parameters
    func main() {
	    hangman.Hangstartstop()
    }
    You can choose beetwin :
    hangman.Hangclassic()

        to run:
            go run main.go **name of the text file**
        the text file is the file where all words are line by line


    hangman.Hangadvanced()

        to run:
            go run main.go **name of the text file**


    hangman.Hangasciiart() 

        to run:
            go run main.go **name of the text file** --letterFile **AsciiArt type**
        AsciiArt type is the style of the Ascii Art
        Here you can choose beetwin:
            standard.txt


    hangman.Hangstartstop()

    to run:
        go run main.go **name of the text file** --letterFile **AsciiArt type**
    you can stop by the input :
    stop or STOP

    and you can restart where you where with:
        go run main.go **name of the text file** --startWith save.txt   
    
    
    


Enjoy 🙂


Features
Classic Hangman game.
Random selection of words to guess.
Display of the hangman's state during the game.
Management of already guessed letters.
Our friend José which is the stickman
The letters in Ascii Art
A start and stop

 







Author
[Mathis_Vassy]