package job

import (
	"bytes"
	"encoding/gob"
)

func (j *Job) TestJob(args interface{}) (err error,taskId string)  {
	j.setInfo("","testJob")
	var scb string
	b := args.([]byte)
	decoder := gob.NewDecoder(bytes.NewReader(b))
	decoder.Decode(&scb)
	j.debug(scb)
	return nil,""
}
