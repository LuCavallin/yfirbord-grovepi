#include <wiringPi.h>
int main(void)
{
    wiringPiSetup();
    pinMode(16, OUTPUT);
    for (;;)
    {
        digitalWrite(16, HIGH);
        delay(500);
        digitalWrite(16, LOW);
        delay(500);
    }
    return 0;
}
