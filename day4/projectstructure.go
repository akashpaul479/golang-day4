package day4

type Students2 struct {
	Name  string
	Age   int
	Marks int
}

func Average(marks []int) float64 {
	sum := 0
	for _, m := range marks {
		sum += m
	}
	return float64(sum) / float64(len(marks))
}
