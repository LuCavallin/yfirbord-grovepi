import BlynkLib
import time
import os
from BlynkTimer import BlynkTimer
from grove.grove_led import GroveLed
from grove.factory import Factory
from grove.grove_temperature_humidity_sensor import DHT

# Load configuration
from os.path import join, dirname
from dotenv import load_dotenv
dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path)

# Initialize Blynk with a BlynkTimer Instance
BLYNK_AUTH = os.getenv("BLYNK_AUTH")
blynk = BlynkLib.Blynk(BLYNK_AUTH)
timer = BlynkTimer()

# Handlers


def toggle_led(seconds):
    led = GroveLed(5)
    led.on()
    time.sleep(seconds)
    led.off()


def buzz(seconds):
    buzzer = Factory.getGpioWrapper("Buzzer", 4)
    buzzer.on()
    time.sleep(seconds)
    buzzer.off()


def dht():
    dht11 = DHT("11", 16)
    humidity, temperature = dht11.read()
    return {temperature: temperature, humidity: humidity}


# Add Timers
timer.set_timeout(1, toggle_led)

# Run
while True:
    blynk.run()
    timer.run()
    toggle_led(2)
    buzz(2)
    print(dht())
