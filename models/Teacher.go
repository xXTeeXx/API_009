// models/Teacher.go
package models

import "gorm.io/gorm"

type Teacher struct {
    gorm.Model // เพิ่ม ID, CreatedAt, UpdatedAt, DeletedAt ให้โครงสร้างข้อมูล

    FirstName string
    LastName  string
    Age       int
    Department string // เปลี่ยน Grade เป็น Department เนื่องจากเรากำลังพูดถึงครู
}
