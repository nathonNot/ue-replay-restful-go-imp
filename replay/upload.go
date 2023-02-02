package replay

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"strings"
)

type FNetworkReplayStartUploadingResponse struct {
	SessionId string `json:"sessionId"`
}

// BeginReplay get session name
func BeginReplay(ctx iris.Context) {
	sessionName := ctx.Params().Get("session_name")
	ctx.JSON(FNetworkReplayStartUploadingResponse{
		SessionId: sessionName,
	})
}

// UploadHeader 开始上传
func UploadHeader(ctx iris.Context) {
	sessionName := ctx.Params().Get("session_name")
	fileName := ctx.Params().Get("file_name")
	//numChunks, _ := ctx.URLParamInt("numChunks")
	//endTimes, _ := ctx.URLParamInt64("time")           // 结尾时间戳,毫秒
	//mTime1, _ := ctx.URLParamInt64("mTime1")           // 当前块开始时间戳,毫秒
	//mTime2, _ := ctx.URLParamInt64("mTime2")           // 当前块结束时间戳,毫秒
	//absSize, _ := ctx.URLParamInt64("absSize")         // 当前块大小
	if strings.Contains(ctx.Path(), "replay.header") { // 头块
		//absSize = 0 // 当前块大小
	}
	if _, err := os.Stat(sessionName); os.IsNotExist(err) {
		err := os.Mkdir(sessionName, os.ModePerm)
		if err != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			return
		}
	}
	if fileName == "replay.header" {
		//numChunks = -1
	}
	savePath := fmt.Sprintf("/replay/%s/%s", sessionName, fileName)
	out, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_, err = io.Copy(out, ctx.Request().Body)
	if err != nil {
		return
	}
	defer out.Close()
	ctx.StatusCode(iris.StatusNoContent)
}

func AddEvent(ctx iris.Context) {
	sessionName := ctx.Params().Get("session_name")
	//group := ctx.URLParam("group")
	//time1, _ := ctx.URLParamInt64("time1")
	//time2, _ := ctx.URLParamInt64("time2")
	meta := ctx.URLParam("meta")
	//incrementSize, _ := ctx.URLParamBool("incrementSize")
	//roomDir := fmt.Sprintf("/replay/%d", roomId)
	if _, err := os.Stat(sessionName); os.IsNotExist(err) {
		err := os.Mkdir(sessionName, os.ModePerm)
		if err != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			return
		}
	}
	savePath := fmt.Sprintf("/replay/%s/%s", sessionName, meta)
	out, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_, err = io.Copy(out, ctx.Request().Body)
	if err != nil {
		return
	}
	defer out.Close()
	ctx.StatusCode(iris.StatusNoContent)
}

func UpdateEvent(ctx iris.Context) {
	eventFullName := ctx.Params().Get("event_full_name")
	nameList := strings.Split(eventFullName, "_")
	sessionName := nameList[0]
	//eventName := nameList[1]
	//group := ctx.URLParam("group")
	//time1, _ := ctx.URLParamInt64("time1")
	//time2, _ := ctx.URLParamInt64("time2")
	//meta := ctx.URLParam("meta")
	//incrementSize, _ := ctx.URLParamBool("incrementSize")
	roomDir := fmt.Sprintf("/replay/%s/%s", sessionName, eventFullName)
	if _, err := os.Stat(roomDir); os.IsNotExist(err) {
		err := os.Mkdir(roomDir, os.ModePerm)
		if err != nil {
			ctx.StatusCode(iris.StatusUnauthorized)
			return
		}
	}
	savePath := fmt.Sprintf("/replay/%s/%s", sessionName, eventFullName)
	out, err := os.OpenFile(savePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	_, err = io.Copy(out, ctx.Request().Body)
	if err != nil {
		return
	}
	defer out.Close()
	ctx.StatusCode(iris.StatusNoContent)
}

type SSessionUserList struct {
	Users []string `json:"users"`
}

func PostUser(ctx iris.Context) {
	//sessionName := ctx.Params().Get("session_name")
	var req SSessionUserList
	ctx.ReadJSON(&req)
	//log.ServerLog().Infof("roomId:%d,session name:%s,post user:%v", roomId, sessionName, req.Users)
}

func StopStreaming(ctx iris.Context) {
	//sessionName := ctx.Params().Get("session_name")
	//endChunks, _ := ctx.URLParamInt("numChunks") // 结尾包
	//endTimes, _ := ctx.URLParamInt64("time")     // 结尾时间戳,毫秒
	//absSize, _ := ctx.URLParamInt("absSize")     // 文件大小
	ctx.StatusCode(iris.StatusNoContent)
}
