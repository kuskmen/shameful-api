.PHONY: build-debug build-prod

build-debug:
	docker build -t shameful-api-debug -f ../build/package/Dockerfile.debug .
build-prod:
	docker build -t shameful-api-prod -f ../build/package/Dockerfile .

.PHONY: run-debug run-prod
run-debug:
	docker run --rm \
		--publish 8093:8093 \
		--publish 2345:2345 \
		--security-opt=seccomp:unconfined \
		--name shameful-api-debug \
		shameful-api-debug
			
run-prod:
	docker run --rm \
		--publish 8093:8093 \
		--publish 2345:2345 \
		--name shameful-api-prod \
		shameful-api-prod
	
