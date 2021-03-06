worker_processes ${BENCH_WORKERS};
daemon off;
error_log /dev/stderr;
events {
	worker_connections 1024;
}
http {
	server {
		server_name localhost;
		listen 8080;
		location /hello {
			default_type text/html;
			content_by_lua_block {
				ngx.print("Hello, World!")
			}
		}

		location /json {
			default_type application/json;
			content_by_lua_block {
				local cjson = require("cjson")
				ngx.req.read_body()
				local json_req = cjson.decode(ngx.var.request_body)
				local json_res = { msg=(json_req.greeting .. " " .. json_req.name) }
				ngx.say(cjson.encode(json_res))
			}
		}
	}
}
