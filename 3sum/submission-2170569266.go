var result [][]int

func init() {
	result = make([][]int, 18000)
	for i := range result {
		result[i] = make([]int, 3)
	}
}

func abs(x int) (int, int) {
	if x < 0 {
		return -x, 1
	}
	return x, 0
}

func threeSum(nums []int) [][]int {
	var freq [100_001][2]uint8
	r, minv, maxv := 0, nums[0], nums[0]
	add := func(a, b, c int) {
		result[r][0], result[r][1], result[r][2] = a, b, c
		r++
	}
	for _, v := range nums {
		minv = min(minv, v)
		maxv = max(maxv, v)
		x, sign := abs(v)
		if freq[x][sign] < 3 {
			freq[x][sign]++
		}
	}
	// case (0, 0, 0)
	if freq[0][0] > 2 {
		add(0, 0, 0)
	}
	// early exit
	if minv >= 0 || maxv <= 0 {
		return result[:r]
	}
	// normalize minv and shrink search space
	minv = min(-minv, maxv*2)
	maxv = min(maxv, minv*2)
	// collect unique positive and negative values
	pos := make([]int, 0, len(nums))
	for v := 1; v <= maxv; v++ {
		if freq[v][0] > 0 {
			pos = append(pos, v)
		}
	}
	neg := make([]int, 0, len(nums))
	for v := 1; v <= minv; v++ {
		if freq[v][1] > 0 {
			neg = append(neg, v)
		}
	}
	// case (-a, 0, +a)
	if freq[0][0] > 0 {
		for _, a := range neg {
			if freq[a][0] > 0 {
				add(-a, 0, a)
			}
		}
	}
	for i := 0; i < len(neg) && neg[i]*2 <= maxv; i++ {
		// case (-a, -a, +2a)
		a := neg[i]
		if c := 2 * a; freq[a][1] > 1 && freq[c][0] > 0 {
			add(-a, -a, c)
		}
		// case (-b, -a, +a+b)
		for j := i + 1; j < len(neg) && neg[j]+a <= maxv; j++ {
			b := neg[j]
			if c := a + b; freq[c][0] > 0 {
				add(-b, -a, c)
			}
		}
	}
	for i := 0; i < len(pos) && pos[i]*2 <= minv; i++ {
		// case (-2a, +a, +a)
		a := pos[i]
		if c := 2 * a; freq[a][0] > 1 && freq[c][1] > 0 {
			add(-c, a, a)
		}
		// case (-a-b, +a, +b)
		for j := i + 1; j < len(pos) && pos[j]+a <= minv; j++ {
			b := pos[j]
			if c := a + b; freq[c][1] > 0 {
				add(-c, a, b)
			}
		}
	}
	return result[:r]
}