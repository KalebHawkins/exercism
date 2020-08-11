//Solution goes in Sources
//
//Calculate the number of grains of wheat on a chessboard given that the number
//on each square doubles.
//
//There once was a wise servant who saved the life of a prince. The king
//promised to pay whatever the servant could dream up. Knowing that the
//king loved chess, the servant told the king he would like to have grains
//of wheat. One grain on the first square of a chess board. Two grains on
//the next. Four on the third, and so on.
//
//There are 64 squares on a chessboard.
//
//Write code that shows:
//- how many grains were on each square, and
//- the total number of grains

import Foundation

struct Grains {
    enum GrainsError: Error {
        case inputTooHigh(message: String)
        case inputTooLow(message: String)
    }
    
    static var total: UInt64 = UInt64.max
    
    static func square(_ squarePosition: Int) throws -> UInt64 {
        let errorString = "Input[\(squarePosition)] invalid. Input should be between 1 and 64 (inclusive)"
        
        guard squarePosition >= 1 else {
            throw GrainsError.inputTooLow(message: errorString)
        }
        
        guard squarePosition <= 64 else {
            throw GrainsError.inputTooHigh(message: errorString)
        }
        
        return UInt64(1) << (squarePosition - 1)
        
    }
}
