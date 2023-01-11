.phony: all chk

chk:
	goimports -w ./cmd
	goimports -w ./pkg

all: chk
	go build -mod=mod -o build/backlight ./cmd/backlight

