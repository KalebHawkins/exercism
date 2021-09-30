let birthday = "Birthday"
let valentine = "Valentine's Day"
let anniversary = "Anniversary"

let space: Character = " "
let exclamation: Character = "!"

let costPerCharacter = 2

func buildSign(for occasion: String, name: String) -> String {
  return "Happy \(occasion) \(name)!"
  
}

func graduationFor(name: String, year: Int) -> String {
  return "Congratulations \(name)!\nClass of \(year)"
}

func costOf(sign: String) -> Int {
  return sign.count * costPerCharacter + 20
}
