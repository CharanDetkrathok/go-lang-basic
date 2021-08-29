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

// ใช้ * เพื่ออ้างอิงถึง ตำแหน้งเก็บข้อมูลของ struct นั้นๆที่ถูกสร้าง และเมื่อเรียก receiver function ที่ต้องการ set value ถ้าไม่ใช้ * เวลา assignment value มาที่ func (p Person) เมื่อต้องการเรียกดูค่ามาที่ GetName() ออกมาดูจะทำให้เป็นค่าว่างๆ ("") 
// ดังนั้นจำเป็นต้องใส่ * เมื่อมีการ assignment value ดังด้านล่างนี้ (p *Person) มันจะไปเปลี่ยนข้อมูล ณ ตำแหน่งจริงนั้นๆใน memory
func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) GetSurname() string {
	return p.surname
}

func (p *Person) SetSurname(surname string) {
	p.surname = surname
}
