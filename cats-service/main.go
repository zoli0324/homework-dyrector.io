package main

func main() {
	db := SetupDB()
	r := SetupRoutes(db)
	r.Run()
}
