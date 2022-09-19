package main

type People struct {
	Number int
	Next   *People
}

//AddPeople 添加一个人
func AddPeople(i int) *People {
	return &People{
		Number: i,
	}
}

//AddMore 多个人组成环
func AddMore(n int) *People {

	var nextPeople *People
	people := AddPeople(0)
	current := people
	for i := 1; i < n; i++ {
		nextPeople = AddPeople(i)
		current.Next = nextPeople
		current = nextPeople
	}
	nextPeople.Next = people
	return people
}

//n个人，数到m的人退出
//返回结果 退出人切片，保留人切片
func game(n, m int) ([]int, []int) {
	if n < 1 {
		return nil, nil
	}
	if n == 1 {
		return nil, []int{0}
	}

	people := AddMore(n)
	exitList := make([]int, 0, n-1)
	leaveList := make([]int, 0, 1)
	count := 1
	leave := n

	for {
		if leave == 1 {
			leaveList = append(leaveList, people.Number)
			break
		}

		if count == m-1 {
			exitList = append(exitList, people.Next.Number)

			people.Next = people.Next.Next
			count = 0
			leave = leave - 1
		}
		count++
		people = people.Next
	}

	return exitList, leaveList
}
