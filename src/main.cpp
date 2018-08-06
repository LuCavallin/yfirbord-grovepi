#include <Bridge.h>
#include <BridgeClient.h>
#include <ArduinoJson.h>
#include <MQTT.h>
#include <Time.h>

// Networking
BridgeClient connection;
MQTTClient client;

// Sensors and actuators libraries
#include "Ultrasonic.h"
#include "DHT.h"
const byte SENSOR_READ = 0;
const byte SENSOR_WRITE = 1;
const byte PIN_ULTRASONIC = 2;
const byte PIN_LED = 3;
const byte PIN_DHT = 4;
const byte PIN_BUZZER = 8;
DHT dht(PIN_DHT, DHT11);
Ultrasonic ultrasonic(PIN_ULTRASONIC);

// Setup MQTT
const String INCOMING_TOPIC = "to-arduino";
const String OUTGOING_TOPIC = "from-arduino";

String handleRequest(String sensor, char *params);
void toggleDigital(int pin, int delayTime);

void onMessageReceived(String &topic, String &payload)
{
    // Debug message
    Serial.println(sprintf("Message received from topic: %s", topic.c_str()));

    // Get and parse request
    String json = payload;
    StaticJsonBuffer<sizeof(payload)> jsonBuffer;
    JsonObject &incomingRoot = jsonBuffer.parseObject(json);

    // Test if parsing succeeds.
    if (!incomingRoot.success())
    {
        Serial.println("Failed to parse request.");
        return;
    }

    // Prepare response and add metadata
    String sensor = incomingRoot["sensor"];
    JsonObject &outgoingRoot = jsonBuffer.createObject();
    // @TODO find way to have a unique ID for the device
    outgoingRoot["id"] = "1394y71289312";
    outgoingRoot["timestamp"] = now();
    outgoingRoot["sensor"] = sensor;
    outgoingRoot["data"] = handleRequest(sensor, incomingRoot["params"]);

    // Converting response to needed type
    String response;
    outgoingRoot.printTo(response);

    // Send response, debug message
    Serial.println(sprintf("Publishing message: %s to topic: %s.", sensor.c_str(), OUTGOING_TOPIC.c_str()));
    client.publish(OUTGOING_TOPIC, response);
}

// Program
void connect()
{
    Serial.println("Connecting to MQTT...");
    while (!client.connect("hytta", MQTT_USERNAME, MQTT_PASSWORD))
    {
        Serial.print(".");
        delay(1000);
    }
    Serial.println("Connected!");

    client.subscribe(INCOMING_TOPIC);
}

void setup()
{
    // Setup coms
    Serial.println("Setting up Bridge and Serial...");
    Bridge.begin();
    Serial.begin(115200);
    client.begin(MQTT_HOSTNAME, MQTT_PORT, connection);
    client.onMessage(onMessageReceived);

    // Setup Sensors
    dht.begin();
    pinMode(PIN_BUZZER, SENSOR_WRITE);
    pinMode(PIN_LED, SENSOR_WRITE);
}

void loop()
{
    client.loop();
    // MQTT
    if (!client.connected())
    {
        connect();
    }

    client.subscribe(INCOMING_TOPIC);
}

// This needs serious re-thought
String handleRequest(String sensor, char *params)
{
    if (sensor == "temperature")
    {
        Serial.println(dht.readTemperature());
        return (String)dht.readTemperature();
    }
    else if (sensor == "humidity")
    {
        return (String)dht.readHumidity();
    }
    else if (sensor == "digital")
    {
        toggleDigital(3, 1000);
        return "success";
    }
    else if (sensor == "buzzer")
    {
        toggleDigital(PIN_BUZZER, 1000);
        return "success";
    }
    else if (sensor == "ultrasonic_ranger")
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
