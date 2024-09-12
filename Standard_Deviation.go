package main
import (
	"fmt"
	"math"
)

func main() {
	var num [5]float64
	var sum, mean, sd float64
	fmt.Println("Enter 5 elements")
	for i := 1; i <= 5; i++ {
		fmt.Printf("Enter %d element: ", i)
		fmt.Scan(&num[i-1])
		sum += num[i-1]
	}
	mean = sum / 5
	for j := 0; j < 5; j++ {
		sd += math.Pow(num[j]-mean, 2)
	}
	sd = math.Sqrt(sd / 5)
	fmt.Println("The Standard Deviation is: ", sd)
}
