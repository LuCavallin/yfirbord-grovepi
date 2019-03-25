import BlynkLib
from BlynkTimer import BlynkTimer

BLYNK_AUTH = 'YourAuthToken'

# Initialize Blynk
blynk = BlynkLib.Blynk(BLYNK_AUTH)

# Create BlynkTimer Instance
timer = BlynkTimer()


# Will only run once after 2 seconds
def hello_world():
    print("Hello World!")


# Will Print Every 5 Seconds
def print_me():
    print("Thanks!")


# Add Timers
timer.set_timeout(2, hello_world)
timer.set_interval(5, print_me)


while True:
    blynk.run()
    timer.run()
