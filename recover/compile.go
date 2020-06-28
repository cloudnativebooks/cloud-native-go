package recover

func compile() {
	defer func() {
		recover()
	}()
}

/*
E:\RainbowMango\GoExpertProgrammingSourceCode>go tool compile -S recover/compile.go
"".compile STEXT size=107 args=0x0 locals=0x18
        0x0000 00000 (recover/compile.go:3)     TEXT    "".compile(SB), ABIInternal, $24-0
        0x0000 00000 (recover/compile.go:3)     MOVQ    TLS, CX
        0x0009 00009 (recover/compile.go:3)     PCDATA  $0, $-2
        0x0009 00009 (recover/compile.go:3)     MOVQ    (CX)(TLS*2), CX
        0x0010 00016 (recover/compile.go:3)     PCDATA  $0, $-1
        0x0010 00016 (recover/compile.go:3)     CMPQ    SP, 16(CX)
        0x0014 00020 (recover/compile.go:3)     PCDATA  $0, $-2
        0x0014 00020 (recover/compile.go:3)     JLS     100
        0x0016 00022 (recover/compile.go:3)     PCDATA  $0, $-1
        0x0016 00022 (recover/compile.go:3)     SUBQ    $24, SP
        0x001a 00026 (recover/compile.go:3)     MOVQ    BP, 16(SP)
        0x001f 00031 (recover/compile.go:3)     LEAQ    16(SP), BP
        0x0024 00036 (recover/compile.go:3)     MOVQ    $0, AX
        0x002b 00043 (recover/compile.go:3)     MOVQ    AX, 8(SP)
        0x0030 00048 (recover/compile.go:3)     PCDATA  $0, $-2
        0x0030 00048 (recover/compile.go:3)     PCDATA  $1, $-2
        0x0030 00048 (recover/compile.go:3)     FUNCDATA        $0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
        0x0030 00048 (recover/compile.go:3)     FUNCDATA        $1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0030 00048 (recover/compile.go:3)     FUNCDATA        $2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0030 00048 (recover/compile.go:3)     FUNCDATA        $5, "".compile.opendefer(SB)
        0x0030 00048 (recover/compile.go:3)     PCDATA  $0, $0
        0x0030 00048 (recover/compile.go:3)     PCDATA  $1, $0
        0x0030 00048 (recover/compile.go:3)     MOVB    $0, ""..autotmp_1+7(SP)
        0x0035 00053 (recover/compile.go:4)     PCDATA  $0, $1
        0x0035 00053 (recover/compile.go:4)     LEAQ    "".compile.func1·f(SB), AX
        0x003c 00060 (recover/compile.go:4)     PCDATA  $0, $0
        0x003c 00060 (recover/compile.go:4)     PCDATA  $1, $1
        0x003c 00060 (recover/compile.go:4)     MOVQ    AX, ""..autotmp_2+8(SP)
        0x0041 00065 (recover/compile.go:7)     MOVB    $0, ""..autotmp_1+7(SP)
        0x0046 00070 (recover/compile.go:7)     CALL    "".compile.func1(SB)
        0x004b 00075 (recover/compile.go:7)     MOVQ    16(SP), BP
        0x0050 00080 (recover/compile.go:7)     ADDQ    $24, SP
        0x0054 00084 (recover/compile.go:7)     RET
        0x0055 00085 (recover/compile.go:7)     CALL    runtime.deferreturn(SB)
        0x005a 00090 (recover/compile.go:7)     MOVQ    16(SP), BP
        0x005f 00095 (recover/compile.go:7)     ADDQ    $24, SP
        0x0063 00099 (recover/compile.go:7)     RET
        0x0064 00100 (recover/compile.go:7)     NOP
        0x0064 00100 (recover/compile.go:3)     PCDATA  $1, $-1
        0x0064 00100 (recover/compile.go:3)     PCDATA  $0, $-2
        0x0064 00100 (recover/compile.go:3)     CALL    runtime.morestack_noctxt(SB)
        0x0069 00105 (recover/compile.go:3)     PCDATA  $0, $-1
        0x0069 00105 (recover/compile.go:3)     JMP     0
        0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
        0x0010 48 3b 61 10 76 4e 48 83 ec 18 48 89 6c 24 10 48  H;a.vNH...H.l$.H
        0x0020 8d 6c 24 10 48 c7 c0 00 00 00 00 48 89 44 24 08  .l$.H......H.D$.
        0x0030 c6 44 24 07 00 48 8d 05 00 00 00 00 48 89 44 24  .D$..H......H.D$
        0x0040 08 c6 44 24 07 00 e8 00 00 00 00 48 8b 6c 24 10  ..D$.......H.l$.
        0x0050 48 83 c4 18 c3 e8 00 00 00 00 48 8b 6c 24 10 48  H.........H.l$.H
        0x0060 83 c4 18 c3 e8 00 00 00 00 eb 95                 ...........
        rel 12+4 t=17 TLS+0
        rel 56+4 t=16 "".compile.func1·f+0
        rel 71+4 t=8 "".compile.func1+0
        rel 86+4 t=8 runtime.deferreturn+0
        rel 101+4 t=8 runtime.morestack_noctxt+0
"".compile.func1 STEXT size=67 args=0x0 locals=0x20
        0x0000 00000 (recover/compile.go:4)     TEXT    "".compile.func1(SB), ABIInternal, $32-0
        0x0000 00000 (recover/compile.go:4)     MOVQ    TLS, CX
        0x0009 00009 (recover/compile.go:4)     PCDATA  $0, $-2
        0x0009 00009 (recover/compile.go:4)     MOVQ    (CX)(TLS*2), CX
        0x0010 00016 (recover/compile.go:4)     PCDATA  $0, $-1
        0x0010 00016 (recover/compile.go:4)     CMPQ    SP, 16(CX)
        0x0014 00020 (recover/compile.go:4)     PCDATA  $0, $-2
        0x0014 00020 (recover/compile.go:4)     JLS     60
        0x0016 00022 (recover/compile.go:4)     PCDATA  $0, $-1
        0x0016 00022 (recover/compile.go:4)     SUBQ    $32, SP
        0x001a 00026 (recover/compile.go:4)     MOVQ    BP, 24(SP)
        0x001f 00031 (recover/compile.go:4)     LEAQ    24(SP), BP
        0x0024 00036 (recover/compile.go:4)     PCDATA  $0, $-2
        0x0024 00036 (recover/compile.go:4)     PCDATA  $1, $-2
        0x0024 00036 (recover/compile.go:4)     FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0024 00036 (recover/compile.go:4)     FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0024 00036 (recover/compile.go:4)     FUNCDATA        $2, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
        0x0024 00036 (recover/compile.go:5)     PCDATA  $0, $1
        0x0024 00036 (recover/compile.go:5)     PCDATA  $1, $0
        0x0024 00036 (recover/compile.go:5)     LEAQ    ""..fp+40(SP), AX
        0x0029 00041 (recover/compile.go:5)     PCDATA  $0, $0
        0x0029 00041 (recover/compile.go:5)     MOVQ    AX, (SP)
        0x002d 00045 (recover/compile.go:5)     CALL    runtime.gorecover(SB)
        0x0032 00050 (recover/compile.go:6)     MOVQ    24(SP), BP
        0x0037 00055 (recover/compile.go:6)     ADDQ    $32, SP
        0x003b 00059 (recover/compile.go:6)     RET
        0x003c 00060 (recover/compile.go:6)     NOP
        0x003c 00060 (recover/compile.go:4)     PCDATA  $1, $-1
        0x003c 00060 (recover/compile.go:4)     PCDATA  $0, $-2
        0x003c 00060 (recover/compile.go:4)     CALL    runtime.morestack_noctxt(SB)
        0x0041 00065 (recover/compile.go:4)     PCDATA  $0, $-1
        0x0041 00065 (recover/compile.go:4)     JMP     0
        0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
        0x0010 48 3b 61 10 76 26 48 83 ec 20 48 89 6c 24 18 48  H;a.v&H.. H.l$.H
        0x0020 8d 6c 24 18 48 8d 44 24 28 48 89 04 24 e8 00 00  .l$.H.D$(H..$...
        0x0030 00 00 48 8b 6c 24 18 48 83 c4 20 c3 e8 00 00 00  ..H.l$.H.. .....
        0x0040 00 eb bd                                         ...
        rel 12+4 t=17 TLS+0
        rel 46+4 t=8 runtime.gorecover+0
        rel 61+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 72 65 63 6f 76 65 72                             recover
go.loc."".compile SDWARFLOC size=0
go.info."".compile SDWARFINFO size=36
        0x0000 03 22 22 2e 63 6f 6d 70 69 6c 65 00 00 00 00 00  ."".compile.....
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
        0x0020 00 00 01 00                                      ....
        rel 0+0 t=24 type.func()+0
        rel 0+0 t=24 type.uint8+0
        rel 12+8 t=1 "".compile+0
        rel 20+8 t=1 "".compile+107
        rel 30+4 t=30 gofile..E:\RainbowMango\GoExpertProgrammingSourceCode\recover\compile.go+0
go.range."".compile SDWARFRANGE size=0
go.debuglines."".compile SDWARFMISC size=22
        0x0000 04 02 11 0a eb 08 56 06 55 06 44 06 41 06 41 08  ......V.U.D.A.A.
        0x0010 15 04 01 03 7e 01                                ....~.
go.loc."".compile.func1 SDWARFLOC size=0
go.info."".compile.func1 SDWARFINFO size=42
        0x0000 03 22 22 2e 63 6f 6d 70 69 6c 65 2e 66 75 6e 63  ."".compile.func
        0x0010 31 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  1...............
        0x0020 00 00 01 9c 00 00 00 00 01 00                    ..........
        rel 18+8 t=1 "".compile.func1+0
        rel 26+8 t=1 "".compile.func1+67
        rel 36+4 t=30 gofile..E:\RainbowMango\GoExpertProgrammingSourceCode\recover\compile.go+0
go.range."".compile.func1 SDWARFRANGE size=0
go.debuglines."".compile.func1 SDWARFMISC size=16
        0x0000 04 02 12 0a eb 9c 06 41 06 6a 71 04 01 03 7d 01  .......A.jq...}.
"".compile.func1·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 "".compile.func1+0
runtime.memequal64·f SRODATA dupok size=8
        0x0000 00 00 00 00 00 00 00 00                          ........
        rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
        0x0000 01                                               .
type..namedata.*func()- SRODATA dupok size=10
        0x0000 00 00 07 2a 66 75 6e 63 28 29                    ...*func()
type.*func() SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 9b 90 75 1b 08 08 08 36 00 00 00 00 00 00 00 00  ..u....6........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00 00 00 00 00                          ........
        rel 24+8 t=1 runtime.memequal64·f+0
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func()-+0
        rel 48+8 t=1 type.func()+0
type.func() SRODATA dupok size=56
        0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
        0x0010 f6 bc 82 f6 02 08 08 33 00 00 00 00 00 00 00 00  .......3........
        0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0030 00 00 00 00                                      ....
        rel 32+8 t=1 runtime.gcbits.01+0
        rel 40+4 t=5 type..namedata.*func()-+0
        rel 44+4 t=6 type.*func()+0
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
        0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
"".compile.opendefer SRODATA dupok size=6
        0x0000 00 09 01 00 08 00                                ......
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........

*/
