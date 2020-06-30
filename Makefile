RANDTAG:=$(shell openssl rand -hex 5)

local:
	operator-sdk build docker.netzlink.com/operators/sormas-operator:$(RANDTAG)
	sed -i 's|\w+.\w+.\w+\/\w+\/sormas-operator:[0-9,a-z,A-Z]+|docker.netzlink.com/operators/sormas-operator:$(RANDTAG)|g' deploy/operator.yaml

push:
	operator-sdk build docker.netzlink.com/operators/sormas-operator:$(RANDTAG)
	sed -i 's|\w+.\w+.\w+\/\w+\/sormas-operator:[0-9,a-z,A-Z]+|docker.netzlink.com/operators/sormas-operator:$(RANDTAG)|g' deploy/operator.yaml
	docker push docker.netzlink.com/operators/sormas-operator:$(RANDTAG)

diagrams:
	plantuml -tsvg doc/sormas-operator.md
