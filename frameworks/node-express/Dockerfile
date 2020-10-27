FROM node:alpine

ENV BENCH_HOST=0.0.0.0
ENV BENCH_PORT=8080
ENV BENCH_WORKERS=1
ENV NODE_ENV=production

WORKDIR /opt/src
COPY . ./

RUN rm -rf node_modules \
	&& npm install \
	&& npm install -g pm2


ENTRYPOINT ["./entrypoint.sh"]
