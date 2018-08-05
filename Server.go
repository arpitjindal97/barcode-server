package main

import (
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var dataModel *CustomTableModel

func StartProcess(m *CustomTableModel) {

	dataModel = m
	PrepareExcelFile()

	fmt.Println("Starting server ...")

	muxHttp := mux.NewRouter()

	muxHttp.HandleFunc("/status", StatusHandler).Methods("GET")
	muxHttp.HandleFunc("/result", ResultHandler).Methods("POST")

	srv := &http.Server{
		Handler:      muxHttp,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

type Status struct {
	Status string `json:"status"`
}
type Response struct {
	Response string `json:"result"`
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	js, _ := json.MarshalIndent(&Status{"Running"}, "", "	")

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
func ResultHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	var result Response
	json.Unmarshal(body, &result)

	dataModel.recieved(TableItem{result.Response})
	js, _ := json.MarshalIndent(result, "", "	")

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	WriteToExcel(result.Response)
}

var fileName string = "sample.xlsx"

func PrepareExcelFile() {
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println("Creating new file")
		xlsx = excelize.NewFile()
		xlsx.SaveAs(fileName)
		return
	}
	xlsx.NewSheet("Sheet2")
	time.Sleep(time.Second * 1)
	xlsx.SetCellValue("Sheet1", "A1", "some random text")
	xlsx.DeleteSheet("Sheet1")
	xlsx.SetSheetName("Sheet2", "Sheet1")
	xlsx.Save()
}

func WriteToExcel(str string) {
	xlsx, _ := excelize.OpenFile(fileName)
	xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(len(xlsx.GetRows("Sheet1"))+1), str)
	xlsx.Save()

}
