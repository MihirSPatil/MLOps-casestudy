package luhn

import (
    "strings"
    "unicode"
)

func Valid(id string) bool {
    clean_id := strings.ReplaceAll(id, " ", "")

    sum := 0
    double := false
    valid := false

    if len(clean_id) > 1 {
        for idx := len(clean_id) - 1; idx >= 0; idx-- {
            val := rune(clean_id[idx])

            if !unicode.IsDigit(val) {
                return false
            }

            digit := int(val - '0')
            if double {
                digit *= 2
                if digit > 9 {
                    digit -= 9
                }
            }
            double = !double
            sum += digit
        }

        if sum%10 == 0 { valid = true }else { valid = false}
    }
    return valid
}
