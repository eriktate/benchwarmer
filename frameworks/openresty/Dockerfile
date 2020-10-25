FROM ubuntu

WORKDIR /opt/openresty

ENV PATH=/usr/local/openresty/nginx/sbin:$PATH

RUN apt-get update \
	&& apt-get install -y wget gnupg ca-certificates lsb-release gettext-base \
	&& wget -O - https://openresty.org/package/pubkey.gpg | apt-key add - \
	&& echo "deb http://openresty.org/package/ubuntu $(lsb_release -sc) main" \
    | tee /etc/apt/sources.list.d/openresty.list \
	&& apt-get update \
	&& apt-get install -y openresty


COPY entrypoint.sh ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
