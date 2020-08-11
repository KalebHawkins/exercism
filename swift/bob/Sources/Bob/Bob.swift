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

import Foundation

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
        return self == self.uppercased()
    }
    
    func hasLetters() -> Bool {
        return self.rangeOfCharacter(from: .letters) != nil
    }
    
    func isQuestion() -> Bool {
        return self.suffix(1) == "?"
    }
    
    func isAllWhiteSpace() -> Bool {
        return self.trimmingCharacters(in: .whitespaces).isEmpty
    }

}
