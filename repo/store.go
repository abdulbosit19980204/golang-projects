package repo

import (
	"fmt"
	"transaction/models"

	"github.com/jmoiron/sqlx"
)

type Repo struct {
	db *sqlx.DB
}

func New(d *sqlx.DB) Repo {
	return Repo{
		db: d,
	}
}

var (
	selectAllOne      = "SELECT * FROM student"
	selectByStudentId = "SELECT * FROM student WHERE student_id=$1"
	// insert            = "INSERT INTO student (country,city,telcode) VALUES($1,$2,$3)"
	sellectWithJoin = " SELECT student.student_id,student.first_name,student.last_name, student.age,study.start_date,study.end_date,address.country,address.city,address.home_address,address.zip_code,study_type.status_title,status.status_title FROM student	INNER JOIN address ON student.address_id = address.address_id INNER JOIN study ON student.student_id = study.student_id INNER JOIN study_type ON study.study_type_id = study_type.stusy_type_id INNER JOIn status ON study.status_id =  status.status_id ORDER BY student.student_id;"
)

func (r Repo) GetByStudentId(id int) ([]models.Student, []models.StudentInnerJoin, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, nil, err
	}

	defer func() {
		if err != nil {
			fmt.Println("Xatolik tufayli ish yakunlanmadi. Bajarilgan amallar ortga qaytarildi")
			_ = tx.Rollback()
		} else {
			fmt.Println("Succseed")
		}
	}()

	//-------------------------------------------
	allInformation := []dbStudent{}
	err = tx.Select(&allInformation, selectAllOne)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("get all information successed")

	gettingInformation := make([]models.Student, 0)

	for i := range allInformation {
		gettingInformation = append(gettingInformation, allInformation[i].StudentToModel())
	}
	//------------------------------------------

	//------------------------------------------
	byStudentId := dbStudent{}
	err = tx.Get(&byStudentId, selectByStudentId, 1)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("get infotmation by student id is complated")
	gettingInformation = append(gettingInformation, byStudentId.StudentToModel())
	//------------------------------------------

	//------------------------------------------
	withStudentInnerJoin := []tbStudentInnerJoin{}

	err = tx.Select(&withStudentInnerJoin, sellectWithJoin)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("All information Joined with each others")
	gettingInformationWithStudentInnerJoin := make([]models.StudentInnerJoin, 0)
	for i := range withStudentInnerJoin {
		gettingInformationWithStudentInnerJoin = append(gettingInformationWithStudentInnerJoin, withStudentInnerJoin[i].InnerJoinToModel())
	}

	//------------------------------------------

	//------------------------------------------

	//------------------------------------------
	if err := tx.Commit(); err != nil {
		return nil, nil, err
	}
	//------------------------------------------
	return gettingInformation, gettingInformationWithStudentInnerJoin, err
}

/*
func (r Repo) GetStudentAddress() ([]models.Address, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			return nil, err
		}
	}()
}
*/

/*
func (r Repo) GetByID(id int) ([]models.Place1, error) {
	collect := make([]models.Place1, 0)

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil {
			fmt.Println("xamma ishlar orqaga qaytti")
			_ = tx.Rollback()
		} else {
			fmt.Println("xamma ishlar bajarildi")
		}
	}()

	res, err := tx.Exec(insert, "usa", "georgia", 1)
	if err != nil {
		return nil, err
	}
	fmt.Println("birinchi query success")

	affacted, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	fmt.Println("birinchi query success")

	if affacted == 0 {
		return nil, errors.New("yaratilmadi")
	}
	fmt.Println("birinchi query tugadi")

	//-----------------------------------------
	all := []dbPlace1{}
	err = tx.Select(&all, selectAllOne)
	if err != nil {
		return nil, err
	}
	fmt.Println("ikkinchi query success")

	for i := range all {
		collect = append(collect, all[i].toModel())
	}
	//-------------------------------------------

	//-------------------------------------------
	byCountry := dbPlace1{}
	err = tx.Get(&byCountry, selectByCountryTwo, "Uzbekistan")
	if err != nil {
		return nil, err
	}
	fmt.Println("uchunchi query success")

	collect = append(collect, byCountry.toModel())

	if err := tx.Commit(); err != nil {
		return nil, err
	}
	//--------------------------------------------

	return collect, nil
}
*/
