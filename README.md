# Advent of Code 2024 (Go)

Advent of Code 2024 repo 🎄 Solutions in Go

Solutions to [Advent of Code 2024](https://adventofcode.com/2024) challenges, a yearly coding event where a new puzzle is released each day from December 1st to December 25th.

### Automatically Download Input
To automatically download your input, create a `.env` file in the `input` directory, add your session cookie and then run get_input.go with the day flag set to the input you'd like to grab:

```bash
mkdir input/.env
echo 'SESSION_COOKIE="YOUR SESSION COOKIE"' > input/.env
cd input
go run input/get_input.go --day 1
```

