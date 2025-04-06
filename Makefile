VERSION ?= latest

docker-build:
	@echo "building version ${VERSION}"
	docker build . --build-arg IMAGE_TAG=${VERSION} -t mgmt-file-transfer:${VERSION}