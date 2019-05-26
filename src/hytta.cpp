#include "application.h"
#line 1 "/Users/luca/Projects/hytta/src/hytta.ino"
#include "PietteTech_DHT.h"
#include <blynk.h>

// Setup Blynk serial output for debug prints and auth
// (get the auth token in the Blynk app project settings)
void setup();
void readDHT();
void loop();
#line 6 "/Users/luca/Projects/hytta/src/hytta.ino"
#define BLYNK_PRINT Serial
char blynkAuth[] = "74dab704d1824061bb43cf03df866244";
BlynkTimer timer;

// Setup DHT sensor DHT11/21/22/AM2301/AM2302 and digital pin (D0 cannot be used)
#define DHTTYPE DHT11
#define DHTPIN D4
PietteTech_DHT DHT(DHTPIN, DHTTYPE);

void setup()
{
  Serial.begin(9600);
  delay(5000);
  // Read sensor once per minute
  timer.setInterval(60000L, readDHT);
  Blynk.begin(blynkAuth);
  DHT.begin();
}

// Blynk currently handles the waiting time, if this wasn't the case it would be
// a good rule to wait ~2500ms between each read according to the PietteTech_DHT library
void readDHT()
{
  int result = DHT.acquireAndWait(1000);

  switch (result)
  {
  case DHTLIB_OK:
    Serial.println("OK");
    Particle.publish("status", "OK", PRIVATE);
    break;
  case DHTLIB_ERROR_CHECKSUM:
    Serial.println("Error\n\r\tChecksum error");
    Particle.publish("status", "Checksum error", PRIVATE);
    break;
  case DHTLIB_ERROR_ISR_TIMEOUT:
    Serial.println("Error\n\r\tISR time out error");
    Particle.publish("status", "ISR time out error", PRIVATE);
    break;
  case DHTLIB_ERROR_RESPONSE_TIMEOUT:
    Serial.println("Error\n\r\tResponse time out error");
    Particle.publish("status", "Response time out error", PRIVATE);
    break;
  case DHTLIB_ERROR_DATA_TIMEOUT:
    Serial.println("Error\n\r\tData time out error");
    Particle.publish("status", "Data time out error", PRIVATE);
    break;
  case DHTLIB_ERROR_ACQUIRING:
    Serial.println("Error\n\r\tAcquiring");
    Particle.publish("status", "Acquiring", PRIVATE);
    break;
  case DHTLIB_ERROR_DELTA:
    Serial.println("Error\n\r\tDelta time too small");
    Particle.publish("status", "Delta time too small", PRIVATE);
    break;
  case DHTLIB_ERROR_NOTSTARTED:
    Serial.println("Error\n\r\tNot started");
    Particle.publish("status", "Not started", PRIVATE);
    break;
  default:
    Serial.println("Unknown error");
    Particle.publish("status", "Unknown error", PRIVATE);
    break;
  }

  // Get temperature and humidity, then send some data to serial
  // and the Particle cloud for debugging
  float temperature = DHT.getCelsius();
  Serial.print("Temperature (oC): ");
  Serial.println(temperature, 2);
  Particle.publish("temperature", String(temperature), PRIVATE);

  float humidity = DHT.getHumidity();
  Serial.print("Humidity (%): ");
  Serial.println(humidity, 2);
  Particle.publish("humidity", String(humidity), PRIVATE);

  // Send data to the Blynk API, once for the monitor and once for the graph
  Blynk.virtualWrite(0, temperature);
  Blynk.virtualWrite(2, temperature);
  Blynk.virtualWrite(1, humidity);
  Blynk.virtualWrite(3, humidity);
}

void loop()
{
  Blynk.run();
  timer.run();
}
