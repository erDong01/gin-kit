package betree

import (
	"github.com/erdong01/kit/base/vector"
	"sort"
)

type (
	BehaviorList struct {
		vector.Vector[IBaseNode]
	}

	IBehaviorList interface {
		AddChild(string, IBaseNode)
		DelChild(string)
		GetChild(string) IBaseNode
		GetChildCount() int
	}

	BehaviorTree struct {
		BehaviorList
	}

	IBehaviorTree interface {
		IBehaviorList
		Init()
		OnExec(int64)
	}
)

func (b *BehaviorList) Less(i, j int) bool {
	return b.Get(i).GetName() < b.Get(j).GetName()
}

func (b *BehaviorList) AddChild(name string, pNode IBaseNode) {
	pNode.Init()
	if pNode.GetType() != COMPOSITE && pNode.GetType() != DECORATOR && pNode.GetType() != ACTION && pNode.GetType() != CONDITION {
		return
	}

	pCurNode := b.GetChild(name)
	if pCurNode != nil {
		if pCurNode.GetType() == COMPOSITE && pNode.(IComposite) != nil {
			pCurNode.(IComposite).AddChild(name, pNode)
		} else if pCurNode.GetType() == DECORATOR && pNode.(IDecorator) != nil {
			pCurNode.(IDecorator).SetChild(pNode)
		}
	} else {
		pNode.SetName(name)
		b.PushFront(pNode)
		sort.Sort(b)
	}
}

func (b *BehaviorList) DelChild(name string) {
	nIndex := sort.Search(b.Len(), func(i int) bool {
		return b.Get(i).(IBaseNode).GetName() >= name
	})
	if nIndex < b.Len() && b.Get(nIndex).(IBaseNode).GetName() == name {
		b.Erase(nIndex)
	}
}

func (b *BehaviorList) GetChild(name string) IBaseNode {
	nIndex := sort.Search(b.Len(), func(i int) bool {
		return b.Get(i).(IBaseNode).GetName() >= name
	})
	if nIndex < b.Len() && b.Get(nIndex).(IBaseNode).GetName() == name {
		return b.Get(nIndex).(IBaseNode)
	}
	return nil
}

func (b *BehaviorList) GetChildCount() int {
	return b.Len()
}

func (b *BehaviorTree) Init() {
}

func (b *BehaviorTree) OnExec(tick int64) {
	for _, v := range b.Values() {
		v.(IBaseNode).OnExec(tick)
	}
}

func NewBeTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Init()
	return tree
}
