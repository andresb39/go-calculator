package calculator

import (
	"fmt"
)

type Front struct{}

func (f Front) Wellcome() {
	banner := `

██╗    ██╗███████╗██╗     ██╗      ██████╗ ██████╗ ███╗   ███╗███████╗    ████████╗ ██████╗ 
██║    ██║██╔════╝██║     ██║     ██╔════╝██╔═══██╗████╗ ████║██╔════╝    ╚══██╔══╝██╔═══██╗
██║ █╗ ██║█████╗  ██║     ██║     ██║     ██║   ██║██╔████╔██║█████╗         ██║   ██║   ██║
██║███╗██║██╔══╝  ██║     ██║     ██║     ██║   ██║██║╚██╔╝██║██╔══╝         ██║   ██║   ██║
╚███╔███╔╝███████╗███████╗███████╗╚██████╗╚██████╔╝██║ ╚═╝ ██║███████╗       ██║   ╚██████╔╝
 ╚══╝╚══╝ ╚══════╝╚══════╝╚══════╝ ╚═════╝ ╚═════╝ ╚═╝     ╚═╝╚══════╝       ╚═╝    ╚═════╝ 
                         ██████╗ █████╗ ██╗      ██████╗                                    
                        ██╔════╝██╔══██╗██║     ██╔════╝                                    
                        ██║     ███████║██║     ██║                                         
                        ██║     ██╔══██║██║     ██║                                         
                        ╚██████╗██║  ██║███████╗╚██████╗                                    
                         ╚═════╝╚═╝  ╚═╝╚══════╝ ╚═════╝                                    

`
	fmt.Println(banner)
}
func (f Front) Menu() {
	f.Wellcome()
	fmt.Println("Select the operation you wish to perform:")
	fmt.Println("1 -> Sum")
	fmt.Println("2 -> Subtract")
	fmt.Println("3 -> Multiply")
	fmt.Println("4 -> Divide")
	fmt.Println("Other Quit")
}

func (f Front) App() {
	option := ReadInput()
	c := Calc{}
	c.Calculate(option)
}
