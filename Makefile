setup:
	sudo pip install -U setuptools pip platformio

build:
	git pull
	platformio build

run:
	ssh -t pi@raspberrypi.local './hytta'
