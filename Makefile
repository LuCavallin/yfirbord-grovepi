GOOS=linux
GOARCH=arm
FILES=hytta sensors.json

build:
	env GOOS=$(GOOS) GOARCH=$(GOARCH) go build ./cmd/hytta

build-default:
	go build ./cmd/hytta

copy:
	scp $(FILES) hytta@hytta-grovepi.local:/home/hytta/
