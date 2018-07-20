package command

import (
	"fmt"
	"global"
	"gopkg.in/urfave/cli.v2"
	"net/url"
)

// ===========================================================
//
//       Operation commands declarations
//
// ===========================================================
var (
	UrlEncodeTranscodeCommand = &cli.Command{
		Name:        "urlencode",
		Usage:       "Urlencode the specified string",
		Description: "Urlencode the specified string",
		Category:    global.CategoryEncode,
		Action:      execUrlEncodeTranscodeCommand,
	}

	UrlDecodeTranscodeCommand = &cli.Command{
		Name:        "urldecode",
		Usage:       "Urldecode the specified string",
		Description: "Urldecode the specified string",
		Category:    global.CategoryEncode,
		Action:      execUrlDecodeTranscodeCommand,
	}
)

// ===========================================================
//
//       Operation commands implementations
//
// ===========================================================
func execUrlEncodeTranscodeCommand(ctx *cli.Context) error {

	if ctx.Args().Get(0) == "" {
		return nil
	}

	if ctx.Bool("reverse") {
		execUrlDecodeTranscodeCommand(ctx)
		return nil
	}

	fmt.Println(url.QueryEscape(ctx.Args().Get(0)))
	return nil
}

func execUrlDecodeTranscodeCommand(ctx *cli.Context) error {

	if ctx.Args().Get(0) == "" {
		return nil
	}

	result, err := url.QueryUnescape(ctx.Args().Get(0))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return nil
}
