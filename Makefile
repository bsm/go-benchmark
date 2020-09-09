GO_MOD_FILES=$(shell find . -name 'go.mod')

default: bench

bench: $(patsubst %/go.mod,bench/%,$(GO_MOD_FILES))

bench/%: %
	cd $< && go test ./... -run=NONE -bench=. -benchmem
