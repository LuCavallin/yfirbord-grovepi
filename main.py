import BlynkLib
from BlynkTimer import BlynkTimer
from grove.grove_led import GroveLed
import time

BLYNK_AUTH = 'YourAuthToken'

# Initialize Blynk with a BlynkTimer Instance
blynk = BlynkLib.Blynk(BLYNK_AUTH)
timer = BlynkTimer()

def hello_world():
    print("Hello World!")

# Add Timers
timer.set_timeout(2, hello_world)


led = GroveLed(5)
while True:
    blynk.run()
    timer.run()
    led.on()
    time.sleep(1)
    led.off()
    time.sleep(1)
