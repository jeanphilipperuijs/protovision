package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-demos/chalk"
	baudprint "ruijs.fr/protovision/BaudPrint"
)

type InputOutput struct {
	Input      string
	Output     []string
	PreAction  string
	PostAction string
}

var (
	baudrate         = 15
	help_games_lines [20]string
	got_joshua       = false

	use_chat_lines = true

	chat_lines_logon = []InputOutput{
		{
			Input:  "Help Logon",
			Output: []string{"\nHELP NOT AVAILABLE"},
		},
		{
			Input: "Help games",
			Output: []string{
				"\n", "`GAMES` REFERS TO MODELS, SIMULATIONS AND GAMES", "WHICH HAVE TACTICAL AND STRATEGIC APPLICATIONS.\n",
				"\n", "List Games\n",
				"FALKEN'S MAZE",
				"BLACK JACK",
				"GIN RUMMY",
				"HEARTS",
				"BRIDGE",
				"CHECKERS",
				"CHESS",
				"POKER",
				"FIGHTER COMBAT",
				"GUERRILLA ENGAGEMENT",
				"DESERT WARFARE",
				"AIR-TO-GROUND ACTIONS",
				"THEATERWIDE TACTICAL WARFARE",
				"THEATERWIDE BIOTOXIC AND CHEMICAL WARFARE",
				"\nGLOBAL THERMONUCLEAR WAR",
			},
		},
		{
			Input:      "Armageddon",
			Output:     []string{"IDENTIFICATION NOT RECOGNIZED BY SYSTEM", "--CONNECTION-TERMINATED__"},
			PostAction: "exit",
		},
		{
			Input:     "joshua",
			Output:    []string{"\nGREETINGS PROFESSOR FALKEN."},
			PreAction: "clear",
		},
	}
	chat_lines_joshua = []InputOutput{
		{
			Input:      "clear",
			Output:     []string{"\n\nBRB"},
			PostAction: "clear",
		},
		{
			Input:  "Hello.",
			Output: []string{"\n\nHOW ARE YOU FEELING TODAY?"},
		},
		{
			Input:  "I'm fine. How are you?",
			Output: []string{"\n\nEXCELLENT.  IT'S BEEN A LONG TIME.  CAN YOU EXPLAIN", "THE REMOVAL OF YOUR USER ACCOUNT NUMBER ON 6/23/73?"},
		},
		{
			Input:  "People sometimes make mistakes",
			Output: []string{"\n\nYES, THEY DO.  SHALL WE PLAY A GAME?"},
		},
		{
			Input:  "Love to. How about Global Thermonuclear War?",
			Output: []string{"\n\nWOULDN'T YOU PREFER A GOOD GAME OF CHESS?"},
		},
		{
			Input: "Later. Let's play Global Thermonuclear War.",
			Output: []string{"\n\nFINE", `
	,__                                                  _,
	\~\|  ~~---___              ,                          | \
	| Wash./ |   ~~~~~~~|~~~~~| ~~---,                VT_/,ME>
/~-_--__| |  Montana |N Dak\ Minn/ ~\~~/Mich.     /~| ||,'
|Oregon /  \         |------|   { WI / /~)     __-NY',|_\,NH
/       |Ida.|~~~~~~~~|S Dak.\    \   | | '~\  |_____,|~,-'Mass.
|~~--__ |    | Wyoming|____  |~~~~~|--| |__ /_-'Penn.{,~Conn (RI)
|   |  ~~~|~~|        |    ~~\ Iowa/  '-' |'~ |~_____{/NJ
|   |     |  '---------, Nebr.\----| IL|IN|OH,' ~/~\,|'MD (DE)
',  \ Nev.|Utah| Colo. |~~~~~~~|    \  | ,'~~\WV/ VA |
|Cal\    |    |       | Kansas| MO  \_-~ KY /'~___--\
',   \  ,-----|-------+-------'_____/__----~~/N Car./
	'_   '\|     |      |~~~|Okla.|    | Tenn._/-,~~-,/
	\    |Ariz.| New  |   |_    |Ark./~~|~~\    \,/S Car.
	~~~-'     | Mex. |     '~~~\___|MS |AL | GA /
		'-,_  | _____|          |  /   | ,-'---~\
			´~'~  \    Texas    |LA'--,~~~~-~~,FL\
					\/~\      /~~~'---'         |  \
						\    /                   \  |
						\  |                     '\'
							'~'`},
		},
		{
			Input:      "bye",
			Output:     []string{"\n\nHOPE TO SEE YOU SOON", "BYE"},
			PostAction: "exit",
		},
	}
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

func clearscreen() {
	fmt.Print("\033[H\033[2J")
}

func writeLine(input string) {
	baudprint.BaudPrint(input, int64(baudrate), 3, false, false)
	fmt.Printf("\n")
}

func check_input(prompt_input string, known_input string) bool {
	if len(prompt_input) > 2 {
		prompt_input_lower := strings.ToLower(prompt_input)
		known_input_lower := strings.ToLower(known_input)
		return prompt_input_lower == known_input_lower || strings.Contains(known_input_lower, prompt_input_lower) || strings.Contains(prompt_input_lower, known_input_lower)
	}
	return false
}

func chat_joshua() {
	prompt_input := InputPrompt("\n")

	gave_output := false
	for _, interaction := range chat_lines_joshua {

		if check_input(prompt_input, interaction.Input) {
			switch interaction.PreAction {
			default:
				break
			}
			for _, o := range interaction.Output {
				writeLine(o)
			}

			switch interaction.PostAction {
			case "clear":
				clearscreen()
			case "exit":
				os.Exit(0)
			default:
				break
			}
			gave_output = true
			break
		}
	}
	if !gave_output {
		writeLine("\n\nI'M SORRY PROFESSOR FALKEN,\nI DO NOT UNDERSTAND")
	}
}

func chat_logon() {
	prompt_input := InputPrompt("\nLOGON: ")

	if use_chat_lines {

		gave_output := false

		for _, v := range chat_lines_logon {

			if check_input(prompt_input, v.Input) {

				if strings.Contains("joshua", prompt_input) {
					got_joshua = true
				}

				switch v.PreAction {
				case "clear":
					clearscreen()
				default:
					break
				}

				for _, o := range v.Output {
					writeLine(o)
				}

				switch v.PostAction {
				case "exit":
					os.Exit(0)
				default:
					break
				}

				gave_output = true
				break
			}
		}

		if !gave_output {
			fmt.Printf("\n%s NOT AVAILABLE\n", strings.ToUpper(prompt_input))
		}
	} else {
		switch prompt_input {

		case "help logon":
			writeLine("\nHELP NOT AVAILABLE")

		case "joshua":
			got_joshua = true
			clearscreen()
			writeLine("\nGREETINGS PROFESSOR FALKEN.")

		case "help games":

			help_games_lines[0] = "\n"
			help_games_lines[1] = "`GAMES` REFERS TO MODELS, SIMULATIONS AND GAMES"
			help_games_lines[2] = "WHICH HAVE TACTICAL AND STRATEGIC APPLICATIONS.\n"
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
				writeLine(help_games_lines[i])
				if i == 19 {
					time.Sleep(time.Second * 1)
				} else {
					time.Sleep(time.Second / 2)
				}
			}

		default:
			fmt.Printf("\n%s NOT AVAILABLE\n", strings.ToUpper(prompt_input))
		}
	}
}

func main() {
	clearscreen()
	fmt.Println(chalk.CyanLight(), chalk.Bold())
	for {
		if !got_joshua {
			chat_logon()
		} else {
			chat_joshua()
		}
	}
}
