package menureuse

import "fmt"
import "bufio"
import "strings"
import "os"

import(. "linkcallback")
import "unsafe"

const CMD_MAX_ARGV_NUM=32

var head *LinkTable 

func help(agrc int, argv [CMD_MAX_ARGV_NUM]string) int {
	ShowAllCmd(head)
	return 0
}


type DataNode struct {
	pNext *LinkTable
	cmd string
	desc string
	handler func(argc int, argv [CMD_MAX_ARGV_NUM]string) int
}

func SearchCondition(pLinkTableNode *LinkTableNode, args unsafe.Pointer) int {
	pcmd := (*string)(unsafe.Pointer(args))
	dNode := (*DataNode)(unsafe.Pointer(pLinkTableNode))
	if dNode.cmd == *pcmd {
		return SUCCESS
	}
	return FAILURE
}

func FindCmd(head *LinkTable, cmd string) *DataNode {
	tNode:= SearchLinkTableNode(SearchCondition, head, unsafe.Pointer(&cmd))
	return (*DataNode)(unsafe.Pointer(tNode))
}

func ShowAllCmd(head *LinkTable) int {
	tNode := GetLinkTableHead(head)
	fmt.Printf("********Menu List:**********\n")
	for tNode != nil {
		dNode := (*DataNode)(unsafe.Pointer(tNode))
		fmt.Printf("%s - %s\n",dNode.cmd, dNode.desc)
		tNode = GetNextLinkTableNode(head,tNode)
	}
	fmt.Printf("****************************\n")
	return 0
}

func MenuConfig(cmd string, desc string, handler func(agrc int, agrv [CMD_MAX_ARGV_NUM]string) int) int {
	if head == nil {
		head = CreateLinkTable()
		var pNode = new(DataNode)
		pNode.cmd = "help"
		pNode.desc = "Show The Command List"
		pNode.handler=help
		AddLinkTableNode(head,(*LinkTableNode)(unsafe.Pointer(pNode)))
	}
	var pNode = new(DataNode)
	pNode.cmd = cmd
	pNode.desc = desc
	pNode.handler=handler
	AddLinkTableNode(head,(*LinkTableNode)(unsafe.Pointer(pNode)))
	return SUCCESS
}



func ExcuteMenu() {
	for true {
		inputReader := bufio.NewReader(os.Stdin)
		var argc int = 0
		var argv [CMD_MAX_ARGV_NUM]string
		fmt.Print("Command->")
		inputs, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("There ware errors reading, exiting program.")
 	        return
		}
		input := strings.Fields(inputs)
		for _, pcmd := range input {
			argv[argc] = pcmd
			argc +=1
		}
		p := FindCmd(head,argv[0])
		if p == nil {
			fmt.Println("Command not exit,you can try help")
			continue
		}
		fmt.Printf("%s - %s\n",p.cmd,p.desc)
		if p.handler != nil {
			p.handler(argc, argv)
		}
	}
}