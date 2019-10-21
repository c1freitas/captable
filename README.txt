# Capitalization Table Exercise

Run `make` to get a list of supported make commands. 
If you do not have make installed you can run each command directly. 
Golang 1.13 is required (1.11+ should work but is untested)

## Building
In order to build the CLI program you should have golang 1.13 installed. 
You can run the command `make build` which will output the file in the `./build/` directory. 
The built executable will be called `captable`.
To build without make, `go build -o build/captable .`

## testing
The included unit tests can be run by running `make tests`
Without make, you can run the tests by `go test -cover -v ./...`

## Running: 
You can run the `./build/captable help` to get a list of arguments.

The CLI takes 2 flags, -d (date) and -f (file) arguments.
date: [optional] The date used when computing investments.
file: [required] The file to read when computing the summary cap table

To run the cli, `captable -f /path/to/file.csv -d 2019-01-01`
If you used the make build command, the path would be `./build/captable`

## Assumptions
- Assumes the supplied file is on the local filesystem and readable. 
- Assumes the file is well formed, only minimal validation occurs (number of fields, valid date, numbers for shares purchased and cash paid).
- Percentages of ownership are rounded to 2 decimal places and total ownership may not equal EXACTLY 100.00%.
- JSON output truncates trailing 0 for decimal places. Values are correct, but should be fixed for display. 
  If this is critical, the JSON marshaling function can be overridden.

- JSON output is rendered to stdout, next version should support output to file, but that was not in the requirements


