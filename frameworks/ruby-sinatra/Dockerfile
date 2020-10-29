FROM ruby:2.7.2-alpine

ENV APP_ENV=production
ENV BENCH_HOST=0.0.0.0
ENV BENCH_PORT=8080
ENV BENCH_WORKERS=1

WORKDIR /opt/src

COPY . /opt/src

RUN apk add openssl-dev g++ musl-dev make \
	&& gem install bundler \
	&& bundler install

ENTRYPOINT [ "./entrypoint.sh" ]
