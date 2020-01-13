package launcher

import (
	"DevTool/command"
	"DevTool/global"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

// ===========================================================
//
//        App declaration
//
// ===========================================================
var (
	app = &cli.App{
		Name:                 "DevTool",
		Version:              "1.0",
		Compiled:             time.Now(),
		Usage:                "An all-round developer toolset.",
		Description:          "The toolset includes encryption, encoding, formatting, generation, searching, etc...",
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Hunter Zhao",
				Email: "zhaohevip@gmail.com / GitHub: https://github.com/MrHunterZhao",
			},
		},
	}
)

// ===========================================================
//
//       App Launcher
//
// ===========================================================
func Launch() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// ===========================================================
//
//       App initializer
//
// ===========================================================
func init() {

	app.Flags = []cli.Flag{
		global.ReverseFlag,
	}

	app.Commands = []*cli.Command{
		// ===============================
		// encrypt operation command
		// ===============================
		command.Base64EncryptionCommand,
		command.Md5EncryptionCommand,

		// ===============================
		// Transcode operation command
		// ===============================
		command.UrlEncodeTranscodeCommand,
		command.UrlDecodeTranscodeCommand,

		// ===============================
		// Generate operation command
		// ===============================
		command.TimeGenerateCommand,
		command.DateGenerateCommand,
		command.TimestampGenerateCommand,
		command.PasswordGenerateCommand,
		command.UuidGenerateCommand,

		// ===============================
		// Format operation command
		// ===============================
		command.TimeFormatCommand,
		command.JSONFormatCommand,

		// ===============================
		// Search operation command
		// ===============================
		command.BaiduSearchCommand,
		command.GoogleSearchCommand,
		command.BaiduMapSearchCommand,
		command.GoogleMapSearchCommand,
		command.WikipediaSearchCommand,

		// ===============================
		// Fetch operation command
		// ===============================
		command.IpFetchCommand,

		// ===============================
		// Server operation command
		// ===============================
		command.StaticServerServeCommand,
	}

}
