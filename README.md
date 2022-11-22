# bdh

A tool to convert binary, decimal or hexadecimal values.

```bash
$ bdh -h        
NAME:
   bdh - convert binary, decimal or hexadecimal values

USAGE:
   bdh [global options] value to convert

VERSION:
   1.0

GLOBAL OPTIONS:
   --binary, --bin, -b       convert a binary value (default: false)
   --decimal, --dec, -d      convert a decimal value (default: false)
   --hexadecimal, --hex, -x  convert a hexadecimal value (default: false)
   --help, -h                show help (default: false)
   --version, -v             print the version (default: false)
```

## Build & Test
```bash
# run
go run ./cmd/bdh

# build
go build -o bdh ./cmd/bdh

# test
go test ./...
```

## Output

The ouptut is always the same. Four fields are returned, one per new line
```bash
# type (bin, dec, hex)
# value converted to binary
# value converted to decimal
# value converted to hexadecimal
```

```bash
$ bdh 42        
dec
0b101010
0x2a
42
```

## Examples

```bash
# use -b to specify a binary
$ bdh -b 1010101
bin
0b1010101
0x55
85
```

```bash
# use -d to specify a decimal
$ bdh -d 59
dec
0b111011
0x3b
59
```

```bash
# use -x to specify a hexadecimal
$ bdh -x 170
hex
0b101110000
0x170
368
```

```bash
# detect binary when 0b prefix is present
$ bdh 0b111011
bin
0b111011
0x3b
59
```

```bash
# detect hexadecimal when 0x prefix is present
$ bdh 0x170   
hex
0b101110000
0x170
368
```

```bash
# guess bin, then hex and finally dec
$ bdh 11001   
bin
0b11001
0x19
25

$ bdh abdef
hex
0b10101011110111101111
0xabdef
703983

$ bdh 123  
dec
0b1111011
0x7b
123
```

