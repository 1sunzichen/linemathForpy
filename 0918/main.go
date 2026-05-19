package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// File content sorting

// Sort file content,
// First sort by column 1 in lexicographic ascending order,
// For ties, sort by column 2 in numeric descending order

// Demo：

// songj123 100
// sj 23
// sj 445
// ritu 56

// After transformation:

// ritu 56
// sj 445
// sj 23
// songj123 100
type line struct {
	col1 string
	col2 int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var rows []line

	for sc.Scan() {
		s := strings.TrimSpace(sc.Text())
		if s == "" {
			continue
		}
		fields := strings.Fields(s)
		if len(fields) < 2 {
			continue
		}
		n, _ := strconv.Atoi(fields[1])
		rows = append(rows, line{col1: fields[0], col2: n})
	}

	sort.Slice(rows, func(i, j int) bool {
		if rows[i].col1 == rows[j].col1 {
			return rows[i].col2 > rows[j].col2 // Column 2 numeric descending
		}
		return rows[i].col1 < rows[j].col1 // Column 1 lexicographic ascending
	})

	for _, r := range rows {
		fmt.Printf("%s %d\n", r.col1, r.col2)
	}
}
