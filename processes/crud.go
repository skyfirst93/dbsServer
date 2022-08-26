package processes

import (
	"dbssever/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var answersKey = "ANSWER_KEY_"
var historyKey = "HISTORY_KEY_"
var userMap = make(map[string]models.Authenticate)
var optMap = make(map[string]bool)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var authenticate models.Authenticate
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data correctly")
		return
	}

	json.Unmarshal(reqBody, &authenticate)
	fmt.Println("data ", authenticate)
	var associate models.Associate
	associate.RequestID = "88ydEE-ioiwe=="
	associate.AssociationID = "375dhjf9-Uydd="
	userMap[associate.AssociationID] = authenticate
	fmt.Println("map ", userMap)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(associate)
}

func AssociateUser(w http.ResponseWriter, r *http.Request) {
	var associate models.Associate
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data correctly")
	}
	json.Unmarshal(reqBody, &associate)
	fmt.Println("data ", associate)
	_, ok := userMap[associate.AssociationID]
	if ok {
		fmt.Println("map ", userMap)
		w.WriteHeader(http.StatusOK)
		var s models.Response
		s.Response = "User Association Sucessful"
		json.NewEncoder(w).Encode(s)
	} else {
		fmt.Println("map ", userMap, associate.AssociationID)
		w.WriteHeader(http.StatusNotFound)
		var s models.Response
		s.Response = "User Association Failed"
		json.NewEncoder(w).Encode(s)
	}
}

func GenerateOtp(w http.ResponseWriter, r *http.Request) {
	var otpDetails models.OtpDetails
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data correctly")
	}
	json.Unmarshal(reqBody, &otpDetails)
	fmt.Println("data ", otpDetails)
	_, ok := userMap[otpDetails.AssociationID]
	if ok {
		w.WriteHeader(http.StatusOK)
		generateOtp(otpDetails.EmailAddress)
		var s models.Response
		s.Response = "Otp Generated"
		json.NewEncoder(w).Encode(s)
	} else {
		w.WriteHeader(http.StatusNotFound)
		var s models.Response
		s.Response = "Failed to generate OTP"
		json.NewEncoder(w).Encode(s)
	}
}

func generateOtp(email string) {
	url := "https://erabhinav.pythonanywhere.com/send-otp"
	method := "POST"

	payload := strings.NewReader(`{"email": ` + email + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	var otp models.Otp
	json.Unmarshal(body, &otp)
	optMap[otp.Otp] = true
}

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	var otp models.Otp
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Kindly enter data correctly")
	}
	json.Unmarshal(reqBody, &otp)
	fmt.Println("data ", otp)
	_, ok := optMap[otp.Otp]
	if ok {
		w.WriteHeader(http.StatusOK)
		var s models.Response
		s.Response = "Otp Verified"
		json.NewEncoder(w).Encode(s)
	} else {
		w.WriteHeader(http.StatusNotFound)
		var s models.Response
		s.Response = "Otp Verification Failed"
		json.NewEncoder(w).Encode(s)
	}

}
