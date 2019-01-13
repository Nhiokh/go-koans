package goKoans

func aboutAnonymousFunctions() {
	{
		i := 1
		increment := func() {
			i++
		}
		increment()

		assert(i == __int__) // closures function in an obvious way
	}

	{
		i := 1
		increment := func(x int) {
			x++
		}
		increment(i)

		assert(i == __int__) // although anonymous functions need not always be closures
	}

	{
		double := func(x int) int { return x * 2 }

		assert(double(3) == __int__) // they can do anything our hearts desire
	}
}