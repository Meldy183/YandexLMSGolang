package main

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (stu *Student) IsExcellentStudent() bool {
	return float64(stu.solvedProblems)*stu.scoreForOneTask > stu.passingScore
}
