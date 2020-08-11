//Solution goes in Sources

//Given a DNA strand, return its RNA complement (per RNA transcription).
//
//Both DNA and RNA strands are a sequence of nucleotides.
//
//The four nucleotides found in DNA are adenine (A), cytosine (C), guanine (G) and thymine (T).
//
//The four nucleotides found in RNA are adenine (A), cytosine (C), guanine (G) and uracil (U).
//
//Given a DNA strand, its transcribed RNA strand is formed by replacing each nucleotide with its complement:
//
//G -> C
//C -> G
//T -> A
//A -> U


struct Nucleotide {
    let dnaStrand: String
    private var complments: [Character: Character] = ["G":"C", "C":"G", "T":"A", "A":"U"]
    
    init(_ dnaStrand: String) {
        self.dnaStrand = dnaStrand
    }
    
    func complementOfDNA() throws -> String {
        
        // Begin mapping each character(Nucleotide) in the DNA sequence to the appropriate RNA complement.
        // I could've wrapped the entire map statement with String() but I feel like it reduced readability so I convert
        // it from a Character array to a String in the return statement.
        let transcription = try dnaStrand.map { char -> Character in
            
            // Guard against invalid nucleotides from trying to be mapped.
            // If an invalid nucleotide is passed an error is thrown.
            guard let nucleotide = self.complments[char] else {
                throw TranscriptionError.invalidNucleotide("\(char) is not a valid Nucleotide")
            }
            
            // Return the RNA nucleotide.
            return nucleotide
        }
        
        return String(transcription)
    }
    
}

enum TranscriptionError: Error {
    case invalidNucleotide(String)
}
