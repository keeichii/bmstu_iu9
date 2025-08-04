package main

import "fmt"

func encode(utf32 []rune) []byte {
    var result []byte
    for _, r := range utf32 {
        switch {
        case r <= 0x7F:
            result = append(result, byte(r))
        case r <= 0x7FF:
            result = append(result,
                0xC0|byte(r>>6),
                0x80|byte(r&0x3F))
        case r <= 0xFFFF:
            result = append(result,
                0xE0|byte(r>>12),
                0x80|byte((r>>6)&0x3F),
                0x80|byte(r&0x3F))
        case r <= 0x10FFFF:
            result = append(result,
                0xF0|byte(r>>18),
                0x80|byte((r>>12)&0x3F),
                0x80|byte((r>>6)&0x3F),
                0x80|byte(r&0x3F))
        }
    }
    return result
}

func decode(utf8 []byte) []rune {
    var result []rune
    for i := 0; i < len(utf8); {
        b := utf8[i]
        var r rune
        var size int
        
        switch {
        case b&0x80 == 0x00:
            r = rune(b)
            size = 1
        case b&0xE0 == 0xC0:
            r = rune(b&0x1F)<<6 | rune(utf8[i+1]&0x3F)
            size = 2
        case b&0xF0 == 0xE0:
            r = rune(b&0x0F)<<12 | 
                rune(utf8[i+1]&0x3F)<<6 | 
                rune(utf8[i+2]&0x3F)
            size = 3
        case b&0xF8 == 0xF0:
            r = rune(b&0x07)<<18 | 
                rune(utf8[i+1]&0x3F)<<12 | 
                rune(utf8[i+2]&0x3F)<<6 | 
                rune(utf8[i+3]&0x3F)
            size = 4
        }
        
        result = append(result, r)
        i += size
    }
    return result
}

func main() {
    examples := []string{"Hello", "Привет", "😊"}
    
    for _, s := range examples {
        utf32 := []rune(s)
        utf8Bytes := encode(utf32)
        decoded := decode(utf8Bytes)
        fmt.Println(string(decoded))
    }
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
