
#include <DHT.h>
#include <Ultrasonic.h>
#include <Bridge.h>
#include <BridgeServer.h>
#include <BridgeClient.h>

#define DHT_PIN 4
#define DHT_TYPE DHT11
#define BUZZER_PIN 8
#define LED_BLUE_PIN 3
#define ULTRASONIC_RANGER_PIN 2

DHT dht(DHT_PIN, DHT_TYPE);
Ultrasonic ultrasonic(ULTRASONIC_RANGER_PIN);
BridgeServer server;

void setup()
{
  // Sensors setup
  dht.begin();
  pinMode(BUZZER_PIN, OUTPUT);
  pinMode(LED_BLUE_PIN, OUTPUT);

  // External and openwrt communications setup
  Bridge.begin();
  server.begin();
}

void loop()
{
  BridgeClient client = server.accept();

  if (client)
  {
    process(client);
    client.stop();
  }

  delay(1000);
}

void process(BridgeClient client)
{
  String command = client.readStringUntil('/');

  if (command == "digital")
  {
    toggleDigital(13, 1000);
  }
  if (command == "buzzer")
  {
    toggleDigital(BUZZER_PIN, 1000);
  }
  if (command == "ultrasonic_ranger")
  {
    long ultrasonicCentimeters = ultrasonic.MeasureInCentimeters();
    client.write(ultrasonicCentimeters);
  }
  if (command == "humidity")
  {
    float h = dht.readHumidity();
    client.write(h);
  }
  if (command == "temperature")
  {
    float t = dht.readTemperature();
    client.write(t);
  }
}

void toggleDigital(int pin, int delayTime)
{
  digitalWrite(pin, HIGH);
  delay(delayTime);
  digitalWrite(pin, LOW);
}
