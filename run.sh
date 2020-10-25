#!/bin/bash
run () {
	framework=$1
	benchmark=$2
	echo "Running ${benchmark} for ${framework}"
	docker run \
		-it \
		--network=benchwarmer_default \
		-v $(pwd)/wrk-scripts:/opt/wrk2/wrk-scripts \
		wrk2 -c1000 -t500 -d30s -R10000 -s ./wrk-scripts/${benchmark}.lua http://${framework}:8080/${benchmark} | grep "JSON Output:" -A5000 | tail -n+2 > ./reports/${framework}_${benchmark}.json

}

for framework in ./frameworks/*/
do
	framework=${framework%*/}
	framework=${framework##*/}

	run $framework hello
	run $framework json
done

