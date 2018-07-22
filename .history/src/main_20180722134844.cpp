#include <ArduinoJson.h>

#include <MQTTClient.h>
#include <MQTT.h>
#include <system.h>

#include "Ultrasonic.h"
#include "DHT.h"

#define DHTTYPE DHT11

MQTTClient client;
int ultrasonicPin = 2;
int DHTPin = 4;
int buzzerPin = 8;
int ledPin = 3;

DHT dht(DHTPin, DHTTYPE);
Ultrasonic ultrasonic(ultrasonicPin);

void connect()
{
    Serial.print("\nConnecting to MQTT...");
    while (!client.connect("arduino", "try", "try"))
    {
        Serial.print(".");
        delay(1000);
    }

    Serial.println("\nconnected!");

    client.subscribe("/hello");
    // client.unsubscribe("/hello");
}

void messageReceived(String &topic, String &payload)
{
    Serial.println("incoming: " + topic + " - " + payload);
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

void setup()
{
    // MQTT
    // client.begin("broker.shiftr.io", net);
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

    long ultrasonicCentimeters;

    ultrasonicCentimeters = ultrasonic.MeasureInCentimeters(); // two measurements should keep an interval
    Serial.println(ultrasonicCentimeters);
    if (ultrasonicCentimeters > 50)
    {
        toggleDigital(buzzerPin, 250);
    }
    else
    {
        toggleDigital(ledPin, 1000);
    }

    // Reading temperature or humidity takes about 250 milliseconds!
    // Sensor readings may also be up to 2 seconds 'old' (its a very slow sensor)
    float h = dht.readHumidity();
    float t = dht.readTemperature();

    // check if returns are valid, if they are NaN (not a number) then something went wrong!
    if (isnan(t) || isnan(h))
    {
        Serial.println("Failed to read from DHT");
    }
    else
    {
        printDHT(h, t);
    }

    Serial.println("\n");
    delay(1000);
}
