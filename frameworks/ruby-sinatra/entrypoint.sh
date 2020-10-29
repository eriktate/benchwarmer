#!/bin/sh

bundler exec thin -R config.ru -a $BENCH_HOST -p $BENCH_PORT start
