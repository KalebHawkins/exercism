//Solution goes in Sources

//Given an age in seconds, calculate how old someone would be on:
//
//   - Earth: orbital period 365.25 Earth days, or 31557600 seconds
//   - Mercury: orbital period 0.2408467 Earth years
//   - Venus: orbital period 0.61519726 Earth years
//   - Mars: orbital period 1.8808158 Earth years
//   - Jupiter: orbital period 11.862615 Earth years
//   - Saturn: orbital period 29.447498 Earth years
//   - Uranus: orbital period 84.016846 Earth years
//   - Neptune: orbital period 164.79132 Earth years
//
//So if you were told someone were 1,000,000,000 seconds old, you should
//be able to say that they're 31.69 Earth-years old.


import Foundation

//Solution goes in Sources

struct SpaceAge {
    var seconds: Double
    private var secondsPerYearOnEarth: Double = 31557600.0

    init(_ seconds: Double) {
        self.seconds = seconds
    }

    var onEarth: Double { convertOrbitablePeriod(from: seconds, on: .onEarth).rounded(to: 2) }
    var onMercury: Double { convertOrbitablePeriod(from: seconds, on: .onMercury).rounded(to: 2) }
    var onVenus: Double { convertOrbitablePeriod(from: seconds, on: .onVenus).rounded(to: 2) }
    var onMars: Double { convertOrbitablePeriod(from: seconds, on: .onMars).rounded(to: 2) }
    var onJupiter: Double { convertOrbitablePeriod(from: seconds, on: .onJupiter).rounded(to: 2) }
    var onSaturn: Double { convertOrbitablePeriod(from: seconds, on: .onSaturn).rounded(to: 2) }
    var onUranus: Double { convertOrbitablePeriod(from: seconds, on: .onUranus).rounded(to: 2) }
    var onNeptune: Double { convertOrbitablePeriod(from: seconds, on: .onNeptune).rounded(to: 2) }
    
    private func convertOrbitablePeriod(from seconds: Double, on planet: orbitalPeriod) -> Double {
        if planet == .onEarth {
            return (seconds / secondsPerYearOnEarth)
        }

        let planetsAgeInYears = seconds / secondsPerYearOnEarth / planet.rawValue
        return planetsAgeInYears
    }
}

fileprivate enum orbitalPeriod: Double {
    typealias RawValue = Double

    case onEarth = 365.25
    case onMercury = 0.2408467
    case onVenus = 0.61519726
    case onMars = 1.8808158
    case onJupiter = 11.862615
    case onSaturn = 29.447498
    case onUranus = 84.016846
    case onNeptune = 164.79132
}

extension Double {
    func rounded(to decimalPlaces: Int) -> Double {
        let divisor = pow(10.0, Double(decimalPlaces))
        return ( self * divisor ).rounded() / divisor
    }
}
