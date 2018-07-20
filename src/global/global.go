package global

import "gopkg.in/urfave/cli.v2"

// ===========================================================
//
//       Command category declarations
//
// ===========================================================
const (
	CategoryEncrypt  = "ENCRYPT"
	CategoryFetch    = "FETCH"
	CategoryFormat   = "FORMAT"
	CategoryGenerate = "GENERATE"
	CategorySearch   = "SEARCH"
	CategoryServer   = "SERVER"
	CategoryEncode   = "ENCODE"


	Source = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_+<>?[]"
)

// ===========================================================
//
//       Cross-platform commands definitions
//
// ===========================================================
var (
	Commands = map[string]map[string]map[string]map[string]string{
		"windows": {
			"actions": {
				"openBrowser": {
					"cmd":  "cmd",
					"args": "/c start ",
				},
			},
		},
		"darwin": {
			"actions": {
				"openBrowser": {
					"cmd":  "open",
					"args": "",
				},
			},
		},
		"linux": {
			"actions": {
				"openBrowser": {
					"cmd":  "xdg-open",
					"args": "",
				},
			},
		},
	}
)

// ===========================================================
//
//       Global flag declarations
//
// ===========================================================
var (
	ReverseFlag = &cli.BoolFlag{
		Name:    "reverse",
		Aliases: []string{"r"},
		Usage:   "Reverse current operation",
	}
)
