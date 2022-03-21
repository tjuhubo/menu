package main
import "fmt"

func main()  {
	fmt.Println("input your command:")
	fmt.Println("1.help 2.list 3.quit")
	var cmd string
	for{
		fmt.Scanln(&cmd)
		if(cmd=="quit"){
			break;
		}
		if(cmd=="help"){
			fmt.Println("I will help you!")
		}
		if(cmd=="list"){
			fmt.Println("list all files")
		}
	}
	fmt.Println("Bye Bye")
}
