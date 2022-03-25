package job

import (
	"bytes"
	"encoding/gob"
)

func (j *Job) OrderJob(args interface{}) (err error,traceId string)  {
	j.setInfo("","orderJob")
	var scb string
	b := args.([]byte)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	decoder.Decode(&scb)
	j.debug(scb)
	j.info(scb)
	return nil,""
}
