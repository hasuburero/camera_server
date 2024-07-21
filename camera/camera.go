package camera

import (
	"fmt"
	"os/exec"
)

type Camera struct {
	Filename string
}

func (self *Camera) Get() error {
	//self.filename = "output.jpg"
	cmd := exec.Command("libcamera-jpeg", "-t", "1", "-o", self.Filename)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Println("capture succeeded")
	return nil
}
