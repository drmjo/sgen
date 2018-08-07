DOCKER_IMAGE:=drmjo/site-generator-go-cli:latest
CURRENT_DIRECTORY:=$(shell pwd)
GO_USER:=go
REPOSITORY:=github.com/drmjo/sgen
WORKING_DIRECTORY:=/home/${GO_USER}/app

.PHONY: build
build:
	go get -v ${REPOSITORY}
	go install -v ${REPOSITORY}

.PHONY: install
install:
	go install -v ${REPOSITORY}

# below commands should be used outside of a docker container
# this is to provision a development container
.PHONY: build-cli
build-cli:
	/usr/local/bin/docker build \
		-t ${DOCKER_IMAGE} \
		-f dockerfiles/cli.Dockerfile \
		--build-arg USER=${GO_USER} \
		dockerfiles

# build the cli
.PHONY: cli
cli: build-cli
	/usr/local/bin/docker run -it --rm \
		-v /tmp/go-site-generator/.cache:/home/go/.cache \
		-v /tmp/go-site-generator/.go-src:${WORKING_DIRECTORY}/src \
		-v ${CURRENT_DIRECTORY}:${WORKING_DIRECTORY}/src/${REPOSITORY} \
		-v ${CURRENT_DIRECTORY}/Makefile:${WORKING_DIRECTORY}/Makefile \
		-v ${CURRENT_DIRECTORY}/terraform:${WORKING_DIRECTORY}/terraform \
		-v ${CURRENT_DIRECTORY}/certs:/etc/sgen/certs \
		-e DOCKER_TLS_VERIFY=1 \
		-e DOCKER_HOST=tcp://192.168.99.100:2376 \
		-e DOCKER_CERT_PATH=/etc/sgen/certs \
		--hostname sgen.lab.mjo.io \
		--name site-generator-cli \
		${DOCKER_IMAGE} bash
