package numbers

type NumberSystem int

const (
	UnknownNumberSystem NumberSystem = iota
	BinaryNumberSystem
	DecimalNumberSystem
	HexadecimalNumberSystem
)

func (ns NumberSystem) String() string {
	switch ns {
	case BinaryNumberSystem:
		return "bin"
	case DecimalNumberSystem:
		return "dec"
	case HexadecimalNumberSystem:
		return "hex"
	case UnknownNumberSystem:
		fallthrough
	default:
		return "unknown"
	}
}
