package SendPriorityMessageToAssets

import (
	"github.com/AiRISTAFlowInc/fs-core/data/coerce"
)

type Input struct {
	MessageOptions string `md:"MessageOptions,allowed(Beep,Twinkle,Stairs,Siren,Silent)"`
	MessageExpiration          int `md:"MessageExpiration"`  
	IP          string `md:"IP,required"`
	Message     string `md:"Message,required"`
	Password    string `md:"Password,required"`
	StaffIdList string `md:"StaffIdList,required"`
	Username    string `md:"Username,required"`
}

func (i *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["MessageOptions"])
	i.MessageOptions = strVal

	intVal, _ := coerce.ToInt(values["MessageExpiration"])
	i.MessageExpiration = intVal

	strVal, _ = coerce.ToString(values["IP"])
	i.IP = strVal

	strVal, _ = coerce.ToString(values["Username"])
	i.Username = strVal

	strVal, _ = coerce.ToString(values["Password"])
	i.Password = strVal

	strVal, _ = coerce.ToString(values["StaffIdList"])
	i.StaffIdList = strVal

	strVal, _ = coerce.ToString(values["Message"])
	i.Message = strVal

	return nil
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"MessageOptions":  i.MessageOptions,
		"MessageExpiration":    i.MessageExpiration,
		"IP":          i.IP,
		"Username":    i.Username,
		"Password":    i.Password,
		"StaffIdList": i.StaffIdList,
		"Message":     i.Message,
	}
}

type Output struct {
	Status string `md:"Status"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	Status, ok := (values["Status"]).(string) // type assertion
	if ok {
		o.Status = Status
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Status": o.Status,
	}
}

type Asset struct {
	ID int `json:"Id"`
}

type Login struct {
	CustomerID  int     `json:"CustomerId"`
}