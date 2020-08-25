package main

func add(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	return a + b
}

func substract(a, b int) int {
	if a == 0 {
		return -b
	}
	if b == 0 {
		return a
	}
	if b < 0 {
		return add(a, -b)
	}
	return a - b
}
