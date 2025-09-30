package main

func sort(num []int) {
	n := len(num)
	sortIndex := 0

	for sortIndex < n {
		flag := false
		for i := n - 1; i > sortIndex; i-- {
			if num[i] < num[i-1] {
				num[i], num[i-1] = num[i-1], num[i]
			}
			flag = true
		}
		if !flag {
			break
		}
		sortIndex++
	}
}
func main() {
	n := []int{2, 3, 1, 2, 5, 5, 55, 12, 13}
	sort(n)
	println(n)
}
