REGISTRY:=docker.io
USER:=nicokahlert
IMAGE:=sormas-operator

RANDTAG:=$(shell openssl rand -hex 5)

local:
	operator-sdk build $(REGISTRY)/$(USER)/$(IMAGE):$(RANDTAG)
	sed -Ei 's|$(REGISTRY)/$(USER)/$(IMAGE):\w+|$(REGISTRY)/$(USER)/$(IMAGE):$(RANDTAG)|g' deploy/operator.yaml

push:	local
	docker push $(REGISTRY)/$(USER)/$(IMAGE):$(RANDTAG)

apply:	push
	kubectl apply -f deploy/crds/sormas.netzlink.com_sormas_crd.yaml
	kubectl apply -f deploy

cluster:
	kind create cluster --name $(IMAGE)
	kubectl cluster-info --context kind-$(IMAGE)

diagrams:
	plantuml -tsvg doc/$(IMAGE).md