# Hytta Pi

[![Build Status](https://travis-ci.org/lucavallin/hytta-pi.svg?branch=master)](https://travis-ci.org/lucavallin/hytta-pi)
[![codecov](https://codecov.io/gh/lucavallin/hytta-pi/branch/master/graph/badge.svg)](https://codecov.io/gh/lucavallin/hytta-pi)



Hytta Pi, a cloud UI for your IoT devices.


## Usage

Make sure you install the dependencies via glide, then build and copy the tool to the RaspberryPi with:

```sh
$ make build copy
```

If this doesn't work for you, check the parameters in the Makefile.

## Crontab

To make it work properly configure add a crontab entry to run hytta every amount-of-time-you-prefer.
