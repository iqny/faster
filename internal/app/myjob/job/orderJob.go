package job

import (
	"bytes"
	"encoding/gob"
)

func (j *Job) OrderJob(args interface{},taskId string)error  {
	j.setInfo(taskId,"orderJob")
	var scb string
	b := args.([]byte)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	decoder.Decode(&scb)
	j.debug(scb)
	j.info(scb)
	return nil
}
