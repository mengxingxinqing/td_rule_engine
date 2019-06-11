package node

import "rule/td_rule_engine/engine/core"

type Node interface {
	GetName()
	GetBoolResult()
	GetMapResult()
	GetId()
	GetType() int
	Run()
}

func CreateNode(node Node) *Node{
	if node.GetType() == core.NodeTypeCompute {
		//newNode :=
	}
}
