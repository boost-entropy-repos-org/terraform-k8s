package main

import (
	"os"

	cmdVersion "github.com/hashicorp/terraform-k8s/subcommand/version"
	cmdSyncWorkspace "github.com/hashicorp/terraform-k8s/subcommand/workspace"
	"github.com/hashicorp/terraform-k8s/version"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all available terraform-k8s commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	Commands = map[string]cli.CommandFactory{
		"sync-workspace": func() (cli.Command, error) {
			return &cmdSyncWorkspace.Command{UI: ui}, nil
		},
		"version": func() (cli.Command, error) {
			return &cmdVersion.Command{UI: ui, Version: version.GetHumanVersion()}, nil
		},
	}
}

func helpFunc() cli.HelpFunc {
	// This should be updated for any commands we want to hide for any reason.
	// Hidden commands can still be executed if you know the command, but
	// aren't shown in any help output. We use this for prerelease functionality
	// or advanced features.
	hidden := map[string]struct{}{
		"inject-connect": struct{}{},
	}

	var include []string
	for k := range Commands {
		if _, ok := hidden[k]; !ok {
			include = append(include, k)
		}
	}

	return cli.FilteredHelpFunc(include, cli.BasicHelpFunc("terraform-k8s"))
}