package main 


import (
	"os/exec"
	"os"
	"github.com/fatih/color"
	"strings"
	"fmt"
	"bytes"
	"reflect"
	"unsafe"
)


func RunCmdSmart(cmd string) (string,error) {
	 parts := strings.Fields(cmd)
  //	fmt.Println(parts[0],parts[1:])
    var out *exec.Cmd
  
    out = exec.Command(parts[0],parts[1:]...)	
   
    
	var ou ,our bytes.Buffer
	out.Stdout = &ou
	out.Stderr = &our

	fmt.Println(BytesToString(our.Bytes()))
	err := out.Run()
	if err != nil {
	//	fmt.Println("%v", err.Error())
		return our.String(), err
	}
	return ou.String(),nil
}
func BytesToString(b []byte) string {
				    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
				    sh := reflect.StringHeader{bh.Data, bh.Len}
				    return *(*string)(unsafe.Pointer(&sh))
				}

func main() {
	fil := ""
	 if len(os.Args) > 1 {

	 	fil = strings.Join(os.Args[1:], " " )
	 } else {
	 		fil= ""
	 }
	 	
	 	 log_build,err := RunCmdSmart("go run " + fil)
						 if err != nil {
						 	//fmt.Println(err.Error())
						  color.Red("Your build failed, Here is why :>")
						 	lines := strings.Split(log_build,"\n")
						 	 for i, line := range lines {
						 	 	if i > 0 {
						        if strings.Contains(line,"imported and") {
						        	line_part  := strings.Split(line,":")
						        	color.Red(strings.Join(line_part[2:]," - ") + " in file " + line_part[0])
						        } else {
						        	if line != "" {
						           line_part  := strings.Split(line,":")
						          
						         
								    
								    //fmt.Println(line_part[len(line_part) - 1])
								 
							
								  	color.Magenta("Verify your file " + line_part[0] + " on line : " + line_part[1] + " | " + strings.Join(line_part[2:]," - "))

								   

								  
						         
						         	
						        }
						    }
						    	}
						    }
						    color.Red("Full compiler build log : ")
						    fmt.Println(log_build)
						    return
						 } else {
						 	color.Green("Success! run with :")
						 	//fmt.Print(log_build)
						 	color.Magenta("go run " + fil)
						
						 }
	
}