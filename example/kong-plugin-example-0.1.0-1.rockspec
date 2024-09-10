
package = "kong-plugin-example"
version = "0.1.0-1"

source = {
	url = ""
}

build = {
    type = "builtin",
    modules = {
        ["kong.plugins.example.handler"] = "src/handler.lua",
        ["kong.plugins.example.schema"] = "src/schema.lua"
    }
}
	