require 'sinatra'

set :bind, ENV["BENCH_HOST"]
set :port, ENV["BENCH_PORT"]
set :logging, nil

get '/hello' do
  "Hello, World!"
end

post '/json' do
  content_type :json
  request.body.rewing
  payload = JSON.parse request.body.read
  { msg: "%s %s" % [payload["greeting"], payload["name"]] }.to_json
end
