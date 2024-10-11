local plugin = {
	VERSION = "1.0.0",
	PRIORITY = 10,
}

function plugin:access(conf)
	-- plugin logic here
	kong.log("This is an example lua plugin handler, conf: ", conf)
end

return plugin
