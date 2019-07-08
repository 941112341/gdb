package learn

type Comparable interface {
	Compare(other interface{}) int
}

func Lt(c1, c2 Comparable) bool {
	return c1.Compare(c2) < 0
}

func Lte(c1, c2 Comparable) bool {
	return c1.Compare(c2) <= 0
}

func Gt(c1, c2 Comparable) bool {
	return c1.Compare(c2) > 0
}

func Gte(c1, c2 Comparable) bool {
	return c1.Compare(c2) >= 0
}

func Eq(c1, c2 Comparable) bool {
	return c1.Compare(c2) == 0
}
