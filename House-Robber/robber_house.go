package House_Robber

func robbing(number []int) int {
	if len(number) == 0 {
		return 0
	}
	rh:= make([]int,len(number))
	rh[0] = number[0]
	max := rh[0]
	for i := 1; i < len(number); i ++{
		rh[i] = number[i]
		for j := 0; j <= i - 2; j ++ {
			if rh[j] + number[i] > rh[i] {
				rh[i] = rh[j] + number[i]
			}
		}
		if rh[i]>max{
   		max=rh[i]
		}
	}
	return max
}

func maxvalue(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Rob(num []int) int {
	if len(num) == 0 {
		return 0
	}
	if len(num) == 1 {
		return num[0]
	}
	return maxvalue(robbing(num[1 :]), robbing(num[0 : len(num) - 1]))
}