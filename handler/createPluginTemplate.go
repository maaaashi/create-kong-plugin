package handler

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreatePluginTemplate(pluginName string, language string, mkdir bool) {
	pluginRootDir := "."

	if mkdir {
		pluginRootDir = filepath.Join(".", pluginName)
	}
	srcDir := filepath.Join(pluginRootDir, "/src")

	if err := os.MkdirAll(srcDir, os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	switch language {
	case "Lua":
		WriteLuaTemplate(srcDir, pluginRootDir, pluginName)
	case "Go":
		WriteGoTemplate(srcDir, pluginRootDir, pluginName)
	case "JavaScript":
		WriteJSTemplate(srcDir, pluginRootDir, pluginName)
	}

	fmt.Printf("Kong plugin template for '%s' created successfully!\n", pluginName)
}
