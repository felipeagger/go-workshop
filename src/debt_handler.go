package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetDebts find all debts
func GetDebts(c *gin.Context) {

	var debts []debt
	allDebts := selectAll(&debts, c)

	c.JSON(200, allDebts)
}

//GetDebt find debt by ID
func GetDebt(c *gin.Context) {

	ID := c.Param("id")

	debt, _ := selectDebtByID(ID, c)

	c.JSON(200, debt)
}

//GetDebtsByUser find debt by userID
func GetDebtsByUser(c *gin.Context) {

	ID, _ := strconv.Atoi(c.Param("id"))

	debts := selectDebtsByUser(ID, c)

	c.JSON(200, debts)
}

//PostDebt Create debt by Body
func PostDebt(c *gin.Context) {

	var newDebt debt
	if err := c.ShouldBindJSON(&newDebt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbConnect()
	db.Create(&newDebt)

	c.JSON(201, newDebt)
}

//PutDebt Update debt by Body
func PutDebt(c *gin.Context) {

	ID := c.Param("id")

	var updateDebt debt
	if err := c.ShouldBindJSON(&updateDebt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	debt, db := selectDebtByID(ID, c)

	debt.UserID = updateDebt.UserID
	debt.CompanyName = updateDebt.CompanyName
	debt.Value = updateDebt.Value
	debt.Date = updateDebt.Date
	debt.Status = updateDebt.Status

	db.Save(&debt)

	c.JSON(200, debt)
}

//DeleteDebt delete debt by ID
func DeleteDebt(c *gin.Context) {

	ID := c.Param("id")

	debt, db := selectDebtByID(ID, c)

	if debt.ID != "" {
		db.Delete(&debt)
	}

	c.JSON(204, nil)
}
