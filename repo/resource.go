package repo

import (
	"time"
	"transaction/models"
)

type dbStudent struct {
	Student_id  int    `db:"student_id"`
	First_name  string `db:"first_name"`
	Last_name   string `db:"last_name"`
	Father_name string `db:"father_name"`
	Age         int    `db:"age"`
	Address_id  int    `db:"address_id"`
}

func (st dbStudent) StudentToModel() models.Student {
	return models.Student{
		Student_id:  st.Student_id,
		First_name:  st.First_name,
		Last_name:   st.Last_name,
		Father_name: st.Father_name,
		Age:         st.Age,
		Address_id:  st.Address_id,
	}
}

type tbAddress struct {
	Address_id   int    `db:"address_id"`
	Country      string `db:"country"`
	City         string `db:"city"`
	Zip_code     string `db:"zip_code"`
	Home_address string `db:"home_address"`
}

func (ad tbAddress) AddressToModel() models.Address {
	return models.Address{
		Address_id:   ad.Address_id,
		Country:      ad.Country,
		City:         ad.City,
		Zip_code:     ad.Zip_code,
		Home_address: ad.Home_address,
	}
}

type tbStudentInnerJoin struct {
	Student_id       int       `db:"student_id"`
	First_name       string    `db:"first_name"`
	Last_name        string    `db:"last_name"`
	Father_name      string    `db:"father_name"`
	Age              int       `db:"age"`
	Address_id       int       `db:"address_id"`
	Country          string    `db:"country"`
	City             string    `db:"city"`
	Zip_code         string    `db:"zip_code"`
	Home_address     string    `db:"home_address"`
	Status_id        int       `db:"status_id"`
	Status_title     string    `db:"status_title"`
	Study_id         int       `db:"study_id"`
	Start_date       time.Time `db:"start_date"`
	End_date         time.Time `db:"end_date"`
	Stusy_type_title string    `db:"stusy_type.status_title"`
	Study_type_id    int       `db:"study_type_id"`
	//Stusy_type_id int       `db:"stusy_type_id"`
}

func (inJst tbStudentInnerJoin) InnerJoinToModel() models.StudentInnerJoin {
	return models.StudentInnerJoin{
		Student_id:       inJst.Student_id,
		First_name:       inJst.First_name,
		Last_name:        inJst.Last_name,
		Father_name:      inJst.Father_name,
		Age:              inJst.Age,
		Country:          inJst.Country,
		City:             inJst.City,
		Zip_code:         inJst.Zip_code,
		Home_address:     inJst.Home_address,
		Status_title:     inJst.Status_title,
		Start_date:       inJst.Start_date,
		End_date:         inJst.End_date,
		Stusy_type_title: inJst.Stusy_type_title,
		Address_id:       inJst.Address_id,
		Status_id:        inJst.Status_id,
		Study_id:         inJst.Study_id,
		Stusy_type_id:    inJst.Study_type_id,
	}
}
