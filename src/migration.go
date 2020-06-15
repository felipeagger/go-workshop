package src

//AutoMigration update the database schema
func AutoMigration() {
	db := dbConnect()
	defer db.Close()

	db.AutoMigrate(user{})
}
