FROM python:3.9-alpine

ENV HOST_ADDR=""

WORKDIR /opt/src

COPY . /opt/src

RUN pip install pipenv \
	&& pipenv install --dev

ENTRYPOINT ["pipenv", "run", "pytest", "-s"]
