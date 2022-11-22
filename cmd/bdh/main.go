package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/skorens/bdh/internal/numbers"
	"github.com/urfave/cli/v2"
)

type BDH struct {
	ns       numbers.NumberSystem
	flags    numbers.ConversionBits
	value    string
	valueInt int
}

var (
	bdh = &cli.App{
		Name:                   "bdh",
		Version:                "1.0",
		Usage:                  "convert binary, decimal or hexadecimal values",
		UseShortOptionHandling: true,
		HideHelpCommand:        true,
		Before:                 setup,
		Action:                 convert,
		ArgsUsage:              "value to convert",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     "binary",
				Usage:    "convert a binary value",
				Aliases:  []string{"bin", "b"},
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "decimal",
				Usage:    "convert a decimal value",
				Aliases:  []string{"dec", "d"},
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "hexadecimal",
				Usage:    "convert a hexadecimal value",
				Aliases:  []string{"hex", "x"},
				Required: false,
			},
		},
	}
)

func setup(c *cli.Context) error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s\n", c.App.Version)
	}

	if c.NArg() < 1 {
		return fmt.Errorf("missing value to convert")
	}

	return nil
}

func main() {
	if err := bdh.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func convert(c *cli.Context) error {
	b := &BDH{
		ns:    numbers.UnknownNumberSystem,
		value: strings.ToLower(c.Args().Get(0)),
	}

	err := b.detect(c)
	if err != nil {
		return err
	}

	b.valueInt, err = numbers.Convert(b.ns, b.value)

	if err != nil {
		return err
	}

	b.print()
	return nil
}

func (b *BDH) detect(c *cli.Context) error {
	// detect
	if strings.HasPrefix(b.value, "0x") {
		b.ns = numbers.HexadecimalNumberSystem
		b.flags = b.flags.Set(numbers.NotBinary | numbers.NotDecimal)
		b.value = strings.TrimPrefix(b.value, "0x")
	} else if strings.HasPrefix(b.value, "0b") {
		b.ns = numbers.BinaryNumberSystem
		b.flags = b.flags.Set(numbers.NotHexadecimal | numbers.NotDecimal)
		b.value = strings.TrimPrefix(b.value, "0b")
	}

	// confirm flag matches possible values
	if c.Bool("decimal") && (b.ns != numbers.UnknownNumberSystem && b.ns != numbers.DecimalNumberSystem) {
		return fmt.Errorf("specified -d but value is not a decimal")
	} else if c.Bool("hexadecimal") && (b.ns != numbers.UnknownNumberSystem && b.ns != numbers.HexadecimalNumberSystem) {
		return fmt.Errorf("specified -h but value is not a hexadecimal")
	} else if c.Bool("binary") && (b.ns != numbers.UnknownNumberSystem && b.ns != numbers.BinaryNumberSystem) {
		return fmt.Errorf("specified -b but value is not a binary")
	}

	// check flags
	if c.Bool("decimal") {
		b.ns = numbers.DecimalNumberSystem
	} else if c.Bool("binary") {
		b.ns = numbers.BinaryNumberSystem
	} else if c.Bool("hexadecimal") {
		b.ns = numbers.HexadecimalNumberSystem
	}

	// validate
	if err := b.validate(b.value); err != nil {
		return err
	}

	// determine
	if b.ns == numbers.UnknownNumberSystem {
		if b.flags.Has(numbers.NotHexadecimal) && b.flags.Has(numbers.NotDecimal) && b.flags.Has(numbers.ValidBinary) {
			b.ns = numbers.BinaryNumberSystem
		} else if b.flags.Has(numbers.NotBinary) && b.flags.Has(numbers.NotDecimal) && b.flags.Has(numbers.ValidHexadecimal) {
			b.ns = numbers.HexadecimalNumberSystem
		} else if b.flags.Has(numbers.NotBinary) && b.flags.Has(numbers.NotHexadecimal) && b.flags.Has(numbers.ValidDecimal) {
			b.ns = numbers.DecimalNumberSystem
		}
	}

	// guess
	if b.ns == numbers.UnknownNumberSystem {
		if b.flags.Has(numbers.ValidBinary) {
			b.ns = numbers.BinaryNumberSystem
		} else if b.flags.Has(numbers.ValidDecimal) {
			b.ns = numbers.DecimalNumberSystem
		} else if b.flags.Has(numbers.ValidHexadecimal) {
			b.ns = numbers.HexadecimalNumberSystem
		} else {
			return fmt.Errorf("input not valid")
		}
	}

	return nil
}

func (b *BDH) print() {
	// always print in the same order for ease of parsing separated by new lines
	// type
	// bin
	// dec
	// hex
	fmt.Printf("%v\n", b.ns)
	fmt.Println(numbers.BinaryString(b.valueInt))
	fmt.Println(numbers.DecimalString(b.valueInt))
	fmt.Println(numbers.HexadecimalString(b.valueInt))
}

func (b *BDH) validate(value string) error {
	for _, d := range value {
		if d != '0' && d != '1' {
			b.flags = b.flags.Set(numbers.NotBinary)
		}
		if (d >= 'G' && d <= 'Z') || (d >= 'g' && d <= 'z') {
			b.flags = b.flags.Set(numbers.NotHexadecimal | numbers.NotDecimal)
			return fmt.Errorf("not a valid number to convert")
		}
		if (d >= 'A' && d <= 'F') || (d >= 'a' && d <= 'f') {
			b.flags = b.flags.Set(numbers.NotDecimal | numbers.NotBinary)
		}
	}

	if !b.flags.Has(numbers.NotBinary) {
		b.flags = b.flags.Set(numbers.ValidBinary)
	}

	if !b.flags.Has(numbers.NotHexadecimal) {
		b.flags = b.flags.Set(numbers.ValidHexadecimal)
	}

	if !b.flags.Has(numbers.NotDecimal) {
		b.flags = b.flags.Set(numbers.ValidDecimal)
	}

	return nil
}
