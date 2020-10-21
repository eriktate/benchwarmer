FROM python:3.9-slim

EXPOSE 8080
ENV BENCH_HOST=0.0.0.0
ENV BENCH_PORT=8080

WORKDIR /opt/src

COPY . /opt/src

RUN pip install pipenv \
	&& pipenv install

ENTRYPOINT ["pipenv", "run", "python", "main.py"]
