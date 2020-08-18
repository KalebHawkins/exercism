//Solution goes in Sources

struct Raindrops {
    var numberOfDrops: Int
    
    init(_ numberOfDrops: Int) {
        self.numberOfDrops = numberOfDrops
    }
    
    private static let harmonics = [
        (3, "Pling"),
        (5, "Plang"),
        (7, "Plong")
    ]
    
    var sounds: String {
        let resultingString = Raindrops.harmonics.reduce("") {
            $0 + (numberOfDrops % $1.0 == 0 ? $1.1 : "")
        }
        
        return resultingString.isEmpty ? String(numberOfDrops) : resultingString
    }
}

extension Int {
    var factors: [Int] {
        return (1...self).filter({ self % $0 == 0 })
    }
}
