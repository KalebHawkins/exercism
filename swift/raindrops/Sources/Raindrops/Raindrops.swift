//Solution goes in Sources

struct Raindrops {
    var numberOfDrops: Int
    
    init(_ numberOfDrops: Int) {
        self.numberOfDrops = numberOfDrops
    }
    
    var sounds: String {
        let factors = numberOfDrops.getFactors()

        var resultingString = String()
        
        if factors.contains(3) {
            resultingString.append("Pling")
        }
        if factors.contains(5) {
            resultingString.append("Plang")
        }
        if factors.contains(7) {
            resultingString.append("Plong")
        }
        
        return resultingString.isEmpty ? String(self.numberOfDrops) : resultingString
    }
}

extension Int {
    func getFactors() -> [Int] {
        return (1...self).filter({ self % $0 == 0 })
    }
}
