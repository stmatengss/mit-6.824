package main
import "fmt"

type Camera struct{
}

type Phone struct{
}

func (c *Camera) TakePic() string {
		return "Click"
}

func (p *Phone) Call() string {
		return "Ring"
}

type CameraPhone struct{
		Camera
		Phone
}


func main () {
		cp := new(CameraPhone)
		fmt.Println(cp.TakePic())
		fmt.Println(cp.Call())
}
