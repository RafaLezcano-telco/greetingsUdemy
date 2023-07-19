package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)
func Hello(name string)(string,error){
	if(name==""){
		return "", errors.New("nombre vacio")
	}
	message:= fmt.Sprintf(randomFormat(),name)
	return message,nil
}

func Hellos(names []string)(map [string]string,error){
	messages:=make(map[string]string)

	for _,name := range names{
		message,err:= Hello(name)
		if err != nil {
			return nil, err 
		}
		messages[name]=message 
	}
	return messages,nil
}




func randomFormat()string{
	formats:=[] string{
		"hola %v, bienvenido",
		"Bueno verte, %v",
		"Saludo, %v",
	}

	//Genero un numero aleatorio
	randomIndex:=rand.Intn(len(formats))
	fmt.Println(randomIndex)
	return formats[randomIndex]
}