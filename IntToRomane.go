package main

func main() {

	intToRoman(110)

}

func intToRoman(num int) string {
	nums := []int{1100, 1000, 900, 600, 500, 400, 110, 100, 90, 60, 50, 40, 11, 10, 9, 6, 5, 4, 1}
	symbols := []string{"MC", "M", "CM", "DC", "D", "CD", "CX", "C", "XC", "LX", "L", "XL", "XI",
		"X", "IX", "VI", "V", "IV", "I"}
	res, i := "", 0
	for num != 0 {
		for num < nums[i] {
			i++

		}
		num -= nums[i]
		res += symbols[i]

		print("-")
		print(res)

	}
	return res
}
