package flow

import "encoding/json"


func InitFlow(config string) {
	flowData := &FlowData{}
	err := json.Unmarshal([]byte(config), flowData)
	if err!=nil{
	//	todo::处理异常
	}
	for node := range flowData.NodeList{

	}
}
