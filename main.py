import BlynkLib
from BlynkTimer import BlynkTimer
from grove.grove_led import GroveLed
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
def toggle_led():
    led.on()
    time.sleep(1)
    led.off()

# Add Timers
timer.set_timeout(1, toggle_led)

# Run
while True:
    blynk.run()
    timer.run()
