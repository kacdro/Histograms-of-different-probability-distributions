# Histograms-of-different-probability-distributions

This Go project is a statistical distribution visualizer that generates and displays histograms of different probability distributions. It's a console application that uses the gonum package, a set of numerical libraries for the Go programming language, to create plots.

Main Components:

xoshift function: This is a custom pseudo-random number generator function based on the xorshift algorithm, which is known for its simplicity and speed. The function modifies a uint64 state variable with bit shifts and XOR operations and returns a floating-point number in the range [0, 1).

User Interaction: The main function starts by seeding the random number generator with the current time in nanoseconds. It then prompts the user to choose the type of distribution they want to visualize. The user can select from Normal, Uniform, Geometric, Exponential, Zipf, or Bernoulli distributions.

![geometric](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/51e07ef7-ad56-4893-aba5-a5bf8218c535)

![uniform](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/3f924604-f8af-4eb5-89f9-dd64febdb5a7)
![normal](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/2aeeb4cb-569b-4c8e-a15d-2d6d9050a5ba)

![zipf](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/a6b46785-7e0a-4d6c-851b-717a704c68ba)
