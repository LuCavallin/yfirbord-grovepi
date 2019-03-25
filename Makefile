setup:
	curl -sL https://github.com/Seeed-Studio/grove.py/raw/master/install.sh | sudo bash -s -
	pip install -r requirements.txt

copy:
	scp . pi@raspberrypi.local:/home/pi/hytta

run:
	ssh -t pi@raspberrypi.local 'python3 ./hytta/main.py'
