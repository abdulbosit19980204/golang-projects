package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) user(c *gin.Context) {
	data, data1, err := h.store.GetByStudentId(1)
	fmt.Println("data", data)
	fmt.Println("data1", data1)
	fmt.Println("err", err)
	if err != nil {
		c.JSON(500, gin.H{
			"error_code": -1,
			"error_note": err.Error(),
			"data":       nil,
			"data1":      data1,
		})
		return
	}

	c.JSON(200, gin.H{
		"error_code": 0,
		"error_note": nil,
		"data":       data,
		"data1":      data1,
	})
}

//func(h *Handler)user-filter()

/*
	POST http://localhost:8080/user -- create
	models user(id,f,i,o,age,ad_id,sutd_id) info relation address(id,full...) sutdy(id,title,startq_date,end_date,status(complate,rejected,inporgress))

	GET  http://localhost:8080/users -- get all users

	GET  http://localhost:8080/users -- get all faqat paramterlar orqali: status lar boyicha(c,r,i),
	GET  http://localhost:8080/users?status=complated
	GET  http://localhost:8080/users?start=date1&end=date2 - ikki xil yillar orasidagi bitirganlarni oling

	POST http://localhost:8080/users -- body() search rule yozish kerak
	{
		"studne ismi":"abduvali",
		"tugantgan oliygohi dan kop bosin":2,
		"oliygolari xaqida malumot":[
				{
				"nomi",
				"tugantan sanasi"
			},
			{
				"nomi",
				"tugantan sanasi"
			}
		]
	}

	POST http://localhost:8080/users create
	POST http://localhost:8080/address create
	POST http://localhost:8080/study create

	xammasi uchun delete va update yozasiz
*/
