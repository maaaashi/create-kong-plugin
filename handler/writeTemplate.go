package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
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

func WriteLuaTemplate(srcDir string, pluginRootDir string, pluginName string) {
	// rockspec テンプレート
	rockspecTemplate := `package = "kong-plugin-{{.PluginName}}"
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
	handlerTemplate := `local plugin = {
	VERSION = "1.0.0",
	PRIORITY = 10,
}

function plugin:access(conf)
	-- plugin logic here
	kong.log("This is an example lua plugin handler, conf: ", conf)
end

return plugin
`
	// schema.lua テンプレート
	schemaTemplate := `return {
	name = "{{.PluginName}}",
	fields = {},
}
`

	// テンプレートをファイルに書き込む
	writeTemplateToFile(filepath.Join(pluginRootDir, "kong-plugin-"+pluginName+"-0.1.0-1.rockspec"), rockspecTemplate, pluginName)
	writeTemplateToFile(filepath.Join(srcDir, "handler.lua"), handlerTemplate, pluginName)
	writeTemplateToFile(filepath.Join(srcDir, "schema.lua"), schemaTemplate, pluginName)
}

func WriteGoTemplate(srcDir string, pluginRootDir string, pluginName string) {

	mainGoTemplate := `package main

import (
	"log"

	"github.com/Kong/go-pdk"
	"github.com/Kong/go-pdk/server"
)

func main() {
	server.StartServer(New, Version, Priority)
}

var Version = "1.0.0"
var Priority = 10

type Config struct {
	Message string
}

func New() interface{} {
	return &Config{}
}

func (conf Config) Access(kong *pdk.PDK) {
	log.Println("This is an example golang plugin handler, conf: ", conf)
}
`

	goModTemplate := `module github.com/your/repo

go 1.21

require github.com/Kong/go-pdk v0.11.0

require (
	github.com/ugorji/go/codec v1.2.12 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
`

	writeTemplateToFile(filepath.Join(srcDir, "main.go"), mainGoTemplate, pluginName)
	writeTemplateToFile(filepath.Join(srcDir, "go.mod"), goModTemplate, pluginName)
}

func WriteJSTemplate(srcDir string, pluginRootDir string, pluginName string) {
	// todo
}
