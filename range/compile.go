package _range

func compile() {
	s := "hello"
	for range s {
	}
}

/*
从汇编码无法清晰地看出其实现原理
E:\RainbowMango\GoExpertProgrammingSourceCode>go tool compile -S ./range/compile.go
"".compile STEXT size=122 args=0x0 locals=0x30
        0x0000 00000 (./range/compile.go:3)     TEXT    "".compile(SB), ABIInternal, $48-0
        0x0000 00000 (./range/compile.go:3)     MOVQ    TLS, CX
        0x0009 00009 (./range/compile.go:3)     PCDATA  $0, $-2
        0x0009 00009 (./range/compile.go:3)     MOVQ    (CX)(TLS*2), CX
        0x0010 00016 (./range/compile.go:3)     PCDATA  $0, $-1
        0x0010 00016 (./range/compile.go:3)     CMPQ    SP, 16(CX)
        0x0014 00020 (./range/compile.go:3)     PCDATA  $0, $-2
        0x0014 00020 (./range/compile.go:3)     JLS     115
        0x0016 00022 (./range/compile.go:3)     PCDATA  $0, $-1
        0x0016 00022 (./range/compile.go:3)     SUBQ    $48, SP
        0x001a 00026 (./range/compile.go:3)     MOVQ    BP, 40(SP)
        0x001f 00031 (./range/compile.go:3)     LEAQ    40(SP), BP
        0x0024 00036 (./range/compile.go:3)     PCDATA  $0, $-2
        0x0024 00036 (./range/compile.go:3)     PCDATA  $1, $-2
        0x0024 00036 (./range/compile.go:3)     FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0024 00036 (./range/compile.go:3)     FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0024 00036 (./range/compile.go:3)     FUNCDATA        $2, gclocals·568470801006e5c0dc3947ea998fe279(SB)
        0x0024 00036 (./range/compile.go:3)     PCDATA  $0, $0
        0x0024 00036 (./range/compile.go:3)     PCDATA  $1, $0
        0x0024 00036 (./range/compile.go:3)     XORL    AX, AX
        0x0026 00038 (./range/compile.go:5)     CMPQ    AX, $5
        0x002a 00042 (./range/compile.go:5)     JGE     105
        0x002c 00044 (./range/compile.go:5)     PCDATA  $0, $1
        0x002c 00044 (./range/compile.go:5)     LEAQ    go.string."hello"(SB), CX
        0x0033 00051 (./range/compile.go:5)     MOVBLZX (CX)(AX*1), DX
        0x0037 00055 (./range/compile.go:5)     CMPL    DX, $128
        0x003d 00061 (./range/compile.go:5)     JGE     68
        0x003f 00063 (./range/compile.go:5)     PCDATA  $0, $0
        0x003f 00063 (./range/compile.go:5)     INCQ    AX
        0x0042 00066 (./range/compile.go:5)     JMP     38
        0x0044 00068 (./range/compile.go:5)     MOVQ    CX, (SP)
        0x0048 00072 (./range/compile.go:5)     MOVQ    $5, 8(SP)
        0x0051 00081 (./range/compile.go:5)     MOVQ    AX, 16(SP)
        0x0056 00086 (./range/compile.go:5)     CALL    runtime.decoderune(SB)
        0x005b 00091 (./range/compile.go:5)     MOVQ    32(SP), AX
        0x0060 00096 (./range/compile.go:5)     LEAQ    go.string."hello"(SB), CX
        0x0067 00103 (./range/compile.go:5)     JMP     38
        0x0069 00105 (./range/compile.go:5)     PCDATA  $0, $-1
        0x0069 00105 (./range/compile.go:5)     PCDATA  $1, $-1
        0x0069 00105 (./range/compile.go:5)     MOVQ    40(SP), BP
        0x006e 00110 (./range/compile.go:5)     ADDQ    $48, SP
        0x0072 00114 (./range/compile.go:5)     RET
        0x0073 00115 (./range/compile.go:5)     NOP
        0x0073 00115 (./range/compile.go:3)     PCDATA  $1, $-1
        0x0073 00115 (./range/compile.go:3)     PCDATA  $0, $-2
        0x0073 00115 (./range/compile.go:3)     CALL    runtime.morestack_noctxt(SB)
        0x0078 00120 (./range/compile.go:3)     PCDATA  $0, $-1
        0x0078 00120 (./range/compile.go:3)     JMP     0
        0x0000 65 48 8b 0c 25 28 00 00 00 48 8b 89 00 00 00 00  eH..%(...H......
        0x0010 48 3b 61 10 76 5d 48 83 ec 30 48 89 6c 24 28 48  H;a.v]H..0H.l$(H
        0x0020 8d 6c 24 28 31 c0 48 83 f8 05 7d 3d 48 8d 0d 00  .l$(1.H...}=H...
        0x0030 00 00 00 0f b6 14 01 81 fa 80 00 00 00 7d 05 48  .............}.H
        0x0040 ff c0 eb e2 48 89 0c 24 48 c7 44 24 08 05 00 00  ....H..$H.D$....
        0x0050 00 48 89 44 24 10 e8 00 00 00 00 48 8b 44 24 20  .H.D$......H.D$
        0x0060 48 8d 0d 00 00 00 00 eb bd 48 8b 6c 24 28 48 83  H........H.l$(H.
        0x0070 c4 30 c3 e8 00 00 00 00 eb 86                    .0........
        rel 12+4 t=17 TLS+0
        rel 47+4 t=16 go.string."hello"+0
        rel 87+4 t=8 runtime.decoderune+0
        rel 99+4 t=16 go.string."hello"+0
        rel 116+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFINFO dupok size=0
        0x0000 5f 72 61 6e 67 65                                _range
go.string."hello" SRODATA dupok size=5
        0x0000 68 65 6c 6c 6f                                   hello
go.loc."".compile SDWARFLOC size=0
go.info."".compile SDWARFINFO size=45
        0x0000 03 22 22 2e 63 6f 6d 70 69 6c 65 00 00 00 00 00  ."".compile.....
        0x0010 00 00 00 00 00 00 00 00 00 00 00 00 01 9c 00 00  ................
        0x0020 00 00 01 0a 73 00 04 00 00 00 00 00 00           ....s........
        rel 12+8 t=1 "".compile+0
        rel 20+8 t=1 "".compile+122
        rel 30+4 t=30 gofile..E:\RainbowMango\GoExpertProgrammingSourceCode\range\compile.go+0
        rel 39+4 t=29 go.info.string+0
go.range."".compile SDWARFRANGE size=0
go.debuglines."".compile SDWARFMISC size=21
        0x0000 04 02 11 0a eb 06 9b 06 25 06 37 06 02 27 ff 71  ........%.7..'.q
        0x0010 04 01 03 7e 01                                   ...~.
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·568470801006e5c0dc3947ea998fe279 SRODATA dupok size=10
        0x0000 02 00 00 00 02 00 00 00 00 02                    ..........

*/
