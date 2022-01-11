package main

func main() {
	// s := Stack{}
	// s.pushMin(12)
	// s.pushMin(10)
	// s.pushMin(20)
	// s.pushMin(40)
	// s.pushMin(5)

	// fmt.Println(s.min())

	// 10
	// 10, 5
	// 10, 6, 5
	// 10, 6, 5, 2
	// 10, 6, 5, 2, 1
	//

	s := Stack{}
	s.push(10)
	s.push(5)
	s.push(6)
	s.push(2)
	s.push(1)
	s.push(20)
	s.push(14)
	s.push(3)
	newS := s.sort()

	newS.display()
}
