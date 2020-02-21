package main

import "fmt"

type Human interface {
	//メソッド名だけを宣言（必ず使わなければならない）
	Say() string
}

type Person struct {
	Name string
}

type Dog struct {
	Name string
}

func (p *Person) Say() string{
	// struct書き換えでポインタ型を指定した場合
	p.Name = "Mr." + p.Name
	fmt.Println(p.Name)
	return p.Name
}

// Human（インターフェイス）を持ったDriveCarメソッド
func DriveCar(human Human){
	// mikeなら"Run"を出力、違うなら"Get out"を出力
	// stringで比べたいため、Sayの返り値はstringを設定する
	if human.Say() == "Mr.Mike"{
		fmt.Println("Run")
	}else{
		fmt.Println("Get out")
	}
}

func main() {
	// Human（インターフェイス）にstructを入れる
	// ポインタを扱う場合アドレスを渡さないといけないため、&を記述する
	var mike Human = &Person{"Mike"}
	var x Human = &Person{"x"}
	// var dog Dog = Dog{"dog"}
	DriveCar(mike)
	DriveCar(x)
	// 関数DriveCarはインターフェイスHumanで明示したメソッドSayを扱わないと動作しないため、
	// 、エラーが起きる。（インターフェイスのダックタイピング）
	// DriveCar(dog)
}