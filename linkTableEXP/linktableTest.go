package main

import (
	"fmt";
)

//链表节点
type node struct {
	val int
	pNext  *node
}
//链表
type linktable struct {
	pHead *node
	pTail *node
	SumOfNode int
}

// 创建一个节点
func newNode(val int) (pNode *node) {
    pNode = &node{
        val: val,
    }

    return pNode
}

//创建一个链表
func createLinkTable() *linktable {
	var pLinkTable *linktable = new(linktable)
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	return pLinkTable
}

//删除链表
func deleteLinkTable(pLinkTable *linktable) *linktable {
	pLinkTable.pHead = nil
	pLinkTable.pTail = nil
	pLinkTable.SumOfNode = 0
	return nil
}

//在链表添加节点
func (pLinkTable *linktable) addNode(pNode *node) *linktable {
	if pLinkTable==nil||pNode==nil {
		fmt.Println("链表或者节点为空")
		return pLinkTable
	}
	//若链表为空则插入第一个节点
	if pLinkTable.pHead==nil && pLinkTable.pTail==nil {
		pLinkTable.pHead = pNode
		pLinkTable.pTail = pNode
	} else {
		pLinkTable.pTail.pNext = pNode
		pLinkTable.pTail = pNode
	}
	pLinkTable.pTail.pNext = nil
	pLinkTable.SumOfNode += 1
	return nil
}

//在链表删除节点
func (pLinkTable *linktable) deleteNode(pNode *node) *linktable {
	if pLinkTable==nil {
		fmt.Println("链表为空无法删除")
		return pLinkTable
	}
	var pNow *node = pLinkTable.pHead
	var preNode *node = pNow
	if pLinkTable.pHead == pNode {
		pLinkTable.pHead = pNow.pNext
		pLinkTable.SumOfNode--
		return nil
	}
	pNow = pNow.pNext
	for pNow != nil {
		if pNow == pNode {
			preNode.pNext = pNow.pNext
			pLinkTable.SumOfNode--
			return nil
		}
		preNode = pNow
		pNow = pNow.pNext
	}
	fmt.Println("节点不存在")
	return pLinkTable
}
//查找指定节点
func (pLinkTable *linktable) findNode(pNode *node) *node {
	var nowNode *node = pLinkTable.pHead
	for nowNode != nil {
		if nowNode == pNode {
			return nowNode
		}
		nowNode = nowNode.pNext
	}
	fmt.Println("链表中不存在该节点")
	return nil
}
//获取下一个节点
func (pLinkTable *linktable) getNext(pNode *node) *node {
	var pNow *node = pLinkTable.pHead
	for pNow != nil {
		if pNow == pNode {
			return pNow.pNext
		}
		pNow = pNow.pNext
	}
	return nil
}

func main()  {
	linktable := createLinkTable()
	node1 := newNode(1)
	linktable.addNode(node1)
	fmt.Println(linktable.findNode(node1))
	fmt.Println("插入节点1并且查找节点1结束")
	node2 := newNode(2)
	linktable.addNode(node2)
	fmt.Println(linktable.getNext(node1))
	fmt.Println("获取node1的下一个节点node2成功")
	fmt.Println(linktable.findNode(node2))
	fmt.Println("插入节点2并且查找节点2结束")
	linktable.deleteNode(node2);
	fmt.Println("删除节点成功")
}