IMAGE ?= ocean-demo-mac
ARCH ?= amd64
CONTAINER ?= ocean-demo-build

build:
docker build --build-arg ARCH=$(ARCH) -f Dockerfile.mac -t $(IMAGE) .

extract:
docker create --name $(CONTAINER) $(IMAGE)
docker cp $(CONTAINER):/ocean-demo ./ocean-demo
docker rm $(CONTAINER)

run: extract
./ocean-demo

.PHONY: build extract run
