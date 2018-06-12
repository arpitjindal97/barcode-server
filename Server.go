package main

import (
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

	muxHttp.HandleFunc("/Status", StatusHandler).Methods("GET")
	muxHttp.HandleFunc("/Result", ResultHandler).Methods("POST")

	srv := &http.Server{
		Handler:      muxHttp,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Running")
}
func ResultHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	strBody := string(body)

	dataModel.recieved(TableItem{strBody})
	fmt.Fprint(w, strBody)
	WriteToExcel(strBody)
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
	xlsx.DeleteSheet("Sheet1")
	xlsx.SetSheetName("Sheet2", "Sheet1")
	xlsx.Save()
}

func WriteToExcel(str string) {
	xlsx, _ := excelize.OpenFile(fileName)
	xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(len(xlsx.GetRows("Sheet1"))+1), str)
	xlsx.Save()

}
