package services

type RedisDM[T comparable] interface {
	ChangeFirstHash()
	ChangeSecondHash()
	AddSpace()
	DownSpace()
	AddEle(T)
	GetP()
}
