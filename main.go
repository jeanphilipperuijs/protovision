package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-demos/chalk"
)

func InputPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+"")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

var help_games_lines [20]string
var got_joshua = false

func clearscreen() {
	fmt.Print("\033[H\033[2J")
}

func joshua_chat() {
	chat := InputPrompt("\n")
	if chat == "Hello." {
		fmt.Println("\n\nHOW ARE YOU FEELING TODAY?")

	} else if chat == "I'm fine. How are you?" {
		fmt.Println("\n\nEXCELLENT.  IT'S BEEN A LONG TIME.  CAN YOU EXPLAIN")
		fmt.Println("THE REMOVAL OF YOUR USER ACCOUNT NUMBER ON 6/23/73?")

	} else if strings.Contains(chat, "People sometimes make mistak") {
		fmt.Printf("\n\nYES, THEY DO.")
		time.Sleep(time.Second * 5)
		fmt.Println("  SHALL WE PLAY A GAME?")

	} else if chat == "Love to. How about Global Thermonuclear War?" {
		fmt.Println("\n\nWOULDN'T YOU PREFER A GOOD GAME OF CHESS?")

	} else if chat == "Later. Let's play Global Thermonuclear War." {
		fmt.Println("\n\nFINE")

	} else {
		fmt.Println(chat)
	}
}

func logon_prompt() {
	message := strings.ToLower(InputPrompt("\nLOGON: "))

	if message == "help logon" {
		fmt.Printf("\nHELP NOT AVAILABLE\n")

	} else if message == "joshua" {
		got_joshua = true
		clearscreen()
		fmt.Printf("\nGREETINGS PROFESSOR FALKEN.\n")

	} else if message == "help games" {
		help_games_lines[0] = "\n"
		help_games_lines[1] = "`GAMES` REFERS TO MODELS, SIMULATIONS AND GAMES"
		help_games_lines[2] = "WHICH HAVE TACTICAL AND STRATEGIC APPLICATIONS."
		help_games_lines[3] = "\n"
		help_games_lines[4] = "List Games\n"
		help_games_lines[5] = "FALKEN'S MAZE"
		help_games_lines[6] = "BLACK JACK"
		help_games_lines[7] = "GIN RUMMY"
		help_games_lines[8] = "HEARTS"
		help_games_lines[9] = "BRIDGE"
		help_games_lines[10] = "CHECKERS"
		help_games_lines[11] = "CHESS"
		help_games_lines[12] = "POKER"
		help_games_lines[13] = "FIGHTER COMBAT"
		help_games_lines[14] = "GUERRILLA ENGAGEMENT"
		help_games_lines[15] = "DESERT WARFARE"
		help_games_lines[16] = "AIR-TO-GROUND ACTIONS"
		help_games_lines[17] = "THEATERWIDE TACTICAL WARFARE"
		help_games_lines[18] = "THEATERWIDE BIOTOXIC AND CHEMICAL WARFARE"
		help_games_lines[19] = "\nGLOBAL THERMONUCLEAR WAR"

		for i := 0; i < len(help_games_lines); i++ {
			fmt.Println(help_games_lines[i])
			if i == 19 {
				time.Sleep(time.Second * 2)
			}
			time.Sleep(time.Second / 1)
		}
		time.Sleep(time.Second * 5)

	} else {
		fmt.Printf("\n%s NOT AVAILABLE\n", strings.ToUpper(message))
	}
}
func main() {
	clearscreen()
	fmt.Println(chalk.CyanLight(), chalk.Bold())
	for {
		if !got_joshua {
			logon_prompt()
		} else {
			joshua_chat()
		}
	}
}
