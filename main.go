package main

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"strings"
)

type DemoPlugin struct{}

func (p *DemoPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	apiVersion, _ := cliConnection.ApiVersion()
	fmt.Println("cf API Version: " + apiVersion)
	cliVersionOutput, _ := cliConnection.CliCommandWithoutTerminalOutput("version")
	cliFullVersion := strings.Split(cliVersionOutput[0], " ")[2]
	cliVersion := strings.Split(cliFullVersion, "+")[0]
	fmt.Println("cf CLI Version: " + cliVersion)
	return
}

func (p *DemoPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "DemoPlugin",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "demo",
				HelpText: "GA's v8!",
				UsageDetails: plugin.Usage{
					Usage: "cf demo",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(DemoPlugin))
}
