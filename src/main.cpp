#include <Bridge.h>
#include <BridgeClient.h>
#include <MQTT.h>
#include <ArduinoJson.h>

// Sensors and actuators libs
#include "Ultrasonic.h"
#include "DHT.h"

// Networking
BridgeClient net;
MQTTClient client;

// Sensors and actuators
#define DHTTYPE DHT11

int ultrasonicPin = 2;
int DHTPin = 4;
int buzzerPin = 8;
int ledPin = 3;

DHT dht(DHTPin, DHTTYPE);
Ultrasonic ultrasonic(ultrasonicPin);

// Program
void connect()
{
    Serial.print("Connecting...");
    while (!client.connect("arduino", MQTT_USERNAME, MQTT_PASSWORD))
    {
        Serial.print(".");
        delay(1000);
    }

    Serial.println("\nConnected!");
    client.subscribe("/hello");
}

void messageReceived(String &topic, String &payload)
{
    Serial.println("incoming: " + topic + " - " + payload);
}

void setup()
{
    Bridge.begin();
    Serial.begin(115200);

    client.begin(MQTT_HOSTNAME, net);
    client.onMessage(messageReceived);
    connect();

    Serial.begin(9600);
    dht.begin();
    pinMode(buzzerPin, OUTPUT);
    pinMode(ledPin, OUTPUT);
}

void loop()
{
    // MQTT
    client.loop();
    if (!client.connected())
    {
        connect();
    }

    // read message and operate

    delay(1000);
}

void toggleDigital(int pin, int delayTime)
{
    digitalWrite(pin, HIGH);
    delay(delayTime);
    digitalWrite(pin, LOW);
}

void printDHT(float h, int t)
{
    Serial.print("Humidity: ");
    Serial.print(h);
    Serial.print("\t Temperature: ");
    Serial.print(t);
    Serial.println(" *C");
}
