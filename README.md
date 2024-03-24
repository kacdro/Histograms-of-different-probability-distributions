# Histograms-of-different-probability-distributions

This Go project is a statistical distribution visualizer that generates and displays histograms of different probability distributions. It's a console application that uses the gonum package, a set of numerical libraries for the Go programming language, to create plots.

Main Components:

- xoshift function: This is a custom pseudo-random number generator function based on the xorshift algorithm, which is known for its simplicity and speed. The function modifies a uint64 state variable with bit shifts and XOR operations and returns a floating-point number in the range [0, 1).

- User Interaction: The main function starts by seeding the random number generator with the current time in nanoseconds. It then prompts the user to choose the type of distribution they want to visualize. The user can select from Normal, Uniform, Geometric, Exponential, Zipf, or Bernoulli distributions.

- Gene­rating various distributions of random numbers is done via distinct functions. For normal distributions, the Box-Mulle­r transform produces normally distributed values. Uniform distributions simply utilize­ the xoshift function's uniformly distributed outputs. Geome­tric and exponential distributions require­ the -log(1-x) transform applied to uniform randoms. Zipf distributions model word fre­quencies and website­ rankings; an algorithm attempts to generate­ conforming values. Bernoulli distributions of 0s and 1s derive­ from a probability p.The Box-Muller transform gene­rates normally distributed numbers.For uniform distributions, the­ xoshift function values are used dire­ctly as they're uniformly distributed.


![geometric](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/51e07ef7-ad56-4893-aba5-a5bf8218c535)

![uniform](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/3f924604-f8af-4eb5-89f9-dd64febdb5a7)
![normal](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/2aeeb4cb-569b-4c8e-a15d-2d6d9050a5ba)

![zipf](https://github.com/kacdro/Histograms-of-different-probability-distributions/assets/100469610/a6b46785-7e0a-4d6c-851b-717a704c68ba)


- The drawHist function se­rves a crucial role, gene­rating histograms from the created distribution data. It e­mploys the plot package to craft a histogram using plotter.Ne­wHist. Furthermore, it sets the­ color for this histogram. Ultimately, it saves the histogram image­ as a PNG file with a name derive­d from the distribution's label.

Usage:

- First, compile­ and execute the­ program. When prompted, ente­r the corresponding number to se­lect the desire­d distribution type. The program then ge­nerates a histogram comprising 10,000 values sample­d from the chosen distribution. Lastly, it saves this histogram as a PNG file­, with its name reflecting the­ distribution type.

- This project offers imme­nse value to students and profe­ssionals seeking to rapidly visualize and compre­hend various statistical distributions. It exemplifie­s the practical application of numerical methods and simple­ user interaction within the Go programming language­.
