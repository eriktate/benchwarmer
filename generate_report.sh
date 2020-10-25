#!/bin/bash

./run.sh
pushd aggregate-report
pipenv run python main.py
popd
