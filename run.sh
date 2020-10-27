#!/bin/bash
run () {
	framework=$1
	benchmark=$2
	echo "Running ${benchmark} for ${framework}"
	docker run \
		-it \
		--network=benchwarmer_default \
		-v $(pwd)/benchmarks:/opt/wrk2/benchmarks \
		wrk2 -c1000 -t500 -d30s -R10000 -s ./benchmarks/${benchmark}.lua http://${framework}:8080/${benchmark} | grep "JSON Output:" -A5000 | tail -n+2 > ./reports/${framework}_${benchmark}.json
}

if [ -z "$1" ]; then
	for framework in ./frameworks/*/
	do
		for bench in ./benchmarks/*
		do
			bench=${bench%.*}
			bench=${bench##*/}
			framework=${framework%*/}
			framework=${framework##*/}

			run $framework $bench
			sleep 5s
		done
	done
else
	run $1 $2
fi
