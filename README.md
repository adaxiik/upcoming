# upcoming
Tool with which you can save and print your upcoming appointments, to get stressed even more.


## Build
```bash
$ go build -o upcoming
```

## Install
```bash
$ go install
```
- Note: You may need to add go/bin to your PATH

## Usage
```bash
$ upcoming -h
Usage of upcoming:
  -filepath string
        filepath (always required)
  -mode string
        add, remove, print (default "print")
  -name string
        name to add (required in add and remove mode)
  -time string
        time to add (required in add mode)
```

## Supported Time Formats
#### Smart formats (automatically set to current day/year)
- dd.mm
- d.m
- hh:mm
- hh:mm:ss
#### Full formats
- yyyy-mm-dd hh:mm:ss
- yyyy-mm-dd hh:mm
- yyyy-mm-dd
- dd.mm.yyyy hh:mm:ss
- dd.mm.yyyy hh:mm
- dd.mm.yyyy
- dd.mm.yy hh:mm:ss
- dd.mm.yy hh:mm
- dd.mm.yy
- d.m.yyyy hh:mm:ss
- d.m.yyyy hh:mm
- d.m.yyyy
- d.m.y hh:mm:ss
- d.m.y hh:mm
- d.m.y
#### Relative formats
- Relative formats must start with `+`
- Supported units: `h`, `m`, `s`
- Examples:
  - `+1h` -> 1 hour from now
  - `+1h30m` -> 1 hour and 30 minutes from now
  - `+1h30m10s` -> 1 hour, 30 minutes and 10 seconds from now

## Examples
```bash
$ upcoming -filepath /path/to/file -mode add -name "Appointment" -time +12h
$ upcoming -filepath /path/to/file -mode add -name "Party" -time 2023-05-07
$ upcoming -filepath /path/to/file -mode add -name "Meeting" -time 12.3.2023 12:00
$ upcoming -filepath /path/to/file -mode print
Party: 79d 6h 0m (1902h0m31s)
Meeting: 23d 6h 0m (558h0m31s)
Appointment: 12h 0m (12h0m0s)
$ upcoming -filepath /path/to/file -mode remove -name "Party"
$ upcoming -filepath /path/to/file
Meeting: 23d 5h 59m (557h59m20s)
Appointment: 11h 58m (11h58m49s)
```
- `/path/to/file` is json file, which will be created if it does not exist.
- All appointments that are in the past will be removed automatically.

## Expected usage
- `upcoming -filepath /path/to/file | head -n 1` in your `.bashrc` or `.zshrc` to get stressed every time you open a new terminal