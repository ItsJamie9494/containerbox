package main

import (
	"os"
	
	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
	"github.com/ItsJamie9494/containerbox/cmd"
)

var errorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FF0000"))

func main() {
    app := &cli.App{
    	EnableBashCompletion: true,
    	Name:		      "containerbox",
    	Version:              "v1.0",
    	Usage:		      "Containerise your system",
    	UsageText:	      "containerbox command [options] [arguments...]",
    	Flags: []cli.Flag{
    	    &cli.BoolFlag{
    	    	Name:    "verbose",
    	    	Value:   false,
    	    	Aliases: []string{"V"},
    	    },
    	    &cli.BoolFlag{
    	    	Name:    "root",
    	    	Value:   false,
    	    	Usage:   "Execute all lower-level functions with Root permissions. Use caution",
    	    },
    	},
    	Commands: []*cli.Command{
    	    {
    	    	Name:      "create",
    	    	Usage:     "Create a new container",
    	    	UsageText: "containerbox create <idk yet>",
    	    	Aliases:   []string{"new"},
    	    	Action:    cmd.Create,
    	    },
    	    {
    	    	Name:      "enter",
    	    	Usage:     "Enter an existing container",
    	    	UsageText: "containerbox enter <idk yet>",
    	    	Aliases:   []string{"start", "run"},
    	    	Action:    cmd.Enter,
    	    },
    	    {
    	    	Name:      "list",
    	    	Usage:     "List all containers",
    	    	UsageText: "containerbox list",
    	    	Action:    cmd.List,
    	    },
    	    {
    	    	Name:      "stop",
    	    	Usage:     "Stop a started container",
    	    	UsageText: "containerbox stop <idk yet>",
    	    	Aliases:   []string{"kill", "quit"},
    	    	Action:    cmd.Stop,
    	    },
    	    {
    	    	Name:      "remove",
    	    	Usage:     "Remove an existing container",
    	    	UsageText: "containerbox remove <idk yet>",
    	    	Aliases:   []string{"delete", "rm",},
    	    	Action:    cmd.Remove,
    	    },
    	},
    }
    
    if err := app.Run(os.Args); err != nil {
    	println(errorStyle.Render("Error: ") + err.Error())
    	os.Exit(1)
    }
}
