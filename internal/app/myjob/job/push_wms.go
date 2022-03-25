package job

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/opentracing/opentracing-go"
)

// PushWmsObj 推送单据到wms
func (j *Job) PushWmsObj(args interface{})  (err error,taskId string) {
	var data = make(map[string]interface{})
	b := args.([]byte)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	decoder.Decode(&data)
	taskId = data["traceId"].(string)
	j.setInfo(taskId, "pushWmsObj")
	j.debug(data["no"])
	j.info("原数据：",data)
	j.debug(data["traceId"])

	//j.info(j.svr.GetName())
	ctx:=context.Background()
	addOrderSpan,nextCtx := opentracing.StartSpanFromContext(ctx, "push")
	defer addOrderSpan.Finish()
	addOrderSpan.SetTag("add.order start", true)
	addOrderSpan.LogKV("data", data)
	j.svr.GetName(nextCtx)
	return
}
