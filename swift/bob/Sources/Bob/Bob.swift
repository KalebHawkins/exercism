//Solution goes in Sources
//
//Bob is a lackadaisical teenager. In conversation, his responses are very limited.
//
//Bob answers 'Sure.' if you ask him a question.
//
//He answers 'Whoa, chill out!' if you yell at him.
//
//He says 'Fine. Be that way!' if you address him without actually saying
//anything.
//
//He answers 'Whatever.' to anything else.

struct Bob {
    static func hey(_ input: String) -> String {
        var output: String = "Whatever."
        
        if input.isAllLettersUpper() && input.hasLetters() {
            output = "Whoa, chill out!"
        } else if input.isQuestion() {
            output = "Sure."
        } else if input.isAllWhiteSpace() {
            output = "Fine. Be that way!"
        }
        
        
        return output
    }
}

extension String {
    func isAllLettersUpper() -> Bool {
        for char in self {
            if char.isLetter && char.isLowercase  {
                return false
            }
        }
        
        return true
    }
    
    func hasLetters() -> Bool {
        for char in self {
            if char.isLetter {
                return true
            }
        }
        return false
    }
    
    func isQuestion() -> Bool {
        if self.suffix(1) == "?" {
            return true
        }
        
        return false
    }
    
    func isAllWhiteSpace() -> Bool {
        for char in self {
            if !char.isWhitespace {
                return false
            }
        }
        return true
    }
}
