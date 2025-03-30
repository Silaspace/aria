.EQU DDRD = 0x0A
.EQU PORTD = 0x0B
.EQU PIND = 0x09

	sbi DDRD, 4
loop:
	sbis PIND, 7
	rjmp button_down
	cbi PORTD, 4
	rjmp loop
button_down:
	sbi PORTD, 4
	rjmp loop