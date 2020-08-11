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

    var onEarth: Double { convertOrbitablePeriod(from: seconds, on: .Earth).rounded(to: 2) }
    var onMercury: Double { convertOrbitablePeriod(from: seconds, on: .Mercury).rounded(to: 2) }
    var onVenus: Double { convertOrbitablePeriod(from: seconds, on: .Venus).rounded(to: 2) }
    var onMars: Double { convertOrbitablePeriod(from: seconds, on: .Mars).rounded(to: 2) }
    var onJupiter: Double { convertOrbitablePeriod(from: seconds, on: .Jupiter).rounded(to: 2) }
    var onSaturn: Double { convertOrbitablePeriod(from: seconds, on: .Saturn).rounded(to: 2) }
    var onUranus: Double { convertOrbitablePeriod(from: seconds, on: .Uranus).rounded(to: 2) }
    var onNeptune: Double { convertOrbitablePeriod(from: seconds, on: .Neptune).rounded(to: 2) }
    
    private func convertOrbitablePeriod(from seconds: Double, on planet: orbitalPeriod) -> Double {
        if planet == .Earth {
            return (seconds / secondsPerYearOnEarth)
        }

        let planetsAgeInYears = seconds / secondsPerYearOnEarth / planet.rawValue
        return planetsAgeInYears
    }
}

fileprivate enum orbitalPeriod: Double {
    typealias RawValue = Double

    case Earth = 365.25
    case Mercury = 0.2408467
    case Venus = 0.61519726
    case Mars = 1.8808158
    case Jupiter = 11.862615
    case Saturn = 29.447498
    case Uranus = 84.016846
    case Neptune = 164.79132
}

extension Double {
    func rounded(to decimalPlaces: Int) -> Double {
        let divisor = pow(10.0, Double(decimalPlaces))
        return ( self * divisor ).rounded() / divisor
    }
}
