#include "PietteTech_DHT.h"

#define DHTTYPE DHT11 // Sensor type DHT11/21/22/AM2301/AM2302
#define DHTPIN D4     // Digital pin for communications

// Lib instantiate
PietteTech_DHT DHT(DHTPIN, DHTTYPE);
int n; // counter

void setup()
{
  Serial.begin(9600);
  while (!Serial.available() && millis() < 30000)
  {
    Serial.println("Press any key to start.");
    Particle.process();
    delay(1000);
  }
  Serial.println("DHT Simple program using DHT.acquireAndWait");
  Serial.print("LIB version: ");
  Serial.println(DHTLIB_VERSION);
  Serial.println("---------------");
  DHT.begin();
}

void loop()
{
  Serial.print("\n");
  Serial.print(n);
  Serial.print(": Retrieving information from sensor: ");
  Serial.print("Read sensor: ");

  int result = DHT.acquireAndWait(1000); // wait up to 1 sec (default indefinitely)

  switch (result)
  {
  case DHTLIB_OK:
    Serial.println("OK");
    break;
  case DHTLIB_ERROR_CHECKSUM:
    Serial.println("Error\n\r\tChecksum error");
    break;
  case DHTLIB_ERROR_ISR_TIMEOUT:
    Serial.println("Error\n\r\tISR time out error");
    break;
  case DHTLIB_ERROR_RESPONSE_TIMEOUT:
    Serial.println("Error\n\r\tResponse time out error");
    break;
  case DHTLIB_ERROR_DATA_TIMEOUT:
    Serial.println("Error\n\r\tData time out error");
    break;
  case DHTLIB_ERROR_ACQUIRING:
    Serial.println("Error\n\r\tAcquiring");
    break;
  case DHTLIB_ERROR_DELTA:
    Serial.println("Error\n\r\tDelta time to small");
    break;
  case DHTLIB_ERROR_NOTSTARTED:
    Serial.println("Error\n\r\tNot started");
    break;
  default:
    Serial.println("Unknown error");
    break;
  }
  Serial.print("Humidity (%): ");
  Serial.println(DHT.getHumidity(), 2);

  Serial.print("Temperature (oC): ");
  Serial.println(DHT.getCelsius(), 2);

  Serial.print("Temperature (oF): ");
  Serial.println(DHT.getFahrenheit(), 2);

  Serial.print("Temperature (K): ");
  Serial.println(DHT.getKelvin(), 2);

  Serial.print("Dew Point (oC): ");
  Serial.println(DHT.getDewPoint());

  Serial.print("Dew Point Slow (oC): ");
  Serial.println(DHT.getDewPointSlow());

  n++;
  delay(2500);
}
