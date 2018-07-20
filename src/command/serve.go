package command

import (
	"fmt"
	"global"
	"gopkg.in/urfave/cli.v2"
	"net/http"
	"time"
	"tool"
)

// ===========================================================
//
//       Operation commands declarations
//
// ===========================================================
var (
	StaticServerServeCommand = &cli.Command{
		Name:        "serve",
		Aliases:     []string{"s"},
		Usage:       "Start a static file server at port 9090",
		Description: "Start a static file server at port 9090",
		Category:    global.CategoryServer,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "port", Aliases: []string{"p"}, Usage: "Listen the specified port"},
			&cli.BoolFlag{Name: "silent", Aliases: []string{"s"}, Usage: "Do NOT open the browser after server started"},
		},
		Action: execStaticServerServeCommand,
	}
)

// ===========================================================
//
//       Operation commands implementations
//
// ===========================================================
func execStaticServerServeCommand(ctx *cli.Context) error {

	dir := "."
	if ctx.Args().Get(0) != "" {
		dir = ctx.Args().Get(0)
	}

	port := "9090"
	if ctx.String("port") != "" {
		port = ctx.String("port")
	}

	// if not a valid directory
	if !tool.IsDir(dir) {
		fmt.Printf("Directory %s not exists", dir)
		return nil
	}

	// add handle function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	})

	if !ctx.Bool("silent") {
		// open the uri in browser
		go openServerUri(port)
	}

	fmt.Printf("Serving HTTP on 0.0.0.0 port %s ...\n", port)
	// start listening
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}

// open the server uri via browser
func openServerUri(port string) {

	time.Sleep(time.Millisecond * time.Duration(300))
	url := "http://127.0.0.1:" + port
	tool.OpenURI(url)
}
