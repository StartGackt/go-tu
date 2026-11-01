package main

import "fmt"

func main() {
    coreTemp := 1000
    pressure := 250
    coolantLevel := 50
    isManualLock := false
    overrideCode := "BETA"

    action := ""

    // ต้องเรียงลำดับจาก กฎข้อ E -> D -> C -> B -> A -> F
    if coreTemp < 0 || pressure < 0 || coolantLevel < 0 || coolantLevel > 100 {
        // กฎข้อ E เซนเซอร์พัง
        action = "SENSOR_ERR"
    } else if coreTemp > 1500 || coolantLevel < 20 || (isManualLock && overrideCode != "ALPHA") {
        // กฎข้อ D ตัดการทำงานฉุกเฉิน
        action = "SCRAM"
    } else if (coreTemp > 1000 || pressure >= 250) && !(coreTemp <= 1200 && overrideCode == "OMEGA" && !isManualLock) {
        // กฎข้อ C สภาวะวิกฤต (พร้อมดักข้อยกเว้นด้วย !)
        action = "CRITICAL"
    } else if coreTemp >= 800 && pressure >= 200 && coolantLevel >= 40 {
        // กฎข้อ B สภาวะเร่งเครื่อง (ยกเว้นน้ำหล่อเย็นต่ำกว่า 40 แปลว่าต้อง >= 40)
        action = "OVERDRIVE"
    } else if coreTemp >= 500 && pressure > 100 {
        // กฎข้อ A สภาวะปกติ
        action = "NORMAL"
    } else {
        // กฎข้อ F เดินเครื่องต่ำ
        action = "IDLE"
    }

    // สั่งงานระบบด้วย switch
    switch action {
    case "SENSOR_ERR":
        fmt.Println("ERR: เซนเซอร์ขัดข้อง ตัดเข้าโหมดสำรอง")
    case "SCRAM":
        fmt.Println("DANGER: ทิ้งแท่งควบคุมและปิดเตาปฏิกรณ์ทันที!")
    case "CRITICAL":
        switch overrideCode {
        case "ALPHA", "BETA":
            fmt.Println("WARN: สภาวะวิกฤต เปิดวาล์วระบายแรงดันฉุกเฉิน")
        default:
            fmt.Println("WARN: สภาวะวิกฤต อพยพเจ้าหน้าที่ออกจากพื้นที่")
        }
    case "OVERDRIVE", "NORMAL":
        if action == "OVERDRIVE" && coolantLevel == 50 {
            fmt.Println("STATUS: เดินเครื่องเต็มกำลัง (เตือน: น้ำหล่อเย็นอยู่ที่ขีดจำกัดครึ่งถัง)")
        } else {
            fmt.Println("STATUS: ระบบทำงานปกติ")
        }
    case "IDLE":
        fmt.Println("STATUS: พลังงานต่ำ เดินเครื่องรอคำสั่ง")
    }
}