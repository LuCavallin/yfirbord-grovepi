package readers

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}
