package utils

import "fmt"

const (
	KiB = 1 << (10 * (iota + 1)) // Kibibytes
	MiB = 1 << (10 * (iota + 1)) // Mibibytes
	GiB = 1 << (10 * (iota + 1)) // Gibibytes
	TiB = 1 << (10 * (iota + 1)) // Tebibytes
	PiB = 1 << (10 * (iota + 1)) // Pebibytes
	EiB = 1 << (10 * (iota + 1)) // Exbibytes
)

func GetSize(value int64) string {
	floatValue := float64(value)
	switch {
	case value < KiB:
		return fmt.Sprintf("%d", value)

	case value < MiB:
		return fmt.Sprintf("%.2f KiB", floatValue/KiB)

	case value < GiB:
		return fmt.Sprintf("%.2f MiB", floatValue/MiB)

	case value < TiB:
		return fmt.Sprintf("%.2f GiB", floatValue/GiB)

	case value < PiB:
		return fmt.Sprintf("%.2f TiB", floatValue/TiB)

	case value < EiB:
		return fmt.Sprintf("%.2f PiB", floatValue/PiB)

	default:
		return fmt.Sprintf("%.2f EiB", floatValue/EiB)
	}
}
