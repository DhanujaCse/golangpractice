package coding1

type Student struct {
	Name    string
	Age     int
	Grades1 [3]int
}

func Calculateavg(s Student) int {
	var sum int

	for i := 0; i < len(s.Grades1); i++ {
		sum = sum + s.Grades1[i]

	}
	avg := sum / len(s.Grades1)
	return avg

}
