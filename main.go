package main

import (
	"github.com/kataras/iris/v12"
	"github.com/nathonNot/ue-replay-restful-go-imp/replay"
)

func InitApi() {
	app := iris.New()
	replayRoute := app.Party(ReplayRoute)
	{
		replayRoute.Post(BeginReplay, replay.BeginReplay)
		replayRoute.Post(UploadHeader, replay.UploadHeader)
		replayRoute.Post(AddEvent, replay.AddEvent)
		replayRoute.Post(UpdateEvent, replay.UpdateEvent)
		replayRoute.Post(StopStreaming, replay.StopStreaming)
		replayRoute.Post(PostUser, replay.PostUser)
	}
	app.Listen("0.0.0.0:5001")
}

func main() {
	InitApi()
}
