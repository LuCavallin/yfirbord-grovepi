package messages

import (
	"testing"
	"time"
)

func TestNewReading(t *testing.T) {
	name := "sound"
	val := 200
	timestamp := time.Now()
	message := NewReading(name, val, timestamp)

	if message.Name != name {
		t.Errorf("message.name was incorrect, got: %s, want: %s.", message.Name, name)
	}

	if message.Value != val {
		t.Errorf("message.value was incorrect, got: %d, want: %d.", message.Value, val)
	}

	if message.Time != timestamp {
		t.Errorf("message.time was incorrect, got: %v, want: %v.", message.Time, timestamp)
	}
}