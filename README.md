# `asmap`

Lightweight method for loading in various filetypes as `map` in Go. 

## Details

Under the hood, `asmap.ReadAsMap()` assumes headers are the first line and searches for a common separator based on the headers, unless you wish to read a .csv, then the function uses `"encoding/csv"` to handle the file. Common separators include `',', '\t', ';', '|', ' '`. `asmap.ReadAsMap()` returns a map with the headers as keys and rows as values of string type-even if your data are `int` or `float64`.
>[!NOTE]
>To convert strings to integers or other numeric data types, you can use `strconv.ParseFloat()`.

## Installation
```bash
go get github.com/w-decker/asmap@latest
```

## Usage

```go
package main

import (
    "github.com/w-decker/asmap"
    "fmt"
)

func main() {
	m, err := asmap.ReadAsMap("data.txt")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
	fmt.Println(m)
}
```
```bash
go run main.go
```




