RANDTAG:=$(shell openssl rand -hex 5)

local:
	operator-sdk build docker.netzlink.com/operators/sormas-operator:$(RANDTAG)
	sed -i 's|REPLACE_IMAGE|docker.netzlink.com/operators/sormas-operator:$(RANDTAG)|g' deploy/operator.yaml

push:
	operator-sdk build docker.netzlink.com/operators/sormas-operator:$(RANDTAG)
	sed -i 's|REPLACE_IMAGE|docker.netzlink.com/operators/sormas-operator:$(RANDTAG)|g' deploy/operator.yaml
	docker push docker.netzlink.com/operators/sormas-operator:$(RANDTAG)
