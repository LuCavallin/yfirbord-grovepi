#include <Bridge.h>
#include <BridgeClient.h>
#include <ArduinoJson.h>
#include <PubSubClient.h>
#include <Time.h>

// Networking
BridgeClient connection;

// Sensors and actuators libraries
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
#define INCOMING_TOPIC "to-arduino"
#define OUTGOING_TOPIC "from-arduino"

void onMessageReceived(char *topic, byte *payload, unsigned int length);
PubSubClient client(MQTT_HOSTNAME, MQTT_PORT, onMessageReceived, connection);

void onMessageReceived(char *topic, byte *payload, unsigned int length)
{
    Serial.print("Message received from topic: ");
    Serial.print(topic);
    Serial.println();

    // Get and parse request
    String json = (char *)payload;
    StaticJsonBuffer<200> jsonBuffer;
    JsonObject &incomingRoot = jsonBuffer.parseObject(json);

    // Test if parsing succeeds.
    if (!incomingRoot.success())
    {
        Serial.println("Failed to parse request.");
        return;
    }

    // Prepare response
    const char *sensor = incomingRoot["sensor"];

    JsonObject &outgoingRoot = jsonBuffer.createObject();
    outgoingRoot["id"] = "1394y71289312";
    outgoingRoot["timestamp"] = now();
    outgoingRoot["sensor"] = sensor;
    JsonArray &data = outgoingRoot.createNestedArray("data");

    // TODO
    // data.add(48.756080);
    // data.add(2.302038);

    // char *response;
    // outgoingRoot.printTo(response, 128);

    // Send response
    Serial.print("Publishing message: ");
    Serial.print(sensor);
    Serial.print(" to topic: ");
    Serial.print(OUTGOING_TOPIC);
    Serial.print(".");
    Serial.println();

    client.publish("from-arduino", response);
}

// Program
void setup()
{
    Serial.println("Setting up Bridge and Serial...");
    Bridge.begin();
    Serial.begin(115200);
}

void loop()
{
    // MQTT
    if (!client.connected())
    {
        Serial.println("Connecting to MQTT...");
        if (client.connect("hytta", MQTT_USERNAME, MQTT_PASSWORD))
        {
            Serial.println("Connected!");
            client.subscribe(INCOMING_TOPIC);
        }
    }

    if (client.connected())
    {
        client.loop();
    }
}
