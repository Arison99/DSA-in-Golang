package main

import (
	"fmt"
	"math/rand"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func bubbleSort(arr []int, steps chan []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				// Send a copy of the array to the channel
				step := make([]int, len(arr))
				copy(step, arr)
				steps <- step
			}
		}
	}
	close(steps)
}

func plotArray(arr []int, filename string) {
	p := plot.New()
	p.Title.Text = "Bubble Sort Visualization"
	p.Y.Label.Text = "Value"

	// Convert []int to plotter.Values
	values := make(plotter.Values, len(arr))
	for i, v := range arr {
		values[i] = float64(v)
	}

	// Create a bar chart
	bars, err := plotter.NewBarChart(values, vg.Points(20))
	if err != nil {
		panic(err)
	}

	p.Add(bars)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := rand.Perm(20) // Generate a random permutation of numbers from 0 to 19

	fmt.Println("Original array:", arr)

	steps := make(chan []int)
	go bubbleSort(arr, steps)

	i := 0
	for step := range steps {
		filename := fmt.Sprintf("step_%02d.png", i)
		plotArray(step, filename)
		i++
	}

	fmt.Println("Sorted array:", arr)
}
