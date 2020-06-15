package src

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUsers find all users
func GetUsers(c *gin.Context) {

	var users []user
	allUsers := selectAll(&users, c)

	c.JSON(200, allUsers)
}

//GetUser find one user
func GetUser(c *gin.Context) {
	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		checkErr(err, 400, c)
		return
	}

	user, _ := selectUserID(ID, c)

	c.JSON(200, user)
}

//PostUser Create user by Body
func PostUser(c *gin.Context) {

	var newUser user

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := dbConnect()
	db.Create(&newUser)

	c.JSON(201, newUser)
}

//PutUser Update one user by Body
func PutUser(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		checkErr(err, 400, c)
		return
	}

	var updateUser user

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, db := selectUserID(ID, c)

	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.BirthDate = updateUser.BirthDate

	db.Save(&user)

	c.JSON(200, user)
}

//DeleteUser remove one user
func DeleteUser(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		checkErr(err, 400, c)
		return
	}

	user, db := selectUserID(ID, c)

	if user.ID > 0 {
		db.Delete(&user)
	}

	c.JSON(204, nil)
}
