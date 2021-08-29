package person

// การกำหนด ชื่อ struct หรือ func เริ่มต้นด้วยตัวใหญ่คือการ expose ให้ด้านนอกมองเห็นได้ หรือ public นั่นเอง ถ้าเป็นตัวเล็กคือ private ใช้ได้แค้ภายใจ package ตัวเอง
// ตัวแปรก็ หลักการเเดียวกัน
type Person struct {
	name    string
	surname string
}

// สร้าง method ให้ struct หรือเรียกว่า receiver function
func (p Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) GetSurname() string {
	return p.surname
}

func (p *Person) SetSurname(surname string) {
	p.surname = surname
}
