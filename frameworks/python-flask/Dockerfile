FROM python:3.9-alpine

EXPOSE 8080
ENV BENCH_HOST=0.0.0.0
ENV BENCH_PORT=8080
ENV BENCH_WORKERS=1

WORKDIR /opt/src

COPY . /opt/src

RUN pip install pipenv \
	&& pipenv install \
	&& pipenv install gunicorn


ENTRYPOINT ["pipenv", "run", "gunicorn", "-w", "${BENCH_WORKERS}", "-b", "${BENCH_HOST}:${BENCH_PORT}", "main:app"]
