#!/usr/bin/lua5.3

http = require("socket.http")

function ReturnIP(url)
	local resp = http.request(url)

	return resp
end

print(ReturnIP('http://ip.wang-li.top:93/4u6385IP'))