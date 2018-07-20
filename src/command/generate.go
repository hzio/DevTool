package command

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gopkg.in/urfave/cli.v2"

	"global"
	"strconv"
	"time"
	"tool"
)

// ===========================================================
//
//       Operation commands declarations
//
// ===========================================================
var (
	TimeGenerateCommand = &cli.Command{
		Name:        "time",
		Aliases:     []string{"t", "now"},
		Usage:       "Generate current date-time string",
		Description: "Generate current date-time string",
		Category:    global.CategoryGenerate,
		Action:      execTimeGenerateCommand,
	}

	DateGenerateCommand = &cli.Command{
		Name:        "date",
		Aliases:     []string{"d"},
		Usage:       "Generate current date string",
		Description: "Generate current date string",
		Category:    global.CategoryGenerate,
		Action:      execDateGenerateCommand,
	}

	TimestampGenerateCommand = &cli.Command{
		Name:        "timestamp",
		Aliases:     []string{"ts"},
		Usage:       "Generate current timestamp",
		Description: "Generate current timestamp",
		Category:    global.CategoryGenerate,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "ms", Usage: "Generate current timestamp millisecond"},
		},
		Action: execTimestampGenerateCommand,
	}

	PasswordGenerateCommand = &cli.Command{
		Name:        "passwd",
		Usage:       "Generate a random password of 16 bits",
		Description: "Generate a random password of 16 bits",
		Category:    global.CategoryGenerate,
		Action: execPasswordGenerateCommand,
	}

	UuidGenerateCommand = &cli.Command{
		Name:        "uuid",
		Usage:       "Generate a UUID",
		Description: "Generate a UUID",
		Category:    global.CategoryGenerate,
		Action:      execUuidGenerateCommand,
	}
)

// ===========================================================
//
//       Operation commands implementations
//
// ===========================================================

func execTimeGenerateCommand(ctx *cli.Context) error {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

func execDateGenerateCommand(ctx *cli.Context) error {
	fmt.Println(time.Now().Format("2006-01-02"))
	return nil
}

func execTimestampGenerateCommand(ctx *cli.Context) error {

	ts := time.Now().Unix()
	// millisecond
	if ctx.Bool("ms") {
		ts = ts * 1000
	}

	fmt.Println(int64(ts))
	return nil
}

func execPasswordGenerateCommand(ctx *cli.Context) error {

	finalLen := 16
	if ctx.Args().Get(0) != "" {

		targetLen, err := strconv.Atoi(ctx.Args().Get(0))
		if err != nil || targetLen <= 0 {
			fmt.Println("Invalid length")
			return nil
		} else {
			finalLen = targetLen
		}
	}
	fmt.Println(tool.GenRandomString(finalLen))
	return nil
}


func execUuidGenerateCommand(ctx *cli.Context) error {

	data := []byte(ctx.Args().Get(0))
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	fmt.Println(hex.EncodeToString(cipherStr))
	return nil
}
