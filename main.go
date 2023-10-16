package main

import "josephu/tools"

func main() {
	var n = 5
	var startId = 1
	var killNum = 1
	head := &tools.Node{} //初始化头结点

	for i := 1; i <= n; i++ { //初始化环
		hero := &tools.Node{
			Id: i,
		}
		tools.InsertNode(head, hero)
	}
	tools.ListNode(head)

	tools.DelNode(head, startId, killNum)
}
