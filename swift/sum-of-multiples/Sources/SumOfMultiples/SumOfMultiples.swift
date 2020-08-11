//Solution goes in Sources

//Given a number, find the sum of all the unique multiples of particular numbers up to but not including that number.
//If we list all the natural numbers below 20 that are multiples of 3 or 5, we get 3, 5, 6, 9, 10, 12, 15, and 18.
//The sum of these multiples is 78.

struct SumOfMultiples {
    static func toLimit(_ limit: Int, inMultiples: [Int]) -> Int {
        // If the limit is zero then the result will always be zero.
        guard limit != 0 else {
            return 0
        }
        
        // Zeros need to be filtered out of the inMultiples array. This is done easily with the filter function.
        let inMultiplesZerosStripped = inMultiples.filter { $0 != 0 }
        var multiples = [Int]()
        
        // For each number in the array find each multiple and store them in an array.
        for number in inMultiplesZerosStripped {
            multiples.append(contentsOf: getMultiples(limit: limit, number: number))
        }
        
        // The array must be stipped of duplcates and then summed.
        // This is achieved by converting the multiples array into a set and back into an array
        // The numbers in the array are then summed using the reduce function.
        return Array(Set(multiples)).reduce(0, +)
    }
    
    // getMultiples is a helper function to find the multiples of a number passed to it.
    // the multiples are returned as an array.
    private static func getMultiples(limit: Int, number: Int) -> [Int] {
        var arr = [Int]()
        
        // Don't start this loop at 0... Division by zero is infinitly bad :).
        for x in 1..<limit {
            if x % number == 0 {
                arr.append(x)
            }
        }
        
        return arr
    }
}
