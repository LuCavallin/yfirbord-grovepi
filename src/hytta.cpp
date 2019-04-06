/*
 * Project hytta
 * Description: Monitor the environment in your house with a Particle Photon.
 * Author: Luca Cavallin <me@lucavall.in>
 * Date: 05/04/2019
 */

#include "Particle.h"
#include "blynk.h"
#include "Ultrasonic.h"

void setup();
void sendRange();
void loop();
#line 12 "/Users/luca/Projects/hytta/src/hytta.ino"
#define BLYNK_PRINT Serial
//#define BLYNK_DEBUG
char blynkAuth[] = "74dab704d1824061bb43cf03df866244";
BlynkTimer timer;

#define PIN_BUZZER D6
#define PIN_LED D4
#define PIN_ULTRASONIC D2
Ultrasonic ultrasonic(PIN_ULTRASONIC);

// setup() runs once, when the device is first turned on.
void setup()
{
  Serial.begin(9600);
  pinMode(PIN_BUZZER, OUTPUT);
  pinMode(PIN_LED, OUTPUT);
  Serial.begin(9600);

  // Setup Blynk
  delay(5000);
  timer.setInterval(1000L, sendRange);
  Blynk.begin(blynkAuth);
}

// Read from ultrasonic ranger and write to Blynk V0
void sendRange()
{
  float cm = ultrasonic.MeasureInCentimeters();
  Blynk.virtualWrite(0, cm);
}

// loop() runs over and over again, as quickly as it can execute.
void loop()
{
  Blynk.run();
  timer.run();
}
