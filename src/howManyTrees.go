package src

func Solution(n, h int) int {
	if n < h {
		return 0
	}
	if n == 0 {
		if h == 0 {
			return 1
		} else {
			return 0
		}
	}
	if h == 0 {
		return Solution(n, 1)
	}
	ans := Solution(n, h+1)
	for i := 1; i <= n; i++ {
		ans += (Solution(i-1, h-1) - Solution(i-1, h)) * (Solution(n-i, 0) - Solution(n-i, h-1))
		ans += (Solution(i-1, 0) - Solution(i-1, h-1)) * (Solution(n-i, h-1) - Solution(n-i, h))
		ans += (Solution(i-1, h-1) - Solution(i-1, h)) * (Solution(n-i, h-1) - Solution(n-i, h))
	}
	return ans
}
