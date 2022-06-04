package boolean

type Boolean interface {
	~bool
}

func IsTrue[T Boolean](b T) bool {
	return b == true
}

type Pointer interface {
	~*bool
}

func IsPTrue[T Pointer](b T) bool {
	if b != nil {
		return *b
	}
	return false
}
