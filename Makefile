RANDTAG:=$(shell openssl rand -hex 5)

local:
	operator-sdk build docker.netzlink.com/operators/sormas-operator:$(RANDTAG)
	sed -i 's|\w+.\w+.\w+\/\w+\/sormas-operator:[0-9,a-z,A-Z]+|docker.netzlink.com/operators/sormas-operator:$(RANDTAG)|g' deploy/operator.yaml

push:	local
	docker push docker.netzlink.com/operators/sormas-operator:$(RANDTAG)

apply:	push
	kubectl apply -f deploy/crds/sormas.netzlink.com_sormas_crd.yaml
	kubectl apply -f deploy

cluster:
	kind create cluster --name sormas-operator 
	kubectl cluster-info --context kind-sormas-operator

diagrams:
	plantuml -tsvg doc/sormas-operator.md