package entities

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user   // 포함된 타입을 비노출 타입으로 선언
	Rights int
}
