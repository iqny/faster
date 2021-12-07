package job

import (
	"bytes"
	"encoding/gob"
)

// PushWmsObj 推送单据到wms
func (j *Job) PushWmsObj(args interface{}, taskId string) error {
	j.setInfo(taskId, "pushWmsObj")
	var data = make(map[string]interface{})
	b := args.([]byte)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	decoder.Decode(&data)
	j.debug(data["no"])
	j.info(data)
	return nil
}
