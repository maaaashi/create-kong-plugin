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
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func selectLanguage(cmd *cobra.Command) (string, error) {
	setLanguage, err := cmd.PersistentFlags().GetString("language")

	if err != nil {
		fmt.Println("Error getting language flag:", err)
		return "", err
	}

	var language string

	if setLanguage != "" {
		language = setLanguage
	} else {
		languages := []string{"Lua", "Go", "x JavaScript", "x Python"}

		prompt := promptui.Select{
			Label: "Select Language",
			Items: languages,
		}

		_, lang, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return "", err
		}

		language = lang
	}

	if language != "" && language != "Lua" && language != "Go" {
		return "", fmt.Errorf("invalid language. Supported languages are: Lua, Go")
	}

	return language, nil
}

func setPluginName(reader *bufio.Reader, args []string) (string, error) {
	pluginName := ""

	if len(args) == 0 {
		fmt.Print("Enter the plugin name: ")

		name, _ := reader.ReadString('\n')
		pluginName = strings.Split(name, "\n")[0]
	} else {
		pluginName = args[0]
	}

	if pluginName == "" {
		return "", fmt.Errorf("plugin name cannot be empty")
	}

	return pluginName, nil
}

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

		pluginName, err := setPluginName(reader, args)

		if err != nil {
			fmt.Println("Error setting plugin name:", err)
			return
		}

		createDirectoryFlag := false

		if pluginName == "." {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Println("Error getting current directory:", err)
				return
			}
			pluginName = filepath.Base(wd)
		} else {
			createDirectoryFlag = true
		}

		language, err := selectLanguage(cmd)

		if err != nil {
			fmt.Println("Error selecting language:", err)
			return
		}

		p := promptui.Prompt{
			Label:     "Name: " + pluginName + ", Language: " + language + ". Continue?",
			IsConfirm: true,
		}

		_, err = p.Run()

		if err != nil {
			return
		}

		handler.CreatePluginTemplate(pluginName, language, createDirectoryFlag)
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringP("language", "l", "", "Language to use for the plugin (Lua, Go)")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
