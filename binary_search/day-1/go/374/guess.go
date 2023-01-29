package guess_number

/** problem 374
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
var Target int

func guess(num int) int {
	if num > Target {
		return -1
	} else if num < Target {
		return 1
	}
	return 0
}
func GuessNumber(n int) int {
	l, r := 0, n
	for l <= r {
		mid := l + (r-l)/2
		res := guess(mid)
		if res == -1 {
			r = mid - 1
		} else if res == 1 {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}
