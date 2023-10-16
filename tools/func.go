package tools

import "fmt"

type Node struct {
	Id   int
	Next *Node //指向后一个节点
}

func InsertNode(head *Node, newNode *Node) {
	if head.Next == nil { //判断是不是第一个节点
		head.Id = newNode.Id
		head.Next = head //指向自己，形成环
	} else {
		temp := head //创建一个辅助节点
		for {        //找环尾
			if temp.Next == head {
				break
			}
			temp = temp.Next
		}
		temp.Next = newNode //将链表末尾指向新节点
		newNode.Next = head //将新节点指向环头
	}
}

func DelNode(head *Node, startId int, killNum int) {
	if head.Next == nil { //如果是空链
		fmt.Println("链表为空")
		return
	}

	if startId > 1 {
		for i := 1; i < startId-1; i++ { //把环头移动到要开始的节点前一个
			head = head.Next
		}
	}
	if startId == 1 { //根据后面的删除算法，这时需要把环头前移一位
		temp := head //需要辅助节点
		for {
			if head.Next == temp {
				break
			}
			head = head.Next
		}
	}
	
	for {
		for i := 1; i < killNum; i++ { //把环头移动到要淘汰的节点前一个
			head = head.Next
		}
		fmt.Printf("%d淘汰\n", head.Next.Id)
		head.Next = head.Next.Next //把环头指向下个节点的下个节点，淘汰下个节点
		if head.Next == head {     //当环只剩一个节点时
			fmt.Printf("%d幸存\n", head.Id)
			break
		}
	}
}

func ListNode(head *Node) {
	temp := head
	if temp.Next == nil {
		fmt.Println("链表为空")
		return
	}
	for {
		fmt.Printf("[%d]->", temp.Id) //输出当前节点
		if temp.Next == head {        //先判断
			break
		}
		temp = temp.Next //再跳转
	}
	fmt.Println(head.Id)
}
