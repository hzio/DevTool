package launcher

import (
	cmds "command"
	"fmt"
	"global"
	"gopkg.in/urfave/cli.v2"
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
		Name:                  "DevTool",
		Version:               "1.0",
		Compiled:              time.Now(),
		Usage:                 "An all-round developer toolset.",
		Description:           "The toolset includes encryption, encoding, formatting, generation, searching, etc...",
		EnableShellCompletion: true,
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
		fmt.Fprintln(os.Stderr, err)
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
		// encrypt operation commands
		// ===============================
		cmds.Base64EncryptionCommand,
		cmds.Md5EncryptionCommand,

		// ===============================
		// Transcode operation commands
		// ===============================
		cmds.UrlEncodeTranscodeCommand,
		cmds.UrlDecodeTranscodeCommand,

		// ===============================
		// Generate operation commands
		// ===============================
		cmds.TimeGenerateCommand,
		cmds.DateGenerateCommand,
		cmds.TimestampGenerateCommand,
		cmds.PasswordGenerateCommand,
		cmds.UuidGenerateCommand,

		// ===============================
		// Format operation commands
		// ===============================
		cmds.TimeFormatCommand,
		cmds.JSONFormatCommand,

		// ===============================
		// Search operation commands
		// ===============================
		cmds.BaiduSearchCommand,
		cmds.GoogleSearchCommand,
		cmds.BaiduMapSearchCommand,
		cmds.GoogleMapSearchCommand,
		cmds.WikipediaSearchCommand,

		// ===============================
		// Fetch operation commands
		// ===============================
		cmds.IpFetchCommand,

		// ===============================
		// Server operation commands
		// ===============================
		cmds.StaticServerServeCommand,
	}

}
