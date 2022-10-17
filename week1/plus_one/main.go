package main


func main() {
    plusOne([]int{1, 2, 3, 4})
}

func plusOne(digits []int) []int {
    lenDigits := len(digits)
    for i := lenDigits - 1; i >= 0; i-- {
        if digits[i] != 9 {
            digits[i] += 1
            return digits
        } else {
            digits[i] = 0
        }
    }

    if digits[0] == 0 && lenDigits > 0{
        digits = append(digits, 1)
        copy(digits[1:], digits[0:])
        digits[0] = 1
    }

    return digits
}
