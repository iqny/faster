package job

import (
	"bytes"
	"encoding/gob"
)

func (j *Job) OrderTransformJob(args interface{}) (err error,taskId string)  {
	j.setInfo("","orderTransformJob")
	var scb string
	b := args.([]byte)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	decoder.Decode(&scb)
	j.debug(scb)
	j.info(scb)
	return nil,""
}
