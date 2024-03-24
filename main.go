package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

var xoshiftState uint64

func xoshift() float64 {
	xoshiftState ^= xoshiftState << 13
	xoshiftState ^= xoshiftState >> 7
	xoshiftState ^= xoshiftState << 17
	return float64(xoshiftState) / (1 << 64)
}

func main() {
	xoshiftState = uint64(time.Now().UnixNano())

	var choice int
	fmt.Print("Choose distribution type:\n1. Normal\n2. Uniform\n3. Geometric\n4. Exponential\n5. Zipf\n6. Bernoulli\n Enter choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		drawNormalDistribution(10000, "normal", "Normal Distribution")
	case 2:
		drawUniformDistribution(10000, "uniform", "Uniform Distribution")
	case 3:
		drawGeometricDistribution(10000, 0.2, "geometric", "Geometric Distribution")
	case 4:
		drawExponentialDistribution(10000, 0.5, "exponential", "Exponential Distribution")
	case 5:
		drawZipfDistribution(10000, 1.5, "zipf", "Zipf Distribution")
	case 6:
		drawBernoulliDistribution(10000, 0.6, "bernoulli", "Bernoulli Distribution")
	default:
		fmt.Println("Invalid choice")
	}
}

func drawNormalDistribution(size int, name, title string) {
	var dist []float64
	for i := 0; i < size; i++ {
		dist = append(dist, generateNormal())
	}
	drawHist(dist, name, title)
}

func generateNormal() float64 {

	u1 := xoshift()
	u2 := xoshift()
	z0 := math.Sqrt(-2*math.Log(u1)) * math.Cos(2*math.Pi*u2)
	return z0
}

func drawUniformDistribution(size int, name, title string) {
	var dist []float64
	for i := 0; i < size; i++ {
		dist = append(dist, xoshift())
	}
	drawHist(dist, name, title)
}

func drawGeometricDistribution(size int, prob float64, name, title string) {
	var dist []float64
	for i := 0; i < size; i++ {
		dist = append(dist, -1/prob*math.Log(1-xoshift()))
	}
	drawHist(dist, name, title)
}

func drawExponentialDistribution(size int, lambda float64, name, title string) {
	var dist []float64
	for i := 0; i < size; i++ {
		dist = append(dist, -1/lambda*math.Log(1-xoshift()))
	}
	drawHist(dist, name, title)
}

func drawZipfDistribution(size int, s float64, name, title string) {
	var dist []float64
	for i := 0; i < 10000; i++ {
		dist = append(dist, float64(generateZipf(1.1, 2.0, 100.0)))
	}
	drawHist(dist, name, title)
}

func generateZipf(s float64, v float64, n float64) int {
	x := xoshift()
	return int(math.Pow(1/(1-x), 1/s) * v * n)
}

func drawBernoulliDistribution(size int, p float64, name, title string) {
	var dist []float64
	for i := 0; i < size; i++ {
		u := xoshift()
		if u <= p {
			dist = append(dist, 1.0)
		} else {
			dist = append(dist, 0.0)
		}
	}
	drawHist(dist, name, title)
}

func drawHist(dist []float64, name, title string) {
	n := len(dist)
	vals := make(plotter.Values, n)
	for i := 0; i < n; i++ {
		vals[i] = dist[i]
	}

	plt := plot.New()
	plt.Title.Text = title
	hist, err := plotter.NewHist(vals, 25)
	if err != nil {
		log.Println("Cannot plot:", err)
	}
	hist.FillColor = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	plt.Add(hist)

	err = plt.Save(800, 400, name+".png")
	if err != nil {
		log.Panic(err)
	}
}
