package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type ttabproceedings struct {
	actionkey string   `xml:"action-key-code" json:"action-key-code"`
	transdate string   `xml:"transaction-date" json:"transaction-date"`
	Proceed   *Proceed `xml:"proceeding-information" json:"proceeding-information.omitempty"`
}

type Proceed struct {
	ProceedEntry *ProceedEntry `xml:"proceeding-entry" json:"proceeding-entry"`
}

type ProceedEntry struct {
	number      string       `xml:"number" json:"number"`
	typecode    string       `xml:"type-code" json:"type-code"`
	filingdate  string       `xml:"filling-date" json:"filling-date"`
	employeenum string       `xml:"employee-number" json:"employee-number"`
	interloc    string       `xml:"interlocutory-attorney-name" json:"interlocutory-attorney-name"`
	loccode     string       `xml:"location-code" json:"location-code"`
	dayloc      string       `xml:"day-in-location" json:"day-in-location"`
	statup      string       `xml:"status-update-date" json:"status-update-date"`
	statuscod   string       `xml:"status-code" json:"status-code"`
	partyinfo   *partyinfo   `xml:"party-information" json :"party-information"`
	proshistory *proshistory `xml:"prosecution-history" json :"prosecution-history"`
}

type proshistory struct {
	proentry *proentry `xml:"prosecution-entry" json:"prosecution-entry"`
}

type proentry struct {
	identifier string `xml:"identifier" json:"identifier"`
	code       string `xml:"code" json:"code"`
	typecode   string `xml:"type-code" json:"type-code"`
	date       string `xml:"date" json:"date"`
	histext    string `xml:"history-text" json:"history-text"`
}

type partyinfo struct {
	party *party `xml:"party" json:"party"`
}

type party struct {
	identifier string    `xml:"identifier" json:"identifier"`
	rolecode   string    `xml:"role-code" json:"role-code"`
	name       string    `xml:"name" json:"name"`
	propinfo   *propinfo `xml:"property-information" json :"property-information"`
	addinfo    *addinfo  `xml:"address-information" json :"address-information"`
}

type propinfo struct {
	property *property `xml:"property" json :"property"`
}

type property struct {
	identifier   string `xml : "identifier" json : "identifier"`
	serialnumber string `xml : "serial-number" json : "serial-number"`
	marktext     string `xml : "mark-text " json : "mark-text" `
}

type addinfo struct {
	procaddr *procaddr `xml:"proceeding-address" json :"proceeding-address"`
}

type procaddr struct {
	identifier string `xml:"identifier" json:"identifier"`
	typecode   string `xml:"type-code" json:"type-code"`
	name       string `xml:"name" json:"name"`
	address    string `xml:"address-1" json:"address-1"`
	city       string `xml:"city" json:"city"`
	state      string `xml:"state" json:"state"`
	country    string `xml:"country" json:"country"`
	passcode   string `xml:"postcode" json:"postcode"`
}

func main() {
	xmlFile, err := os.Open("C:/Users/Asus/Desktop/ML _ Project/tt230101.xml")

	if err != nil {
		fmt.Println(err)
	}

	var data ttabproceedings
	xml.Unmarshal([]byte(xmlFile), &data)
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
	defer xmlFile.Close()
}
