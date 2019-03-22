FILES=hytta

install:
	pip install -U platformio

build:
	platformio run

copy:
	scp $(FILES) pi@raspberrypi.local:/home/pi

run:
	ssh -t pi@raspberrypi.local './hytta'
