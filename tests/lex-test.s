;comment line
    ;indented comment

;label followed by instruction
label:
    break

;label and instr on same line
label2: break

;directive followed by comment
.directive ; comment 3

;instr with register op
instr r4

;instr with immediate op
instr 8756

;instr with hex op
instr 0xfa6

;instr with octal op
instr 0675

;instr with identifier op
instr label

;instr with implicit expression op
instr label << 2

;instr with explicit expression op
instr (label << 2)

;instr with function op
instr FUNC(ident)

;instr with function op and spacing
instr FUNC ( ident ) 

;instr with complex expression op
instr (FUNC(ident) * 3) + 4

;instr with multiple ops
instr r4, r5

;instr with multiple expression ops
instr (HIGH(0xff) - ident), ~07700

;whole word register
adiw r4:r5, 0xff

;unexpected newline
add 7 <
9

;unexpected comment
add 7 < ;comment

;odd label
r23gh

; instr / ident distinguishing
rjmp label
label rjmp

; func / ident distinguishing
func(low)
low(func)