package models

import (
	"time"
)

type Student struct {
	Student_id  int
	First_name  string
	Last_name   string
	Father_name string
	Age         int
	Address_id  int
}

type Address struct {
	Address_id   int
	Country      string
	City         string
	Zip_code     string
	Home_address string
}

type Status struct {
	Status_id    int
	Status_title string
}

type Study struct {
	Study_id      int
	Start_date    time.Time
	End_date      time.Time
	Status_id     int
	Study_type_id int
	Student_id    int
}

type StudyType struct {
	Stusy_type_id int
	Status_title  string
}

type StudentInnerJoin struct {
	Student_id       int
	First_name       string
	Last_name        string
	Father_name      string
	Age              int
	Address_id       int
	Country          string
	City             string
	Zip_code         string
	Home_address     string
	Status_id        int
	Status_title     string
	Study_id         int
	Start_date       time.Time
	End_date         time.Time
	Stusy_type_title string
	Stusy_type_id    int
}
