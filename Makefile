PROTOS := $(wildcard *.proto) $(wildcard */*.proto)
PBGO := $(PROTOS:.proto=.pb.go)

all: $(PBGO)
	go mod tidy
	go build .

clean:
	rm $(PBGO)

%.pb.go: %.proto
	protoc --go_out=. $<

.PHONY: all clean
