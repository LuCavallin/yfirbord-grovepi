FILES=hytta

install:
	pip install -U platformio
	platformio run

build:
	git pull
	platformio build

copy:
	scp $(FILES) pi@raspberrypi.local:/home/pi

run:
	ssh -t pi@raspberrypi.local './hytta'
