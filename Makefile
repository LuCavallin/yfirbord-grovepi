setup:
	echo "deb https://seeed-studio.github.io/pi_repo/ stretch main" | sudo tee /etc/apt/sources.list.d/seeed.list
	curl https://seeed-studio.github.io/pi_repo/public.key | sudo apt-key add -
	sudo apt update
	sudo apt install python3-mraa python3-upm
	sudo apt update
	sudo apt install python3-rpi.gpio
	sudo pip3 install rpi_ws281x
	pip3 install -r requirements.txt

copy:
	scp . pi@raspberrypi.local:/home/pi/hytta

run:
	ssh -t pi@raspberrypi.local 'python3 ./hytta/main.py'
