package api

import (
	"camera/camera"
	"net/http"
)

type API struct {
	Filename string
	cam      camera.Camera
}

var cam camera.Camera

func (self *API) Init() error {
	cam = camera.Camera{"output.jpg"}
	self.cam = cam
	http.HandleFunc("/get", self.Get)

	server := http.Server{
		Addr: ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
