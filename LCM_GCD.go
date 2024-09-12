package main
import "fmt"

func lcm(t1 int, t2 int) {
	var lcmnum int = 1
	if t1 > t2 {
		lcmnum = t1
	} else {
		lcmnum = t2
	}
	for {
		if lcmnum%t1 == 0 && lcmnum%t2 == 0 {
			fmt.Printf("LCM of %d and %d is %d", t1, t2, lcmnum)
			break
		}
		lcmnum++
	}
	return
}

func gcd(t1 int, t2 int) {
	var gcdnum int
	for i := 1; i <= t1 && i <= t2; i++ {
		if t1%i == 0 && gcdnum%i == 0 {
			gcdnum = i
		}
	}
	fmt.Print("gcd is:", gcdnum)
	return
}

func main() {
	var n1, n2, action int
	fmt.Println("enter 2 positive integers:")
	fmt.Scanln(&n1)
	fmt.Scanln(&n2)
	fmt.Println("enter 1 for lcm and 2 for gcd")
	fmt.Scanln(&action)
	switch action {
	case 1:
		lcm(n1, n2)
	case 2:
		gcd(n1, n2)
	}
}
