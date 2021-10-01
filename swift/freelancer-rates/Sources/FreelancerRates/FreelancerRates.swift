// We first establish a few rules between the freelancer and the project manager:
// 
//     The daily rate is 8 times the hourly rate.
//     A month has 22 billable days.
// 

// The freelancer is offering a discount if the project manager chooses to let the freelancer bill per month, which can come in handy 
// if there is a certain budget the project manager has to work with.
// 
// Discounts are modeled as fractional numbers representing percentage, for example 25.0 (25%).

let dailyRate: Double = 8
let monthlyBillableDays: Double = 22

func dailyRateFrom(hourlyRate: Int) -> Double {
  return dailyRate * Double(hourlyRate)
}

func monthlyRateFrom(hourlyRate: Int, withDiscount discount: Double) -> Double {
  return (dailyRateFrom(hourlyRate: hourlyRate) * monthlyBillableDays - dailyRateFrom(hourlyRate: hourlyRate) * monthlyBillableDays * (discount / 100)).rounded()
}

func workdaysIn(budget: Double, hourlyRate: Int, withDiscount discount: Double) -> Double {
  return (budget / dailyRateFrom(hourlyRate: hourlyRate) * (discount / 100) + budget / dailyRateFrom(hourlyRate: hourlyRate)).rounded()
}
