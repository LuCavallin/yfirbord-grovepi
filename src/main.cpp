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

#define SENSOR_READ 0
#define SENSOR_WRITE 1
#define DHTTYPE DHT11

#define PIN_ULTRASONIC 2
#define PIN_DHT 4
#define PIN_BUZZER 8
#define PIN_LED 3
#define PIN_DHT 4

DHT dht(PIN_DHT, DHTTYPE);
Ultrasonic ultrasonic(PIN_ULTRASONIC);

// Setup MQTT
#define INCOMING_TOPIC "to-arduino"
#define OUTGOING_TOPIC "from-arduino"

void onMessageReceived(char *topic, byte *payload, unsigned int length);
String handleRequest(const char *sensor, char *params);
void toggleDigital(int pin, int delayTime);
PubSubClient client(MQTT_HOSTNAME, MQTT_PORT, onMessageReceived, connection);

void onMessageReceived(char *topic, byte *payload, unsigned int length)
{
    // Debug message
    Serial.print("Message received from topic: ");
    Serial.print(topic);
    Serial.println();

    // Get and parse request
    String json = (char *)payload;
    StaticJsonBuffer<MQTT_MAX_PACKET_SIZE> jsonBuffer;
    JsonObject &incomingRoot = jsonBuffer.parseObject(json);

    // Test if parsing succeeds.
    if (!incomingRoot.success())
    {
        Serial.println("Failed to parse request.");
        return;
    }

    // Prepare response and add metadata
    const char *sensor = incomingRoot["sensor"];
    JsonObject &outgoingRoot = jsonBuffer.createObject();
    // @TODO find way to have a unique ID for the device
    outgoingRoot["id"] = "1394y71289312";
    outgoingRoot["timestamp"] = now();
    outgoingRoot["sensor"] = sensor;
    outgoingRoot["data"] = handleRequest(sensor, incomingRoot["params"]);

    // Converting response to needed type
    char response[MQTT_MAX_PACKET_SIZE];
    outgoingRoot.printTo(response, sizeof(response));

    // Send response, debug message
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

    // Setup Sensors
    dht.begin();
    pinMode(PIN_BUZZER, SENSOR_WRITE);
    pinMode(PIN_LED, SENSOR_WRITE);
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

// This needs serious re-thought
String handleRequest(const char *sensor, char *params)
{
    if (strcmp(sensor, "temperature") == 0)
    {
        Serial.println(dht.readTemperature());
        return (String)dht.readTemperature();
    }
    else if (strcmp(sensor, "humidity") == 0)
    {
        return (String)dht.readHumidity();
    }
    else if (strcmp(sensor, "digital") == 0)
    {
        toggleDigital(3, 1000);
        return "success";
    }
    else if (strcmp(sensor, "buzzer") == 0)
    {
        toggleDigital(PIN_BUZZER, 1000);
        return "success";
    }
    else if (strcmp(sensor, "ultrasonic_ranger") == 0)
    {
        return (String)ultrasonic.MeasureInCentimeters();
    }
}

void toggleDigital(int pin, int delayTime)
{
    digitalWrite(pin, HIGH);
    delay(delayTime);
    digitalWrite(pin, LOW);
}
