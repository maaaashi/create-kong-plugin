/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"

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
package = "{{.PluginName}}"
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
end

return {{.PluginName}}
`
	// schema.lua テンプレート
	schemaTemplate := `
return {
	name = "{{.PluginName}}",
	fields = {
			{ config = {
					type = "record",
					fields = {
							{ my_option = { type = "string", required = true } },
					},
			}},
	},
}
`

	// テンプレートをファイルに書き込む
	writeTemplateToFile(filepath.Join(pluginRootDir, "kong-plugin-"+pluginName+"-0.1.0-1.rockspec"), rockspecTemplate, pluginName)
	writeTemplateToFile(filepath.Join(srcDir, "handler.lua"), handlerTemplate, pluginName)
	writeTemplateToFile(filepath.Join(srcDir, "schema.lua"), schemaTemplate, pluginName)

	fmt.Printf("Kong plugin template for '%s' created successfully!\n", pluginName)
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [plugin-name]",
	Short: "Create a new Kong plugin template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pluginName := args[0]
		fmt.Println("Creating a new Kong plugin template: ", pluginName)
		createPluginTemplate(pluginName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
