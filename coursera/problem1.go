package coursera

//doing this recurisive multiplication wise
//so for any given integer we'll need to first get the multiplicative value
//and then we'll multiple the next one by 10, and so forth

func multiply(a, b int64) int64 {
	var multiplied int64
	var i int64
	for i = 0; i < b; i++ {
		multiplied = multiplied + a
	}
	return multiplied
}

//Solve ... here'syourcommentlinter
func Solve(one, two int64) int64 {
	return multiply(one, two)
}
