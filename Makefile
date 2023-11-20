.PHONY: api
# generate api
api:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) api'

.PHONY: config
# generate config
config:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) config'


.PHONY: generate
# generate wire
generate:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) generate'

.PHONY: all
# all 
all:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) all'

.PHONY: build
# build 
build:
	find app -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && pwd && $(MAKE) build'

generate-product-stub:
	kratos proto client api/product/v1/product.proto