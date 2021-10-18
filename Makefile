
# This makefile defines the following targets
#
#   - all (default) - formats the code and downloads vendor libs
#   - fmt - formats the code
#   - vendor - download all third party libraries and puts them inside vendor directory
#   - clean-vendor - removes third party libraries from vendor directory

.PHONY: all				##formats the code and downloads vendor libs
all: clean-vendor fmt vendor

.PHONY: fmt				##formats the code
fmt:
	@gci -w .
	@go fmt ./...
	@gofumpt -w .

.PHONY: vendor			##download all third party libraries and puts them inside vendor directory
vendor:
	@go mod vendor

.PHONY: clean-vendor			##removes third party libraries from vendor directory
clean-vendor:
	-@rm -rf vendor

.PHONY: lint				##runs code analysis tools
lint: clean-vendor
	go vet ./...
	golint ./...
	golangci-lint run ./...

.PHONY: help				##show this help message
help:
	@echo "usage: make [target]\n"; echo "options:"; \fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | sed 's/.PHONY:*//' | sed -e 's/^/  /'; echo "";

CONTROLLER_GEN = controller-gen
CRD_OPTIONS ?= "crd:trivialVersions=true,preserveUnknownFields=false"

controller-gen: ## Download controller-gen locally if necessary.
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.6.1

## The controller-gen tool doesn't generate deep copy and crd files for apis with own go.mod file, i.e. /apis/config/go.mod
## The workaround is:
##   1. move module's go.mod file to a temporary file
##   2  copy main go.mod file to a temporary file to preserve changes
##   3. run "go mod tidy" to ensure main go.mod file contains all required dependencies
##   4. run controller-gen command
##   5. restore go.mod files

manifests: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	for API in ./apis/*; do mv $$API/go.mod $$API/go_temp.mod; done
	cp ./go.mod ./go_temp.mod
	go mod tidy
	$(CONTROLLER_GEN) $(CRD_OPTIONS) paths="./..." output:crd:artifacts:config=crd
	for API in ./apis/*; do mv $$API/go_temp.mod $$API/go.mod; done
	mv ./go_temp.mod ./go.mod

generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	for API in ./apis/*; do mv $$API/go.mod $$API/go_temp.mod; done
	cp ./go.mod ./go_temp.mod
	go mod tidy
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	for API in ./apis/*; do mv $$API/go_temp.mod $$API/go.mod; done
	mv ./go_temp.mod ./go.mod
