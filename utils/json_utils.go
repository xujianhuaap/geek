package utils

import "encoding/json"

func ToJson(obj interface{})(string,error)  {
	objBytes,err:=json.Marshal(obj)
	if(err!=nil){
		return "",err
	}
	userStr:=string(objBytes);
	return userStr,nil
}
