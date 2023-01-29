package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
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
	UnLocks    string
	NeedGame   bool
	NeedJoshua bool
}

var (
	baudrate         int = 25
	variability      int = 5
	got_joshua           = false
	got_games            = false
	start            time.Time
	prompt           = "\nLOGON: "
	chat_lines_logon = []InputOutput{
		{
			Input:  "Help Logon",
			Output: []string{"\nHELP NOT AVAILABLE"},
		},
		{
			Input: "Help games",
			Output: []string{
				"\n",
				"`GAMES` REFERS TO MODELS, SIMULATIONS AND GAMES",
				"WHICH HAVE TACTICAL AND STRATEGIC APPLICATIONS.\n",
			},
			PostAction: "noprompt",
			UnLocks:    "games",
		},
		{
			Input: "List games",
			Output: []string{
				"\n", "FALKEN'S MAZE",
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
			NeedGame: true,
		},
		{
			Input: "Armageddon",
			Output: []string{
				"IDENTIFICATION NOT RECOGNIZED BY SYSTEM",
				"--CONNECTION-TERMINATED__"},
			PostAction: "exit",
		},
		{
			Input: "joshua",
			Output: []string{
				"\n",
				"GREETINGS PROFESSOR FALKEN.",
			},
			UnLocks:   "joshua",
			PreAction: "clear",
		},
		{
			Input: "7KQ201 McKittrick",
			Output: []string{
				"** IDENTIFICATION NOT RECOGNIZED **",
				"\n",
				"** ACCESS DENIED **",
			},
			PostAction: "exit",
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
			Input:  "Later. Let's play Global Thermonuclear War.",
			Output: []string{"\n\nFINE"}, /*	`
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
											'~'`,*/

			PostAction: "hack",
		},
		{
			Input:      "bye",
			Output:     []string{"\n\nHOPE TO SEE YOU SOON", "BYE"},
			PostAction: "exit",
		},
		{
			Input:      "wargames",
			Output:     []string{"\n\nNICE MOVIE", "SOON IN A THEATER NEAR YOU"},
			PostAction: "exit",
		},
		{
			Input: "CPE1704TKS",
			Output: []string{
				"A STRANGE GAME",
				"THE ONLY WINNING MOVE IS",
				"NOT TO PLAY.",
			},
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
	baudprint.BaudPrint(input, int64(baudrate), int(variability), false, false)
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

func randomString(charset string, n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func hackNuclearCodes() {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(time.Now().Unix())

	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numset := "0123456789"
	set1 := randomString(charset, 3)
	set2 := randomString(numset, 4)
	set3 := randomString(charset, 3)
	code := set1 + set2 + set3
	writeLine(code)
	//fmt.Println(code)
	baudrate = 300
	if code == "JPE1704TKS" || code == "CPE1704TKS" {
		writeLine("FOUND " + code)
		fmt.Printf("%s took %v\n", code, time.Since(start))
		os.Exit(0)
	}

	//baudprint.BaudPrint(code, 1000, 1, false, false)
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
			case "hack":
				clearscreen()
				fmt.Print("\033[s")
				start = time.Now()
				for {
					fmt.Print("\033[u\033[K")
					hackNuclearCodes()
				}
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
	prompt_input := InputPrompt(prompt)

	gave_output := false

	for _, v := range chat_lines_logon {

		if check_input(prompt_input, v.Input) {
			if v.NeedGame == got_games {

				switch v.UnLocks {
				case "games":
					got_games = true
				case "joshua":
					got_joshua = true
				default:
					break
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
				case "noprompt":
					prompt = "\n"
				default:
					break
				}

				gave_output = true
				break
			}
		}
	}

	if !gave_output {
		fmt.Printf("\n%s NOT AVAILABLE\n", strings.ToUpper(prompt_input))

	}
}

func main() {

	flag.IntVar(&baudrate, "bd", 20, "Specify baud rate.")
	flag.IntVar(&variability, "var", 3, "Specify variability.")
	//flag.StringVar(&pass, "p", "password", "Specify pass. Default is password")

	flag.Parse() // after declaring flags we need to call it

	if baudrate < variability {
		fmt.Println("baudrate should be bigger then [variability]", variability)
		os.Exit(1)
	}

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
