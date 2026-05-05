package strings

const (
	zero     = 0
	ten      = 10
	hundred  = 100
	thousand = 1000
	million  = 1000000
	billion  = 1000000000
)

var digits = map[int]string{
	1:        "One",
	2:        "Two",
	3:        "Three",
	4:        "Four",
	5:        "Five",
	6:        "Six",
	7:        "Seven",
	8:        "Eight",
	9:        "Nine",
	ten:      "Ten",
	11:       "Eleven",
	12:       "Twelve",
	13:       "Thirteen",
	14:       "Fourteen",
	15:       "Fifteen",
	16:       "Sixteen",
	17:       "Seventeen",
	18:       "Eighteen",
	19:       "Nineteen",
	20:       "Twenty",
	30:       "Thirty",
	40:       "Forty",
	50:       "Fifty",
	60:       "Sixty",
	70:       "Seventy",
	80:       "Eighty",
	90:       "Ninety",
	hundred:  "Hundred",
	thousand: "Thousand",
	million:  "Million",
	billion:  "Billion",
}

// NumberInEnglish solves the problem in O(1) time and O(1) space.
func NumberInEnglish(num int) string {
	panic("not implemented")

}

func outputIfLarger(num, unit int, word string) (int, string) {
	panic("not implemented")

}

func threeDigitWord(num int) string {
	panic("not implemented")

}

func howMany(num, level int) int {
	panic("not implemented")

}
