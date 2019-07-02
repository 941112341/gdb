package list

func (list *List) GetLastK(k int) interface{} {
	kNode := list.head
	eNode := kNode
	for i := 0; i < k; i++ {
		if eNode != nil {
			eNode = eNode.next
		} else {
			return nil
		}
	}
	return eNode.data
}