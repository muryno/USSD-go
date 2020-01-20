package main
import (
	"fmt"
	"log"
	"net/http"
	"os"
)
func ussd_callback(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type","application/json")
	session_id := r.FormValue("sessionId")
	service_code := r.FormValue("serviceCode")
	phone_number := r.FormValue("phoneNumber")
	_ = fmt.Sprintf("%s,%s,%s",session_id,service_code,phone_number)
	text := r.FormValue("text")

	//you can write your logic here
	if len(text) == 0{
		w.Write([]byte("CON Welcome to Agric, Chose your language \n1. English \n2. Yoruba \n3. Igbo \n4. Hausa"))
		return
	}else{
		switch text{

		case "1":
			w.Write([]byte("CON Welcome to agric, Press \n1. Request for pickup\n2. Receive agricultural tips \n3.  Chat with us \n4. Go back to Language menu"))
			return
		case "1*1":
			w.Write([]byte("CON Please select agricultural you want to move \n1. Rice \n2. Tomato \n3. Maize \n4. Soyabeans \n5. Sweat potato"))
			return

		case "1*1*1":
			w.Write([]byte("CON Please enter tons of Rice you want to move"))
			return

		case "1*1*2":
			w.Write([]byte("CON Please enter tons of Tomato you want to move"))
			return

		case "1*1*3":
			w.Write([]byte("CON Please enter tons of Maize you want to move"))
			return

		case "1*1*4":
			w.Write([]byte("CON Please enter tons of Soya beans you want to move"))
			return

		case "1*1*5":
			w.Write([]byte("CON Please enter tons of Sweet potato  you want to move"))
			return

		case "1*1*5*1","1*1*4*2","1*1*3*3","1*1*2*4","1*1*1*4":
			w.Write([]byte("End Your request has been placed, Driver will be with you shortly..\n Thanks"))
			return
		//case "1*3":
		//	w.Write([]byte(fmt.Sprintf("END Your Phone Number is %s",phone_number)))
		//	return
		//
		//case "1*2":
		//	w.Write([]byte("END Your Balance is NGN 20,000"))
		//	return
		default:
			w.Write([]byte("END Invalid input"))
			return
		}
	}
}

func main(){
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	fmt.Println("This is a ussd application hosted using africastalking ussd platform, ")
	http.HandleFunc("/",ussd_callback)
	log.Fatal(http.ListenAndServe(":"+port,nil))
}