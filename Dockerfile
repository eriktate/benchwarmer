FROM ubuntu AS builder

WORKDIR /opt

RUN apt-get update \
	&& apt-get install -y git gcc make libssl-dev zlib1g-dev \
	&& git clone https://github.com/giltene/wrk2 \
	&& cd wrk2 \
	&& make

WORKDIR /opt/wrk2

ENTRYPOINT ["./wrk"]
