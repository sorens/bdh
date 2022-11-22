package numbers

import "testing"

type testData struct {
	value int
	dec   string
	hex   string
	bin   string
}

var (
	data = []testData{
		{
			value: 255,
			dec:   "255",
			hex:   "0xff",
			bin:   "0b11111111",
		},
		{
			value: 368,
			dec:   "368",
			hex:   "0x170",
			bin:   "0b101110000",
		},
		{
			value: 540,
			dec:   "540",
			hex:   "0x21c",
			bin:   "0b1000011100",
		},
		{
			value: 3021,
			dec:   "3021",
			hex:   "0xbcd",
			bin:   "0b101111001101",
		},
		{
			value: 43981,
			dec:   "43981",
			hex:   "0xABCD",
			bin:   "0b1010101111001101",
		},
	}
)

func TestConvertBinary(t *testing.T) {
	for index, h := range data {
		actual, err := Convert(BinaryNumberSystem, h.bin)

		if err != nil {
			t.Fatalf("expected no error, received %v\n", err)
		}

		if actual != data[index].value {
			t.Fatalf("expected %v, received %v\n", data[index].value, actual)
		}
	}
}

func TestConvertDecimal(t *testing.T) {
	for index, h := range data {
		actual, err := Convert(DecimalNumberSystem, h.dec)

		if err != nil {
			t.Fatalf("expected no error, received %v\n", err)
		}

		if actual != data[index].value {
			t.Fatalf("expected %v, received %v\n", data[index].value, actual)
		}
	}
}

func TestConvertHexadecimal(t *testing.T) {
	for index, h := range data {
		actual, err := Convert(HexadecimalNumberSystem, h.hex)

		if err != nil {
			t.Fatalf("expected no error, received %v\n", err)
		}

		if actual != data[index].value {
			t.Fatalf("expected %v, received %v\n", data[index].value, actual)
		}
	}
}
