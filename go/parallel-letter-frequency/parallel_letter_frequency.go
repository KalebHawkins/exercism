package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	result := FreqMap{}
	ch := make(chan FreqMap)

	for _, s := range l {
		go func(s string) {
			ch <- Frequency(s)
		}(s)
	}

	for range l {
		m := <-ch
		for k, v := range m {
			result[k] += v
		}
	}

	return result
}
