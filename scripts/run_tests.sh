docker build -t api-tests -f tests/Dockerfile ./tests

for framework in ./frameworks/*/
do
	framework=${framework%*/}
	framework=${framework##*/}
	docker run \
		-it \
		--network=benchwarmer_default \
		-e HOST_ADDR="${framework}:8080" \
		api-tests
done
