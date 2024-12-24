package utils

// BoolToFloat converts bool to float64
func BoolToFloat(b bool) float64 {
	if b {
		return 1.0
	}
	return 0.0
}

// BoolToInt converts bool to int64
func BoolToInt(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

// FloatToBool converts float64 to bool
func FloatToBool(f float64) bool {
	return f != 0.0
}

// IntToFloat covverts int64 to float64
func IntToFloat(i int64) float64 {
	return float64(i)
}

// IntToBool converts int64 to bool
func IntToBool(i int64) bool {
	return i != 0
}

// FloatToInt converts float64 to int64
func FloatToInt(f float64) int64 {
	return int64(f)
}
