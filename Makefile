run:
	docker run \
		-it \
		--network=benchwarmer_default \
		-v $(shell pwd)/wrk-scripts:/opt/wrk2/wrk-scripts \
		wrk2 -c1000 -t500 -d10s -R10000 -s ./wrk-scripts/${benchmark}.lua http://${target}:8080/${benchmark}
