package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"encoding/json"
)


type Welcome struct {
	Name string
	Time string
}

type JsonResponse struct{
	Value1 string `json:"key1"`
	Value2 string `json:"key2"`
	JsonNested JsonNested `json:"JsonNested"`
}
type JsonNested struct{
	NestedValue1 string `json:"nestedkey1"`
	NestedValue2 string `json:"nestedkey2"`
}
type JsonName struct{
	FName string `json:"First Name"`
	LName string `json:"Last Name"`
	JsonCont JsonCont `json:"Contact"`
	JsonInfo JsonInfo `json:"About"`
}
type JsonCont struct{
	Num string `json:"Phone-No"`
	Add string `json:"Address"`
}
type JsonInfo struct{
	Chrom string `json:"Gender"`
	Status string `json:"Student"`
}

func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))
	
	nested := JsonNested{
		NestedValue1: "first nested value",
		NestedValue2: "second nested value",
	}
	
	jsonResp := JsonResponse{
		Value1:"some Data",
		Value2: "other Data",
		JsonNested: nested,
	}
	//My code
	//JsonName
	from := JsonInfo{
		Chrom: "Gender",
		Status: "Student",
	}

	contact := JsonCont{

		Num: "678-668-1220",
		Add: "3534 University park",
	}
	
	JsonInfo := JsonName{
		
		FName: "Adeola",
		LName: "Ogundipe",
		JsonCont: contact,
		JsonInfo: from,
	}
	http.Handle("/static/", 
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))
			
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
http.HandleFunc("/jasonResponse", func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(jsonResp)
})

http.HandleFunc("/userInfo", func(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(JsonInfo)
})
	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
