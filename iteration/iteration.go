package iterator

const repeatCounter = 5

func Repeat(character string) string {
	return RepeatTimes(character, repeatCounter)
}

func RepeatTimes(character string, repeatTimes int) string {
	var repeated string
	for i := 0; i < repeatTimes; i++ {
		repeated += character
	}

	return repeated
}
