package code

func reverse7(i int) int {
	tmp := 0
	for i != 0 {
		tmp = tmp*10 + i%10
		i = i / 10
	}
	if tmp < -1<<31 || tmp > 1<<31-1 {
		return 0
	}
	return tmp
}
