package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Add the correct JSON struct tags with the required optional directives to the 'Appointment' struct.
type Appointment struct {
	PatientName string    `json:"patientName"`
	Description string    `json:"-"`
	StartTime   time.Time `json:"startTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
}

// DO NOT change the code within the main function! - Your task is only to add the correct JSON struct tags above!
func main() {
	// DO NOT delete! - This code block takes as an input the values for the 'appointment' struct:
	var patientName, description string
	var year, day, hour int
	fmt.Scanln(&patientName, &description, &year, &day, &hour)
	startTime := time.Date(year, time.April, day, hour, 0, 0, 0, time.UTC)
	endTime := startTime.Add(time.Hour * 1)

	appointment := Appointment{
		PatientName: patientName,
		Description: description,
		StartTime:   startTime,
		EndTime:     endTime,
	}

	appointmentJSON, err := json.Marshal(appointment)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(appointmentJSON))
}
