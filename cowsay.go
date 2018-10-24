package main

import (
	"fmt"
	"strings"

	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cli/plugin/models"
	cowsayer "github.com/clijockey/cowsay/cowsayer"
	"gopkg.in/kyokomi/emoji.v1"
)

// Struct implementing the interface defined by the core CLI. It can
// be found at  "code.cloudfoundry.org/cli/plugin/plugin.go"
type Cowsay struct{}

// Run must be implemented by any plugin because it is part of the
// plugin interface defined by the core CLI.
//
// Run(....) is the entry point when the core CLI is invoking a command defined
// by a plugin. The first parameter, plugin.CliConnection, is a struct that can
// be used to invoke cli commands. The second paramter, args, is a slice of
// strings. args[0] will be the name of the command, and will be followed by
// any additional arguments a cli user typed in.
//
// Any error handling should be handled with the plugin itself (this means printing
// user facing errors). The CLI will exit 0 if the plugin exits 0 and will exit
// 1 should the plugin exits nonzero.
func (c *Cowsay) Run(cliConnection plugin.CliConnection, args []string) {

	if args[0] == "cowsay" {
		if len(args) > 1 {
			var err error
			if args[1] == "apps" {
				// Work out the running apps in the space and also thier status
				appsListing, err := cliConnection.GetApps()

				if err != nil {
					//c.ui.Failed(err.Error())
				}

				var appsStatus []string

				for _, app := range appsListing {
					var state string
					if app.State == "started" {
						state = emoji.Sprint("is all good :white_check_mark:")
					} else if app.State == "stopped" {
						state = emoji.Sprint("has stopped man!! :red_circle:")
					} else {
						state = app.State + emoji.Sprint(" :confused:")
					}
					data := "So `" + app.Name + "` " + state
					appsStatus = append(appsStatus, data)
				}

				fmt.Println(cowsayer.Cow(strings.Join(appsStatus, "\n")))

			} else if args[1] == "space" {
				// Work out the space and org that you are in so it can be retruned to the CLI
				var space plugin_models.Space
				var org plugin_models.Organization

				if space, err = cliConnection.GetCurrentSpace(); err != nil {
				}
				if org, err = cliConnection.GetCurrentOrg(); err != nil {
				}
				fmt.Println(cowsayer.Cow("Space: " + space.Name + " in the organisation: " + org.Name))
			} else {
				// Default answer so that it just prints the message you passed to the CLI
				fmt.Println(cowsayer.Simplesay(args))
			}
		} else {
			fmt.Println(cowsayer.Simplesay(args))
		}
	}
}

// GetMetadata must be implemented as part of the plugin interface
// defined by the core CLI.
//
// GetMetadata() returns a PluginMetadata struct. The first field, Name,
// determines the name of the plugin which should generally be without spaces.
// If there are spaces in the name a user will need to properly quote the name
// during uninstall otherwise the name will be treated as seperate arguments.
// The second value is a slice of Command structs. Our slice only contains one
// Command Struct, but could contain any number of them. The first field Name
// defines the command `cf basic-plugin-command` once installed into the CLI. The
// second field, HelpText, is used by the core CLI to display help information
// to the user in the core commands `cf help`, `cf`, or `cf -h`.
func (c *Cowsay) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "cowsay",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 8,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "cowsay",
				HelpText: "Plain old cowsay example for a little bit of fun!",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "cf cowsay\n	cf cowsay <some text>\n cf cowsay space\n cf cowsay apps",
				},
			},
		},
	}
}

// Unlike most Go programs, the `Main()` function will not be used to run all of the
// commands provided in your plugin. Main will be used to initialize the plugin
// process, as well as any dependencies you might require for your
// plugin.
func main() {
	// Any initialization for your plugin can be handled here
	//
	// Note: to run the plugin.Start method, we pass in a pointer to the struct
	// implementing the interface defined at "code.cloudfoundry.org/cli/plugin/plugin.go"
	//
	// Note: The plugin's main() method is invoked at install time to collect
	// metadata. The plugin will exit 0 and the Run([]string) method will not be
	// invoked.
	plugin.Start(new(Cowsay))
	// Plugin code should be written in the Run([]string) method,
	// ensuring the plugin environment is bootstrapped.
}
