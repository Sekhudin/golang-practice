package main

import "fmt"

func HelloWorld() {
	fmt.Println("Hello Wolrd")
}

func exVariable() {
	/**
	 **Cara 1
	 * var VAR_NAME type = value
	 **Cara 2
	 *VAR_NAME := value
	 **Update Variable
	 * VAR_NAME = new value
	 */

	//  cara 1
	var firstName string = "John"
	// cara 2
	lastName := "Doe"
	// atau bisa juga seperti ini (tidak disarankan)
	// var nilai string
	// nilai = "90"

	fmt.Println(firstName, lastName)

	// update
	firstName = "Jean"
	lastName = "Do"
	fmt.Println("update>>", firstName, lastName)
}

func exConstant() {
	/**
	 **Single Constant
	 * cont CONST_NAME type = value
	 **Multiple Constant
	 * const (...CONST_NAME type = value)
	 */

	// single Constant
	const single string = "You!"
	// bisa juga gini
	const alternatif = "alternatif"
	// multiple Constant
	const (
		fistName string = "John"
		lastName string = "doe"
		// bisa juga gini
		fullName = "John Doe"
	)

	fmt.Println(
		single,
		alternatif,
		fistName,
		lastName,
		fullName,
	)
}

func exArray() {
	/**
	 ** Cara 1
	 * var ARRAY_NAME = [panjang] type{ ...values }
	 ** Cara 2
	 * var ARRAY_NAME = [...]type{ ...values }
	 **Cara 3
	 * ARRAY_NAME := [panjang]type{...values}
	 ** Update item array
	 * ARRAY_NAME[index] = new value
	 */

	//  cara 1
	var arr1 = [2]string{"John", "doe"}
	//  cara 2
	var arr2 = [...]string{"john", "Doe"}
	//  cara 3
	arr3 := [2]string{"strong", "man"}
	// atau dengan cara ini
	var arr4 [1]string
	arr4[0] = "hello"

	fmt.Println(arr1, arr2, arr3, arr4)
	// update
	arr1[0] = "John-new"
	fmt.Println(arr1, arr2)
}

func exSlices() {
	/**
	 **Cara 1
	 * var SLICE_NAME = []type{...values}
	 **Cara 2
	 * SLICE_NAME := []type{..values}
	 **Cara 3 (Membuat slice dari array yang ada)
	 * var SLICE_NAME []type = ARRAY_NAME[low:high]
	 * var SLICE_NAME []type = ARRAY_NAME[low:]
	 * var SLICE_NAME []type = ARRAY_NAME[:high]
	 * var SLICE_NAME []type = ARRAY_NAME[:]
	 * *Cara 4 (Membuat slice menggunakan fungsi make())
	 * SLICE_NAME:= make([]type, panjang, kapasitas)
	 */
	// Cara 1
	var slice1 = []int{1, 2, 3}
	// Cara 2
	slice2 := []int{1, 2, 3}
	// Cara 3
	arr := [3]string{"John", "Jean", "Doe"}
	slices3 := arr[:]
	slices4 := arr[2:3]
	slices5 := arr[1:]
	slices6 := arr[:]
	// cara 4 (menggunakan fungsi make())
	slice7 := make([]string, 2, 5)
	// -- mengisi array dari slice
	slice7[0] = "John"
	slice7[1] = "Doe"
	// untuk menambahkan item ke slice menggunakna append()

	//slice7[3] = "Slice7" // error --> panjang maksmimal 2 item

	fmt.Println(
		slice1,
		slice2,
		slices3,
		slices4,
		slices5,
		slices6,
		slice7,
	)
}

func exMap() {
	/**
	 **Cara 1
	 * MAP_NAME := map[type key]type value { ...key:value }
	 **Mengubah data
	 * MAP_NAME[key] = value
	 */

	john := map[string]string{
		"name": "john doe",
	}

	fmt.Println((john))
}

func main() {
	HelloWorld()
	exVariable()
	exConstant()
	exArray()
	exArray()
	exSlices()
	exMap()
}
