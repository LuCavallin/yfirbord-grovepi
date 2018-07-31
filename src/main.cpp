#include <Bridge.h>
#include <BridgeClient.h>
#include <ArduinoJson.h>
#include <PubSubClient.h>

// Networking
BridgeClient connection;

// Sensors and actuators libs
#include "Ultrasonic.h"
#include "DHT.h"

#define DHTTYPE DHT11

int ultrasonicPin = 2;
int DHTPin = 4;
int buzzerPin = 8;
int ledPin = 3;

DHT dht(DHTPin, DHTTYPE);
Ultrasonic ultrasonic(ultrasonicPin);

// Setup MQTT
void onMessageReceived(char *topic, byte *payload, unsigned int length)
{
    Serial.println("received");
    Serial.println(topic);
}

PubSubClient client(MQTT_HOSTNAME, MQTT_PORT, onMessageReceived, connection);

// Program
void setup()
{
    Bridge.begin();
    Serial.begin(115200);
}

void loop()
{
    // MQTT
    if (!client.connected())
    {
        if (client.connect("hytta", MQTT_USERNAME, MQTT_PASSWORD))
        {
            Serial.println("connected");
            client.publish("outTopic", "hello world");
            client.subscribe("outTopic");
        }
    }

    if (client.connected())
    {
        client.loop();
    }
}
