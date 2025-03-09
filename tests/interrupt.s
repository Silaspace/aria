.DEVICE AT90USB162

    jmp start
    jmp 0x00
    jmp 0x00
    jmp 0x00
    jmp 0x00
    jmp 0x00
    jmp 0x00
    jmp 0x00
    jmp toggle

start:
    ldi r28, 0xFF
    ldi r29, 0x02
    out 0x3E, r29
    out 0x3D, r28
    sbi 0x0A, 4
    lds r24, 0x6A
    ori r24, 0x40
    sts 0x6A, r24
    sbi 0x1D, 0x07
    sei

loop:
    in r24, 0x33
    andi r24, 0xF1
    ori r24, 0x01
    out 0x33, r24
    sleep
    in r24, 0x33
    andi r24, 0xFE
    out 0x33, r24
    rjmp loop

toggle:
    in r26, 0x0B
    ldi r27, 0x10
    eor r26, r27
    out 0x0B, r26
    reti