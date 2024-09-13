/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/maaaashi/create-kong-plugin/handler"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "create-kong-plugin [plugin-name]",
	Short: "Create a new Kong plugin template",
	Long: `CreateKongPlugin is a CLI tool designed to generate template files for creating Kong plugins. With a simple command, it creates the necessary Lua files (handler.lua and schema.lua) and directory structure, making it easier for developers to kickstart their Kong plugin development.

Key Features:
・Quickly generate Kong plugin templates
・Customize plugin names and structure
・Easy to use and lightweight
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		pluginName := ""

		if len(args) == 0 {
			fmt.Print("Enter the plugin name: ")

			name, _ := reader.ReadString('\n')
			pluginName = strings.Split(name, "\n")[0]
		} else {
			pluginName = args[0]
		}

		if pluginName == "" {
			fmt.Println("Plugin name cannot be empty")
			return
		}

		if pluginName == "." {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				return
			}
			pluginName = filepath.Base(wd)
			handler.CreatePluginTemplate(pluginName, false)
		} else {
			handler.CreatePluginTemplate(pluginName, true)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
