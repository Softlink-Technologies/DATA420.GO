package main

import (
    "fmt"
    "math"
)

func main() {
    // Define an array of numbers
    nums := []float64{2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0}

    // Calculate the sum of the numbers
    sum := 0.0
    for _, num := range nums {
        sum += num
    }
    fmt.Printf("Sum: %.2f\n", sum)

    // Calculate the mean of the numbers
    mean := sum / float64(len(nums))
    fmt.Printf("Mean: %.2f\n", mean)

    // Calculate the variance of the numbers
    variance := 0.0
    for _, num := range nums {
        variance += math.Pow(num-mean, 2)
    }
    variance /= float64(len(nums))
    fmt.Printf("Variance: %.2f\n", variance)

    // Calculate the standard deviation of the numbers
    stdDev := math.Sqrt(variance)
    fmt.Printf("Standard Deviation: %.2f\n", stdDev)
}
