package kokoEating

// solution 0
// runtime: 48.2 percentile
// memory : 18.11 percentile
func minEatingSpeed(piles []int, h int) int {
	k_min, k_max := max(1, len(piles)/h), 0
	for _, p := range piles {
		k_max = max(k_max, p)
	}
	var speed, hoursTaken int
	for k_min <= k_max {
		speed = k_min + (k_max-k_min)/2
		hoursTaken = 0
		// Check if speed is acceptable
		for _, pile := range piles {
			// same as += ceil(pile / speed)
			hoursTaken += (pile + (speed - 1)) / speed
		}
		if hoursTaken > h {
			// too slow
			k_min = speed + 1
		} else {
			// too fast
			k_max = speed - 1
		}
	}
	return k_min
}

// solution 1
// runtime: 83.08 percentile (Seems variable?)
// memory : 63.32 percentile
func minEatingSpeed(piles []int, h int) int {
	k_min, k_max := 1, 0
	for _, p := range piles {
		k_max = max(k_max, p)
	}
	if len(piles) == h {
		return k_max
	}

	for k_min <= k_max {
		speed := k_min + (k_max-k_min)/2
		hoursTaken := 0
		// Check if speed is acceptable
		for _, pile := range piles {
			// same as += ceil(pile / speed)
			hoursTaken += (pile + (speed - 1)) / speed
		}
		if hoursTaken > h {
			// too slow
			k_min = speed + 1
		} else {
			// too fast
			k_max = speed - 1
		}
	}
	return k_min
}
