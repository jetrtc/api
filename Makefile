PROTOS := rest_service.proto auth_service.proto video_service.proto
PBGO := $(PROTOS:.proto=.pb.go)

all: $(PBGO)

clean:
	rm *.pb.go

%.pb.go: %.proto
	protoc --go_out=. $<

.PHONY: all clean
