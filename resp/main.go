package main
import (
	"fmt"
	"os"
)

func main(){
	args:= os.Args
	fmt.Printf("args= %s", args[len(args)-1:])
}