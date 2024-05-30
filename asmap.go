package asmap

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// function that determines the separator of a line within file
func Separator(line string) string {

	//
	separators := []rune{',', '\t', ';', '|', ' '}
	separatorCounts := make(map[rune]int)

	//
	for _, sep := range separators {
		separatorCounts[sep] = strings.Count(line, string(sep))
	}

	//
	var s string
	maxCount := 0
	for sep, count := range separatorCounts {
		if count > maxCount {
			maxCount = count
			s = string(sep)
		}
	}

	return s
}

// modular function for reading csv
func Csv(f string) (map[string][]string, error) {

	//
	file, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer file.Close()

	//
	r := csv.NewReader(file)

	//
	headers, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	//
	result := make(map[string][]string)
	for _, header := range headers {
		result[header] = []string{}
	}

	//
	for {
		record, err := r.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, fmt.Errorf("error: %v", err)
		}
		for i, value := range record {
			result[headers[i]] = append(result[headers[i]], value)
		}
	}

	return result, nil

}

// modular function for reading all other file types and searches for common line sep
func Other(f string) (map[string][]string, error) {

	//
	file, err := os.Open(f)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	defer file.Close()

	//
	s := bufio.NewScanner(file)

	//
	s.Scan()
	hl := s.Text()

	//
	sep := Separator(hl)
	headers := strings.Split(hl, sep)

	//
	result := make(map[string][]string)
	for _, header := range headers {
		result[header] = []string{}
	}

	//
	for s.Scan() {

		cl := strings.Split(s.Text(), sep)
		if len(cl) != len(headers) {
			panic("Error")
		}
		//
		for i, value := range cl {
			result[headers[i]] = append(result[headers[i]], value)
		}

	}

	//
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return result, nil
}

// function
func ReadAsMap(f string) (map[string][]string, error) {

	//
	m := make(map[string][]string)

	// CSV
	if strings.HasSuffix(f, ".csv") {
		m, _ = Csv(f)
		_, err := Csv(f) // to catch error
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}

	} else { // other file types
		m, _ = Other(f)
		_, err := Other(f)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
	}

	return m, nil
}
