package conversation

type InputOutput struct {
	Input      string   `json:"input"`
	Output     []string `json:"output"`
	PreAction  string   `json:"pre_action"`
	PostAction string   `json:"post_action"`
	UnLocks    string   `json:"unlocks"`
	NeedGame   bool     `json:"need_game"`
	NeedJoshua bool     `json:"need_joshua"`
}

var (
	ChatlinesLogon = []InputOutput{
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
			NeedGame:   true,
			PostAction: "exit",
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

	ChatlinesJoshua = []InputOutput{
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
			Input:      "Later. Let's play Global Thermonuclear War.",
			Output:     []string{"\n\nFINE"},
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
				"\n",
				"HOW ABOUT A NICE GAME OF CHESS?",
			},
			PostAction: "exit",
		},
	}
)
