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
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func writeTemplateToFile(filePath, tmpl, pluginName string) {
	t, err := template.New("plugin").Parse(tmpl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()

	data := struct {
		PluginName string
	}{
		PluginName: pluginName,
	}

	if err := t.Execute(f, data); err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}

func createPluginTemplate(pluginName string) {
	pluginRootDir := filepath.Join(".", pluginName)
	srcDir := filepath.Join(pluginRootDir, "/src")

	if err := os.MkdirAll(srcDir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// rockspec テンプレート
	rockspecTemplate := `
package = "kong-plugin-{{.PluginName}}"
version = "0.1.0-1"

source = {
	url = ""
}

build = {
    type = "builtin",
    modules = {
        ["kong.plugins.{{.PluginName}}.handler"] = "src/handler.lua",
        ["kong.plugins.{{.PluginName}}.schema"] = "src/schema.lua"
    }
}
	`

	// handler.lua テンプレート
	handlerTemplate := `
local {{.PluginName}} = {
	VERSION = "1.0.0",
	PRIORITY = 10,
}

function {{.PluginName}}:access(conf)
	-- plugin logic here
	kong.log("This is an example plugin handler, conf: ", conf)
end

return {{.PluginName}}
`
	// schema.lua テンプレート
	schemaTemplate := `
return {
	name = "{{.PluginName}}",
	fields = {},
}
`

	// テンプレートをファイルに書き込む
	writeTemplateToFile(filepath.Join(pluginRootDir, "kong-plugin-"+pluginName+"-0.1.0-1.rockspec"), rockspecTemplate, pluginName)
	writeTemplateToFile(filepath.Join(srcDir, "handler.lua"), handlerTemplate, pluginName)
	writeTemplateToFile(filepath.Join(srcDir, "schema.lua"), schemaTemplate, pluginName)

	fmt.Printf("Kong plugin template for '%s' created successfully!\n", pluginName)
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

		fmt.Println("Creating a new Kong plugin template: ", pluginName)
		createPluginTemplate(pluginName)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
