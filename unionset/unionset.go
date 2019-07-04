package unionset

type UnionSet struct {
	Parent  *UnionSet
	element interface{}
}

func (Self *UnionSet) IsSibling(other *UnionSet) bool {
	return Self.Root().element == other.Root().element
}

func (Self *UnionSet) Merge(Other *UnionSet) (self *UnionSet, other *UnionSet) {
	Other.Root().Parent = Self.Root() // 更优化的手段是进行对比，进行加权
	return Self, Other
}

func (Self *UnionSet) Root() *UnionSet {
	r := Self
	for r.Parent != nil {
		r = r.Parent
	}
	// 路径压缩
	self := Self
	for self != r {
		sParent := self.Parent
		self.Parent = r
		self = sParent
	}
	return r
}

func New(data interface{}) *UnionSet {
	return &UnionSet{
		element: data,
	}
}

func (Self *UnionSet) Union(data interface{}) (self *UnionSet, other *UnionSet) {
	return Self.Merge(New(data))
}
