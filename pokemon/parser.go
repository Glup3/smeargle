package pokemon

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseGenerationString(s string) ([]int, error) {
	if s == "" {
		return []int{}, nil
	}

	var result []int
	seen := NewSet[int]()
	parts := strings.Split(s, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			// Handle range
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid start number: %s", rangeParts[0])
			}

			if start <= 0 || start > 8 {
				return nil, fmt.Errorf("start number has to be between 1 and 8: %d", start)
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid end number: %s", rangeParts[1])
			}

			if end <= 0 || end > 8 {
				return nil, fmt.Errorf("end number has to be between 1 and 8: %d", end)
			}

			if start > end {
				return nil, fmt.Errorf("start of range is greater than end: %s", part)
			}

			for i := start; i <= end; i++ {
				if !seen.Has(i) {
					result = append(result, i)
					seen.Add(i)
				}
			}
		} else {
			// Handle individual number
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", part)
			}

			if num <= 0 || num > 8 {
				return nil, fmt.Errorf("number has to be between 1 and 8: %d", num)
			}

			if !seen.Has(num) {
				result = append(result, num)
				seen.Add(num)
			}
		}
	}

	return result, nil
}

func ParseOrderByString(s string) (OrderBy, error) {
	switch strings.ToLower(s) {
	case "alphabet":
		return Alphabet, nil
	case "idx":
		return Idx, nil
	default:
		return 0, fmt.Errorf("invalid order by: %s", s)
	}
}

func ParseSortDirectionString(s string) (SortDirection, error) {
	switch strings.ToLower(s) {
	case "asc":
		return Asc, nil
	case "desc":
		return Desc, nil
	default:
		return 0, fmt.Errorf("invalid sort direction: %s", s)
	}
}
