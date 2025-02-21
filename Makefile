
export GO111MODULE=on

.PHONY: test
test:
	go test ./pkg/... ./cmd/... -coverprofile cover.out

.PHONY: bin
bin: fmt vet
	go build -o bin/kubectl-cnp_viz github.com/mostafahussein/kubectl-cnp-viz/cmd/plugin

.PHONY: fmt
fmt:
	go fmt ./pkg/... ./cmd/...

.PHONY: vet
vet:
	go vet ./pkg/... ./cmd/...

.PHONY: kubernetes-deps
kubernetes-deps:
	go get k8s.io/client-go@v0.32.2
	go get k8s.io/api@kubernetes-0.32.2
	go get k8s.io/apimachinery@kubernetes-0.32.2
	go get k8s.io/cli-runtime@kubernetes-0.32.2

.PHONY: setup
setup:
	make -C setup