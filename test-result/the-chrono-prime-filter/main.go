package main

import "fmt"

func main() {
    year := 1900
    energy := 91
    keyType := "GOLD"

    isLeapYear := false
    isPseudoPrime := false
    gateStatus := ""

    // ขั้นตอนที่ 1 ตรวจปีอธิกสุรทิน (ต้องเรียงจากสับเซตแคบสุด 400 -> 100 -> 4)
    if year%400 == 0 {
        isLeapYear = true
    } else if year%100 == 0 {
        isLeapYear = false
    } else if year%4 == 0 {
        isLeapYear = true
    } else {
        isLeapYear = false
    }

    // ขั้นตอนที่ 2 ตรวจจำนวนเฉพาะเทียม
    if energy == 2 || energy == 3 || energy == 5 || energy == 7 {
        isPseudoPrime = true
    } else if energy%2 != 0 && energy%3 != 0 && energy%5 != 0 && energy%7 != 0 {
        isPseudoPrime = true
    } else {
        isPseudoPrime = false
    }

    // ขั้นตอนที่ 3 ตัดสินใจสถานะประตูมิติ
    if year <= 0 || energy < 2 || energy > 100 {
        gateStatus = "CORRUPTED"
    } else if isLeapYear && isPseudoPrime {
        gateStatus = "QUANTUM_SYNC"
    } else if isLeapYear || isPseudoPrime {
        // เนื่องจากกรณีที่เป็น true ทั้งคู่ถูกดักไปแล้ว บรรทัดนี้จึงหมายถึงเป็น true แค่ตัวเดียวโดยอัตโนมัติ
        gateStatus = "PARTIAL_SYNC"
    } else {
        gateStatus = "DESYNC"
    }

    // แสดงผลด้วย switch
    switch gateStatus {
    case "CORRUPTED":
        fmt.Println("ERROR: พิกัดเวลาหรือพลังงานผิดพลาด")
    case "QUANTUM_SYNC":
        fmt.Println("OPEN: เปิดประตูมิติสมบูรณ์แบบ")
    case "PARTIAL_SYNC":
        switch keyType {
        case "GOLD":
            fmt.Println("OPEN: เปิดประตูมิติด้วยกุญแจทองคำทดแทน")
        case "SILVER", "WOOD":
            fmt.Println("LOCKED: พลังงานไม่เสถียร กุญแจของคุณระดับต่ำเกินไป")
        default:
            fmt.Println("ERROR: ไม่รู้จักชนิดกุญแจ")
        }
    case "DESYNC":
        fmt.Println("LOCKED: พิกัดเวลาและพลังงานไม่สอดคล้องกันเลย")
    }
}