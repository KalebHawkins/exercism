//Solution goes in Sources
import Foundation

struct Gigasecond {
    var description: String!
    
    private lazy var dateFormatter: DateFormatter = {
        let formatter = DateFormatter()
        formatter.dateFormat = "yyyy-MM-dd'T'HH:mm:ss"
        formatter.timeZone = TimeZone(secondsFromGMT: 0)
        
        return formatter
    }()
    
    init?(from startDate: String) {
        let gigasecond: Double = 1_000_000_000
        
        if let startDateFromString = dateFormatter.date(from: startDate) {
            let gigasecondFromStartDate = startDateFromString.addingTimeInterval(gigasecond)
            description = dateFormatter.string(from: gigasecondFromStartDate)
        }
    }
}
