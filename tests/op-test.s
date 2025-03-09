; TESTING OPERANDS

; register
add r0

; imm
add 765

; bin
add 0b101001

; hex
add 0xaf8

; octal
add 076

; label
add loop

; multiple op
add r4, r5

; unary op
add -6

; binary op
add 8 - 6

; brackets
add (8 - 6) + 5

; wide op
add 8 < 6
add 8 << 6

; spacing
add 8+6+4  + 9

; complex expr
add 8<<6 + (-8)

; single char of wide op
add 8<6

; eq expr
add 8 =6

; neq expr
add 8 != 6

; more complex expr
add (label << 2) + (4*0xf-077)

; function
add func(arg)

; function with spaces
add func ( arg < 9 )