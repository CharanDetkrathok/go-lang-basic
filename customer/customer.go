package customer

var name = "this Customer package Now!!!"

// การกำหนด ชื่อ func เริ่มต้นด้วยตัวใหญ่คือการ expose ให้ด้านนอกมองเห็นได้ หรือ public นั่นเอง ถ้าเป็นตัวเล็กคือ private ใช้ได้แค้ภายใจ package ตัวเอง
// ตัวแปรก็ หลักการเเดียวกัน
func Hello(otherName string) string {
	return "Hello " + otherName + " " + name
}
