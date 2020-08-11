//Solution goes in Sources

struct Squares {
    var number: Int
    
    init(_ number: Int){
        self.number = number
    }
    
    var squareOfSum: Int {
        var sum: Int = 0
        for num in 1...number {
            sum += num
        }
        
        return sum*sum
    }
    
    var sumOfSquares: Int {
        var sum: Int = 0
        for num in 1...number {
            sum += num*num
        }
        
        return sum
    }
    
    var differenceOfSquares: Int {
        squareOfSum - sumOfSquares
    }
}
