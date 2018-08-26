GOOS=linux
GOARCH=arm
GOARM=6
FILES=hytta

build:
	env GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build ./cmd/hytta

copy:
	scp $(FILES) pi@raspberrypi.local:/home/pi

run:
	ssh -t pi@raspberrypi.local './hytta'
