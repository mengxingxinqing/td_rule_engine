package node

import "rule/td_rule_engine/engine/core"

type ComputeNode struct {
	Name string
	Id string
	Type int
	BoolResult bool
	MapResult map[string]interface{}
	RunTimes int
}

func (node *ComputeNode) GetName() string{
	return node.Name
}

func (node *ComputeNode) GetBoolResult() bool{
	return node.BoolResult
}

func (node *ComputeNode) GetMapResult()map[string]interface{}{
	return node.MapResult
}
func (node *ComputeNode) GetId()string{
	return node.Id
}

func (node *ComputeNode) GetType() int{
	return node.Type
}

func (node *ComputeNode) Run(ctx *core.FlowContext){
	node.RunTimes += 1

}

