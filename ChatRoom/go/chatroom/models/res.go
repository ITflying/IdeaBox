package models

type Res struct {
	Status string
    Msg string
	Result string
}

func InitResponseResultMsg(status, msg string) *Res  {
	res := new(Res)
	res.Status = status
	res.Msg = msg
	return res
}

func InitResponseResult(status, msg, result string) *Res  {
	res := new(Res)
	res.Status = status
	res.Msg = msg
	res.Result = result
	return res
}