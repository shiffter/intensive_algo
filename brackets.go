package main

import "fmt"

func main() {
	s := "{([{}])}{"
	fmt.Println(validate(s))
}

func push(n []string, val string) []string {
	n = append(n, val)
	return n
}

func pop(n []string) (string, []string) {
	if len(n) < 1 {
		fmt.Println("bad len stack")
		return "", nil
	}
	tmpV := n[len(n)-1]

	if len(n) > 0 {
		n = n[:len(n)-1]
	}
	return tmpV, n
}

func validate(str string) bool {

	m1, m2 := make(map[string]interface{}, 3), make(map[string]interface{}, 3)

	st := make([]string, 0)

	m1["["], m1["("], m1["{"] = nil, nil, nil
	m2["]"], m2[")"], m2["}"] = nil, nil, nil

	for _, v := range str {
		if _, ok := m1[string(v)]; ok {
			st = push(st, string(v))
		} else {
			var tmp string
			tmp, st = pop(st)
			fmt.Println("tmp", tmp, "v", string(v))
			if tmp == "{" && "}" == string(v) {
				continue
			} else if tmp == "(" && ")" == string(v) {
				continue
			} else if tmp == "[" && "]" == string(v) {
				continue
			} else {
				return false
			}
		}
	}
	if len(st) != 0 {
		return false
	}
	return true
}
