#include <stdio.h>
#include <wiringPi.h>

void setup()
{
}

void loop()
{
    printf("Hytta.\n");
    digitalWrite(0, HIGH);
    delay(500);
    digitalWrite(0, LOW);
    delay(500);
}

int main(int argc, char *argv[])
{
    wiringPiSetup();
    pinMode(16, OUTPUT);
    while (true)
    {
        loop();
    }
    return 0;
}
