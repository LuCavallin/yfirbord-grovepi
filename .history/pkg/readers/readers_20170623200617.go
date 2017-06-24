package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(type string) (Reader, error) {

}