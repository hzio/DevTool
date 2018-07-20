package command

import (
	"encoding/json"
	"fmt"
	"global"
	"gopkg.in/urfave/cli.v2"
	"strconv"
	"time"
)

// ===========================================================
//
//       Operation commands declarations
//
// ===========================================================
var (
	TimeFormatCommand = &cli.Command{
		Name:        "timeformat",
		Aliases:     []string{"tf"},
		Usage:       "Format timestamp into date-time pattern",
		Description: "Format timestamp into date-time pattern",
		Category:    global.CategoryFormat,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "ms", Usage: "Generate current timestamp millisecond"},
		},
		Action: execTimeFormatCommand,
	}

	JSONFormatCommand = &cli.Command{
		Name:        "jsonformat",
		Aliases:     []string{"jf"},
		Usage:       "Fortmat json string with indent",
		Description: "Fortmat json string with indent",
		Category:    global.CategoryFormat,
		Action:      execJSONFormatCommand,
	}
)

// ===========================================================
//
//       Operation commands implementations
//
// ===========================================================

func execTimeFormatCommand(ctx *cli.Context) error {

	if ctx.Args().Get(0) == "" {
		fmt.Println("Timestamp not specified")
		return nil
	}

	ts, err := strconv.ParseInt(ctx.Args().Get(0), 10, 64)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if ctx.Bool("ms") {
		ts = ts / 1000
	}

	fmt.Println(time.Unix(ts, 0).Format("2006-01-02 15:04:05"))
	return nil
}

func execJSONFormatCommand(ctx *cli.Context) error {

	if ctx.Args().Get(0) == "" {
		fmt.Println("JSON string not specified")
		return nil
	}

	source := string(ctx.Args().Get(0))
	var dat map[string]interface{}

	// parse into a map
	if err := json.Unmarshal([]byte(source), &dat); err != nil {
		fmt.Println(err)
		return nil
	}

	// marshal
	result, err := json.MarshalIndent(dat, "", "\t")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(string(result))
	return nil
}
