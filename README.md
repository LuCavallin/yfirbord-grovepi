# Yfirbord GrovePi

[![Build Status](https://travis-ci.org/LuCavallin/yfirbord-grovepi.svg?branch=master)](https://travis-ci.org/LuCavallin/yfirbord-grovepi)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/719fb2f2ff4d46c7841611b970fa2c15)](https://www.codacy.com/app/lucavallin/yfirbord-grovepi?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=LuCavallin/yfirbord-grovepi&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/LuCavallin/yfirbord-grovepi/branch/master/graph/badge.svg)](https://codecov.io/gh/LuCavallin/yfirbord-grovepi)



Yfirbord GrovePi, the open source environmental monitor.


## Usage

Make sure you install the dependencies via glide, then build and copy the tool to the RaspberryPi with:

```sh
$ make build copy
```

If this doesn't work for you, check the parameters in the Makefile.

## Crontab

To make it work properly configure add a crontab entry to run yfirbord every amount-of-time-you-prefer.