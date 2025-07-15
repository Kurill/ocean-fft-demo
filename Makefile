IMAGE ?= ocean-demo-mac
ARCH ?= amd64
CONTAINER ?= ocean-demo-build

build:
	@if ! docker image inspect $(IMAGE) > /dev/null 2>&1; then \
	docker build --build-arg ARCH=$(ARCH) -f Dockerfile.mac -t $(IMAGE) .; \
	else \
	echo "$(IMAGE) image already exists"; \
	fi

extract: build
	docker create --name $(CONTAINER) $(IMAGE)
	docker cp $(CONTAINER):/ocean-demo ./ocean-demo
	docker rm $(CONTAINER)

run: extract
	./ocean-demo

.PHONY: build extract run
