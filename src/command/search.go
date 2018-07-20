package command

import (
	"gopkg.in/urfave/cli.v2"

	"global"
	"tool"
)

// ===========================================================
//
//       Operation commands declarations
//
// ===========================================================
var (
	BaiduSearchCommand = &cli.Command{
		Name:        "baidu",
		Aliases:     []string{"b"},
		Usage:       "Search the specified keyword with Baidu search engine",
		Description: "Search the specified keyword with Baidu search engine",
		Category:    global.CategorySearch,
		Action:      execBaiduSearchCommand,
	}

	GoogleSearchCommand = &cli.Command{
		Name:        "google",
		Aliases:     []string{"g"},
		Usage:       "Search the specified keyword with Google search engine",
		Description: "Search the specified keyword with Google search engine",
		Category:    global.CategorySearch,
		Action:      execGoogleSearchCommand,
	}

	BaiduMapSearchCommand = &cli.Command{
		Name:        "baidumap",
		Aliases:     []string{"map"},
		Usage:       "Locate the specified place with Baidu map search engine",
		Description: "Locate the specified place with Baidu map search engine",
		Category:    global.CategorySearch,
		Action:      execBaiduMapSearchCommand,
	}

	GoogleMapSearchCommand = &cli.Command{
		Name:        "googlemap",
		Aliases:     []string{"gmap"},
		Usage:       "Locate the specified place with Google map search engine",
		Description: "Locate the specified place with Google map search engine",
		Category:    global.CategorySearch,
		Action:      execGoogleMapSearchCommand,
	}

	WikipediaSearchCommand = &cli.Command{
		Name:        "wiki",
		Aliases:     []string{"wk"},
		Usage:       "Define the specified word with Wikipedia",
		Description: "Define the specified word with Wikipedia",
		Category:    global.CategorySearch,
		Action:      execWikipediaSearchCommand,
	}
)

// ===========================================================
//
//       Operation commands implementations
//
// ===========================================================
func execBaiduSearchCommand(ctx *cli.Context) error {
	url := "https://www.baidu.com/s?wd="
	tool.OpenURI(url + ctx.Args().Get(0))
	return nil
}

func execGoogleSearchCommand(ctx *cli.Context) error {
	url := "https://www.google.com/search?q="
	tool.OpenURI(url + ctx.Args().Get(0))
	return nil
}

func execBaiduMapSearchCommand(ctx *cli.Context) error {
	url := "http://map.baidu.com/?q="
	tool.OpenURI(url + ctx.Args().Get(0))
	return nil
}

func execGoogleMapSearchCommand(ctx *cli.Context) error {
	url := "https://www.google.com/maps/search/"
	tool.OpenURI(url + ctx.Args().Get(0))
	return nil
}

func execWikipediaSearchCommand(ctx *cli.Context) error {
	url := "https://en.wikipedia.org/wiki/"
	tool.OpenURI(url + ctx.Args().Get(0))
	return nil
}
