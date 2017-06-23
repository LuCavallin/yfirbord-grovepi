GOOS=linux
GOARCH=arm
FILES=yfirbord

build:
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build ./cmd/yfirbord

build-default:
	go build ./cmd/yfirbord

copy:
	scp $(FILES) yfirbord@yfirbord-grovepi.local:/home/yfirbord/