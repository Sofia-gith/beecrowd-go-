package beecrowd


import (
	"fmt"
	"sort"
	
)



func RoleBad(){
var n, m int
	fmt.Scan(&n, &m)

	problems := make([]string, 0, n*m)
	for i := 0; i < n*m; i++ {
		var p string
		fmt.Scan(&p)
		problems = append(problems, p)
	}

	sort.Slice(problems, func(i, j int) bool {
		ti := typeOrder(problems[i])
		tj := typeOrder(problems[j])
		if ti != tj {
			return ti < tj
		}
		ci := int(problems[i][0] - '0')
		cj := int(problems[j][0] - '0')
		return ci > cj // decrescente
	})

	for _, p := range problems {
		fmt.Println(p)
	}
}

func typeOrder(p string) int {
	if p[1] == 'V' {
		return 0
	}
	return 1
}