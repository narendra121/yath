package maths

func DigiCount(x int) int {
	if x < 0 {
		return 0
	}
	var count int
	for x > 0 {
		x = x / 10
		count++
	}
	return count
}
func IsPalindrome(num int) bool {
	var pd int
	cp := num
	for num > 0 {
		temp := num % 10
		num = num / 10
		pd = pd*10 + temp
	}
	return pd == cp
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if i != num && num%i == 0 {
			return false
		}
	}
	return true
}

//weather kth bit is 1 or not
func IsBitChanged(num, position int) bool {
	temp := num >> (position - 1)
	return (temp & 1) != 0
}

/*check is val can be some power of num
 	val = 4 num =2  --> true-->2^2=4
	val= 6  num=2   --> false

*/
func IsPowerOf(val, num int) bool {
	for val > 1 {
		if val%num == 1 {
			return false
		}
		val /= num
	}

	return true
}
