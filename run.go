package main

import (
	"encoding/json"
	. "fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func justRun(imeiCode string, distance string) bool {
	ios := "http://client4.aipao.me/api"
	iosb := false
	log.Println("run " + imeiCode + " " + distance)
	randomGenerateTable()
	req, _ := http.NewRequest("GET", apiRoot+"/%7Btoken%7D/QM_Users"+
		"/Login_AndroidSchool?IMEICode="+imeiCode, nil)
	req.Header.Add("version", appVersion)
	resInfo, _ := (&http.Client{}).Do(req)
	dataInfo, _ := ioutil.ReadAll(resInfo.Body)
	// log.Println(string(dataInfo))
	resInfo.Body.Close()
	returnData := &returnInfo{}
	_ = json.Unmarshal(dataInfo, returnData)
	if !returnData.Success {
		req, _ := http.NewRequest("GET", ios+"/%7Btoken%7D/QM_Users"+
			"/LoginSchool?IMEICode="+imeiCode, nil)
		req.Header.Add("Version", appVersionForIOS)
		resInfo, _ := (&http.Client{}).Do(req)
		dataInfo, _ := ioutil.ReadAll(resInfo.Body)
		log.Println(string(dataInfo))
		resInfo.Body.Close()
		_ = json.Unmarshal(dataInfo, returnData)
		iosb = true
	}
	url := apiRoot
	if iosb {
		url = ios
	}
	UserId := string(Sprintf("%d", returnData.Data.UserId))
	timespan := Sprintf("%d", time.Now().UnixNano()/1e6)
	nonce := Sprintf("%d", 100000+rand.Intn(9900000))
	sign := strings.ToUpper(MD5(returnData.Data.Token + nonce + timespan + UserId))
	time.Sleep(1 * time.Second)
	runTime, runDistance, runStep := randomGenerateInfo()
	if distance != "2400" && distance != "" {
		tmp, _ := strconv.Atoi(runDistance)
		tmp -= 400
		runDistance = strconv.Itoa(tmp)
		distance = "2000"
	}
	client := &http.Client{}

	requestRun, _ := http.NewRequest("GET", url+"/"+
		returnData.Data.Token+"/QM_Runs/SRS?S1="+
		longtitude+"&S2="+latitute+"&S3="+distance, nil)
	requestRun.Header.Add("nonce", nonce)
	requestRun.Header.Add("timespan", timespan)
	requestRun.Header.Add("sign", sign)
	requestRun.Header.Add("version", appVersion)
	requestRun.Header.Add("Accept", "text/html")
	requestRun.Header.Add("User-Agent", UserAgent)
	requestRun.Header.Add("Accept-Encoding", "gzip")
	requestRun.Header.Add("Connection", "Keep-Alive")
	resRun, _ := client.Do(requestRun)
	infoData, _ := ioutil.ReadAll(resRun.Body)
	// log.Println(string(infoData))
	resRun.Body.Close()
	returndata := &returnRun{}
	_ = json.Unmarshal(infoData, returndata)
	resEnd, _ := http.Get(url + "/" + returnData.Data.Token +
		"/QM_Runs/ES?S1=" + returndata.Data.RunId + "&S4=" +
		encrypt(runTime) + "&S5=" + encrypt(runDistance) +
		"&S6=" + returndata.Data.Routes + "&S7=1&S8=" +
		Sprintf("%s", table) + "&S9=" + encrypt(runStep))
	dataEnd, _ := ioutil.ReadAll(resEnd.Body)
	// log.Println(string(dataEnd))

	resEnd.Body.Close()
	returnEnd := &returnEnd{}
	_ = json.Unmarshal(dataEnd, returnEnd)
	 log.Println(returnData)
	 log.Println(returndata)
	 log.Println(returnEnd)
	if returnEnd.Success {
		return true
	} else {
		return false
	}
}
