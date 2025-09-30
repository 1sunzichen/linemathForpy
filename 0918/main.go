package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 文件内容排序

// 针对文件内容进行排序，
// 先按照每行的第一列按字典序升序，
// 遇到相同的按照第二列按数值降序

// Demo：

// songj123 100
// sj 23
// sj 445
// ritu 56

// 转化后：

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
			return rows[i].col2 > rows[j].col2 // 第二列数值降序
		}
		return rows[i].col1 < rows[j].col1 // 第一列字典序升序
	})

	for _, r := range rows {
		fmt.Printf("%s %d\n", r.col1, r.col2)
	}
}
