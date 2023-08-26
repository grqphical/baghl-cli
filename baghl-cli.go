package main

import (
	"log"
	"os"

	"github.com/teris-io/cli"
)

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createCmd := cli.NewCommand("create", "Creates new projects").WithArg(cli.NewArg("name", "Name of new project")).
		WithAction(func(args []string, options map[string]string) int {
			name := args[0]
			SetupProjectFiles(name)
			return 0
		})

	getAssetCmd := cli.NewCommand("getasset", "Gets a web asset for the project in the current directory").
		WithArg(cli.NewArg("url", "URL of file")).
		WithArg(cli.NewArg("filename", "Name of downloaded file")).
		WithAction(func(args []string, options map[string]string) int {
			url, filename := args[0], args[1]
			DownloadAssetFile(url, filename)
			return 0
		})

	app := cli.New("Command line interface for creating BAGHL projects").WithCommand(createCmd).WithCommand(getAssetCmd)

	os.Exit(app.Run(os.Args, os.Stdout))
}
