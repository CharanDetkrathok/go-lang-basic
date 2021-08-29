package main

import (
	"fmt"
	"hello/customer"
	"hello/person"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello World")

	a := "กขค"
	B := "abc"
	C := "123"
	println("abc นับความยาวของ string แบบ ภาษาอังกฤษและตัวเลข ", len(B), " นับด้วย Builtin func len(B) ได้ ตรง")
	println("123 นับความยาวของ string แบบ ภาษาอังกฤษและตัวเลข ", len(C), " นับด้วย Builtin func len(C) ได้ ตรง")

	println("กขค นับความยาวของ string แบบ ภาษาอังกฤษและตัวเลข ", len(a), " นับด้วย Builtin func len(a) จะไม่ตรง -กขค- ความยาวต้องเป็น 3")
	println("กขค นับความยาวของ string แบบ ภาษาไทยให้ตรง => ", utf8.RuneCountInString(a), " นับด้วย Builtin func utf8.RuneCountInString(a) ได้ ถูกต้อง")

	// short declaration
	x := 10
	//OR | var x int = 10
	// OR | var x = 10

	y := true
	// var y bool

	z := "Hello World"
	// var z string

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x)

	println()
	println()
	// ----------------- If-Else condition ------------------------------------
	point := 50

	if point >= 50 {

		println("Passes")

	} else if point < 50 {

		println("Not Passes")

	} else {

		println("5555")

	}

	// --------------------- Array -----------------------------------------------
	// array fix ความยาว
	var c [3]int = [3]int{10, 20, 30}
	var d [3]bool = [3]bool{}

	// or

	e := [...]int{10, 20, 30, 40, 50}

	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%#v", e)
	println()
	fmt.Printf("%v", e)
	println()

	// -------------------------- Slide --------------------------------------
	//go lang เรียก slide ภาษาอื่น เรียก list
	// slide ไม่ fix ความยาว
	// index    0  1  2  3  4
	f := []int{10, 20, 30, 40, 50}
	// เพิ่มข้อมูลต่อท้าย slide
	f = append(f, 60)
	// index  0  1  2  3  4  5
	// []int{10,20,30,40,50,60}
	g := f

	fmt.Printf("%v", f)
	println()
	fmt.Printf("%v", g)
	println()

	// sub slide ตั้งแต่ index 3 to ตัวสุดท้ายของ slide
	fmt.Printf("%v", f[3:])
	println()

	// sub slide ตั้งแต่ index 3 to 4 (index ที่ต้องการจะ +1 ดังนั้น ต้องการ index 4 ต้อง +1 จะเป็น 5 ) (slideName[ ไม่ต้อง +1 : +1 ])
	fmt.Printf("%v", f[3:5])
	println()

	// sub slide ตั้งแต่ index 0 to 3
	fmt.Printf("%v", f[:3])
	println()

	// ------------------------- Map ------------------------------------------------
	// การใช้ map ภาษาอื่น เรียก dictionary (การใช้งาน แบบ Key, value)
	// ประกาศตัวแปรแบบเต็ม
	// var h map[string]string = map[string]string{}

	// short declaration
	h := map[string]string{}

	// key        value
	h["th"] = "Thailand"
	h["us"] = "United States"

	println("th", h["th"])
	println("us", h["us"])

	// รับค่า map แล้วเช็คว่าค่ามีอยู่จริงหรือไม่ พร้อมกับ Assignment ค่าให้ตัวแปรอีกตัวด้วย
	// map จะ return two pole ค่าให้ตัวแปรที่รอรับฝั่งซ้ายมือ 2 ค่า (1 ค่าข้อมูลที่ต้องการ, 2 ค่าที่ขอมีน้อมูลจริงหรือไม่ ตัวแปรนั้นชื่อ ok มีชนิดเป็น bool)

	// กรณีที่-มี-ค่า "th" ใน map  ok = true
	i, ok := h["th"]

	if ok {
		println(i)
	} else {
		println("No value")
	}

	// กรณีที่-ไม่มี-ค่า "en" ใน map ok = false
	j, ok := h["en"]

	if ok {
		println(j)
	} else {
		println("No value")
	}

	// ----------------------------- For Loop -----------------------------------------

	sum := 0
	for i := 0; i <= 5; i++ {
		sum += 1
	}
	println("Sum =", sum)

	// for each
	sum2 := 0
	k := []int{1, 2, 3, 4, 5}

	for j := 0; j < len(k); j++ {
		sum2 += k[j]
	}
	println("Sum2 =", sum2)

	// for each (range) อีกแบบ จะ return two pole values มา 2 ตัว (1 ค่าของ index, 2 ค่าของ value)
	sum3 := 0
	sumIndex := 0
	for index, v := range k {
		sum3 += v
		sumIndex += index
	}
	println("Index = ", sumIndex, " Sum2 = ", sum3)

	// กรณีที่ loop แบบ for each (range) จะ return two pole values มา 2 ตัว (1 ค่าของ index, 2 ค่าของ value) แต่ไม่ต้องการ index ต้องการแค่ค่า value ของ slide อย่างเดียว
	// สามารถทำการ ignore ตัว index ได้ (สามารถ ignore ได้ทั้ง index และ value แล้วแต่ว่าเราต้องการอะไร)
	// ignore โดยการใส่ เครื่องหมาย _ เอาไว้ ( _ คือ Blank Declaration )
	sum4 := 0
	for _, v := range k {
		sum4 += v
	}
	println(" Sum4 = ", sum4)

	// ------------------ Pointer ------------------------------------------------------------------------------------------------

	l := 10
	var m int = 10

	// ดึงตำแหน่ง address ของ l ออกมาดู
	println(&l)
	println("ตำแหน่ง m ใน memory ", &m)

	// ดูค่า
	println(l)
	println(m)

	// ให้ n เป็น pointer เก็บตำแหน่งของ int
	// และชี้ไปยังตำแหน่งของ m
	// var n *int
	// ให้ n เก็บตำแหน่งของ m
	n := &m
	println("เอาข้อมูล value ออกมาดู ", *n)
	println("เอาตำแหน่งของตัวเองใน memory ออกมาดู ", &n)
	println("เอาตำแหน่งของ m ใน memory ที่ n ชี้ referace ออกมาดู", n)

	// ในกรณีที่ เราเปลี่ยนค่าของ n หรือ m ค่าข้อมูลจะเปลี่ยนทั้งคู่ เพราะ n และ m อ้างถึงตำแหน่งเก็บข้อมูลเดียวกัน
	*n = 25
	println("หลังจากเปลี่ยนค่า n เอาค่า n ออกมาดู ", *n)
	println("หลังจากเปลี่ยนค่า n เอาค่า m ออกมาดู ", m)

	/// -------------------- struct ---------------------------------------------------------------
	// มีแต่ feild ไม่มี function
	// ถ้าเขียนแบบบรรทัดเดียวไม่จำเป็นต้องมี (,) ทุกตัว
	// st := Person{Name: "charan", Surname: "detkrathok"}

	// แต่ถ้าจะเขียนในรูปแบบขึ้นบรรทัดใหม่แบบด่านล่างนี้ ตัวสุดท้ายต้องมี (,) ปิดท้ายเสมอ ( .... Surname: "detkrathok",} )
	st := person.Person{}

	st.SetName("charan")
	st.SetSurname("detkrathok")

	println("STRUCT = ", st.GetName(), st.GetSurname())
	fmt.Printf("%#v", st.GetName())
	println()

	// ------------------------- function -------------------------------------------------------
	helloeWorld := hello("World")
	println(helloeWorld)

	result := funcName("Charan", "Detkrathok")
	println(result)

	o := []int{10, 20, 30}
	resultSum := sumNumber(o)
	println(resultSum)

	resultSumVariadic := sumVariadic(10, 20, 30)
	println(resultSumVariadic)

	// Anonymous function
	// var p func(int, int) int
	/* Or
	p := func(a, b int) int {
		return a + b
	}
	*/

	p := sumAnonymous
	println(p(10, 20))

	// ทำการส่ง function เข้าไปใน call func
	call(add)
	call(sub)

	// ----------------------------- เรียกใช้งานต่าง package -------------------------------------
	// เรียก function Hello จาก package customer
	// โดย function Hello กำหนดตัวเเรกเป็น - ตัวใหญ่ H - เพื่อทำการ expose function ให้ภายนอก package customer เรียกใช้ได้

	customer.OtherName = "charan detkrathok"
	callCustomerPackage := customer.Hello()
	println(callCustomerPackage)

}

// ------------------------- function -------------------------------------------------------

// การ return การบอกถึงการ return คือ ดูที่ด้านหลังว่ามี return ค่าเป็นแบบไหน เช่น return string
func hello(s string) string {
	return "func Hello " + s
}

// short hand funcName(name, surname string)
// แบบเต็ม funcName(name string, surname string)
func funcName(name, surname string) string {
	return name + " " + surname
}

// slide
func sumNumber(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}

	return sum
}

// Variadic function
func sumVariadic(a ...int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}

	return sum
}

func sumAnonymous(a, b int) int {
	return a + b
}

func call(f func(int, int) int) {
	sum := f(50, 10)
	println(sum)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
