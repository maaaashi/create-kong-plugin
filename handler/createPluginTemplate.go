package handler

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
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

func CreatePluginTemplate(pluginName string, mkdir bool) {
	pluginRootDir := "."

	if mkdir {
		pluginRootDir = filepath.Join(".", pluginName)
	}
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
