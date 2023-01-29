package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/golang-demos/chalk"
	baudprint "ruijs.fr/protovision/BaudPrint"
	conversation "ruijs.fr/protovision/Conversation"
)

var (
	baudrate    int = 25
	variability int = 5
	load        bool
	got_joshua  = false
	got_games   = false
	start       time.Time
	prompt      = "\nLOGON: "

	filename_joshua                              = "joshua.json"
	filename_logon                               = "logon.json"
	chat_lines_logon  []conversation.InputOutput = conversation.ChatlinesLogon
	chat_lines_joshua []conversation.InputOutput = conversation.ChatlinesJoshua
)

func inputPrompt(label string) string {
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
	log.Printf("writeLine '%s'\n", input)
}

func check_input(prompt_input string, known_input string) bool {
	ret := false
	if len(prompt_input) > 2 {
		prompt_input_lower := strings.ToLower(prompt_input)
		known_input_lower := strings.ToLower(known_input)
		ret = prompt_input_lower == known_input_lower || strings.Contains(known_input_lower, prompt_input_lower) || strings.Contains(prompt_input_lower, known_input_lower)
	}
	log.Printf("'%t' comparing user input vs known: '%s' with known input '%s'", ret, prompt_input, known_input)
	return ret
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
		log.Printf("%s took %v\n", code, time.Since(start))
		os.Exit(0)
	}

	//baudprint.BaudPrint(code, 1000, 1, false, false)
}

func chat_joshua() {
	prompt_input := inputPrompt("\n")

	gave_output := false
	for _, interaction := range chat_lines_joshua {
		log.Printf(interaction.Input)
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
	prompt_input := inputPrompt(prompt)

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

func loadConversations() {
	contentLogon, err := os.ReadFile(filename_logon)
	if err != nil {
		log.Printf("Error when opening file '%s'\n", filename_logon)
		chat_lines_logon = conversation.ChatlinesLogon
		saveLogon()
	} else {
		log.Printf("Loaded LOGON conversation from file '%s'\n", filename_logon)

		err = json.Unmarshal(contentLogon, &chat_lines_logon)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
	}

	contentJoshua, err := os.ReadFile(filename_joshua)
	if err != nil {
		log.Printf("Error when opening file '%s'\n", filename_logon)
		chat_lines_joshua = conversation.ChatlinesJoshua
		saveJoshua()
	} else {
		log.Printf("Loaded JOSHUA conversation from file '%s'\n", filename_logon)
		err = json.Unmarshal(contentJoshua, &chat_lines_joshua)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}
	}

}
func saveLogon() {
	file, _ := json.MarshalIndent(conversation.ChatlinesLogon, "", " ")
	_ = os.WriteFile("logon.json", file, 0644)
}

func saveJoshua() {
	file, _ := json.MarshalIndent(conversation.ChatlinesJoshua, "", " ")
	_ = os.WriteFile("joshua.json", file, 0644)
}
func setup_logging() {
	f, err := os.OpenFile("protovision.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//defer to close when you're done with it, not because you think it's idiomatic!
	//defer f.Close()

	//set output of logs to f
	log.SetOutput(f)

	//test case
	log.Println("ProtoVision, I have you know!")
}
func main() {

	setup_logging()

	flag.IntVar(&baudrate, "bd", 300, "Specify baud rate.")
	flag.IntVar(&variability, "var", 30, "Specify variability.")
	flag.BoolVar(&load, "load", false, "Load conversation from files")
	flag.Parse() // after declaring flags we need to call it

	if load {
		loadConversations()
	}

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
