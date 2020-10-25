#!/bin/bash
envsubst < /opt/openresty/conf/nginx.conf.tpl > /opt/openresty/conf/nginx.conf
nginx -c "/opt/openresty/conf/nginx.conf" "$@"
