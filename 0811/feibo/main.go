package feibo

//1 1
//2 1+1 2 2
//3 1+1+1 1+2 2+1 3
//4 1+1+1+1 1+1+2 1+2+1 2+1+1 2+2  5
//5 1+1+1+1+1 1+1+1+2 1+1+2+1 1+2+1+1 2+1+1+1 1+2+2 2+1+2 2+2+1       8

func fei(n int) int {
	if n <= 2 {
		return n
	}
	pre := 1
	cur := 2
	sum := 0
	for i := 3; i <= n; i++ {
		sum = pre + cur
		pre = cur
		cur = sum
	}
	return cur
}
