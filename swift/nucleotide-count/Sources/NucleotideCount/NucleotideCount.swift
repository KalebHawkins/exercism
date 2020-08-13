//Solution goes in Sources

//Given a single stranded DNA string, compute how many times each nucleotide occurs in the string.
//
//The genetic language of every living thing on the planet is DNA.
//DNA is a large molecule that is built from an extremely long sequence of individual elements called nucleotides.
//4 types exist in DNA and these differ only slightly and can be represented as the following symbols: 'A' for adenine, 'C' for cytosine, 'G' for guanine, and 'T' thymine.
//
//Here is an analogy:
//- twigs are to birds nests as
//- nucleotides are to DNA as
//- legos are to lego houses as
//- words are to sentences as...

import Foundation

struct DNA {
    var strand: String
    
    private var nucleotides: [Character] = [ "A", "C", "G", "T" ]
        
    init?(strand: String){
        guard Set(strand).isSubset(of: Set(nucleotides)) else {
            return nil
        }
        
        self.strand = strand
    }
    
    func count(_ nucleotide: Character) -> Int {
        return strand.filter { $0 == nucleotide }.count
    }
    
    func counts() -> [String: Int] {
        return nucleotides.reduce(into: [String: Int]()) { (key, value) in
            key[String(value)] = count(value)
        }
    }
}
