package sensors

// Sensor contains all info on sensor
// Pin should better be in the reader, but in our case it's totally attached to the sensor itself
type Sensor struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Pin         byte   `json:"pin"`
	Mode        string `json:"mode"`
	Type        string `json:"type"`
	Size        int    `json:"size"`
}
