
local example = {
	VERSION = "1.0.0",
	PRIORITY = 10,
}

function example:access(conf)
	-- plugin logic here
	kong.log("This is an example plugin handler, conf: ", conf)
end

return example
