setup-pi:
	sudo rpi-update -y
	sudo apt-get update
	sudo apt-get -y upgrade
	sudo adduser pi gpio
	sudo chown root.gpio /dev/gpiomem
	sudo chmod g+rw /dev/gpiomem
	sudo reboot

setup:
	echo "deb https://seeed-studio.github.io/pi_repo/ stretch main" | sudo tee /etc/apt/sources.list.d/seeed.list
	curl https://seeed-studio.github.io/pi_repo/public.key | sudo apt-key add -
	sudo apt update
	sudo apt install -y python3-mraa python3-upm python3-rpi.gpio
	pip3 install -r requirements.txt

copy:
	scp . pi@raspberrypi.local:/home/pi/hytta

run:
	ssh -t pi@raspberrypi.local 'python3 ./hytta/main.py'
