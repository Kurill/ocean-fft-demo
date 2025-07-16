BUILDER_IMAGE ?= ocean-demo-builder
ARCH ?= amd64

container:
	docker build -f Dockerfile.build -t $(BUILDER_IMAGE) .

build:
	docker run --rm -v $(PWD):/src --entrypoint "" $(BUILDER_IMAGE) \
	bash -c 'xgo --targets=darwin/$(ARCH) --pkg cmd/ocean-demo -out ocean-demo . && mv /build/ocean-demo-darwin-$(ARCH) /src/ocean-demo'

run: build
	./ocean-demo

clean:
	rm -f ocean-demo

.PHONY: container build run clean
