

SELECT student.student_id,student.first_name,student.last_name, student.age,study.start_date,study.end_date,address.country,address.city,address.home_address,address.zip_code,study_type.status_title,status.status_title FROM student 
INNER JOIN address ON student.address_id = address.address_id
INNER JOIN study ON student.student_id = study.student_id
INNER JOIN study_type ON study.study_type_id = study_type.stusy_type_id
INNER JOIn status ON study.status_id =  status.status_id

ORDER BY student.student_id;