package main

import (
	"github.com/emicklei/go-restful"
	"io"
	"net/http"
	"os"
	"os/exec"
)

var port string = ":8075"

func main() {
	ws := new(restful.WebService)
	ws.Path("/balr")
	ws.Route(ws.POST("/maintenance").To(createMaintenance))
	ws.Route(ws.DELETE("/maintenance").To(removeMaintenance))
	ws.Route(ws.GET("/maintenance").To(statusMaintenance))
	restful.Add(ws)
	http.ListenAndServe(port, nil)
}

func removeMaintenance(req *restful.Request, resp *restful.Response) {
	cmd := exec.Command("rm", "/tmp/maintenance.flag")
	out, err := cmd.Output()
	if err != nil {
		io.WriteString(resp, err.Error())
		return
	}
	io.WriteString(resp, string(out))
}

func createMaintenance(req *restful.Request, resp *restful.Response) {
	cmd := exec.Command("touch", "/tmp/maintenance.flag")
	out, err := cmd.Output()
	if err != nil {
		io.WriteString(resp, err.Error())
		return
	}
	io.WriteString(resp, string(out))
}

func statusMaintenance(req *restful.Request, resp *restful.Response) {
	if _, err := os.Stat("/tmp/maintenance.flag"); err == nil {
		io.WriteString(resp, "Maintenance mode is currently enabled")
	} else {
		io.WriteString(resp, "Maintenance mode is currently NOT enabled")
	}
}
