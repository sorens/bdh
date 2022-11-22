package numbers

type ConversionBits uint8

const (
	NotSet ConversionBits = 1 << iota
	NotBinary
	ValidBinary
	NotHexadecimal
	ValidHexadecimal
	NotDecimal
	ValidDecimal
)

func (c ConversionBits) Set(flag ConversionBits) ConversionBits    { return c | flag }
func (c ConversionBits) Clear(flag ConversionBits) ConversionBits  { return c | flag }
func (c ConversionBits) Toggle(flag ConversionBits) ConversionBits { return c ^ flag }
func (c ConversionBits) Has(flag ConversionBits) bool              { return c&flag != 0 }

func (c ConversionBits) String() string {
	result := "flags: "

	if c.Has(NotBinary) {
		result += "b"
	}
	if c.Has(NotHexadecimal) {
		result += "h"
	}
	if c.Has(NotDecimal) {
		result += "d"
	}
	if c.Has(ValidBinary) {
		result += "B"
	}
	if c.Has(ValidHexadecimal) {
		result += "H"
	}
	if c.Has(ValidDecimal) {
		result += "D"
	}

	return result
}
