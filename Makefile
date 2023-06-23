
CONTROLLER_GEN ?= go run sigs.k8s.io/controller-tools/cmd/controller-gen@v0.12.0

manifest:
	$(CONTROLLER_GEN) rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases
	kubectl create -k config/crd --dry-run=client -o yaml > deploy/crds.yaml

generate: manifest
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."