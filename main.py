import BlynkLib
from BlynkTimer import BlynkTimer
from grove.grove_led import GroveLed
from grove.factory import Factory
from grove.grove_temperature_humidity_sensor import DHT
import time

# Load configuration
from os.path import join, dirname
from dotenv import load_dotenv
dotenv_path = join(dirname(__file__), '.env')
load_dotenv(dotenv_path)

# Initialize Blynk with a BlynkTimer Instance
BLYNK_AUTH = os.getenv("BLYNK_AUTH");
blynk = BlynkLib.Blynk(BLYNK_AUTH)
timer = BlynkTimer()

# Handlers
led = GroveLed(5)
def toggle_led(seconds):
    led.on()
    time.sleep(seconds)
    led.off()

buzzer = Factory.getGpioWrapper("Buzzer", 4)
def buzz(seconds):
    buzzer.on()
    time.sleep(seconds)
    buzzer.off()

dht11 = DHT("11", 16)
def dht()
    humidity, temperature = dht11.read()
    return {temperature: temp, humidity: humidity}


# Add Timers
timer.set_timeout(1, toggle_led)

# Run
while True:
    blynk.run()
    timer.run()
    toggle_led(2)
    buzz(2)
