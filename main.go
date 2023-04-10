package main

import(
"os"
"golang-restaurant-management/database"

)


func main(){
port := os.Getenv("PORT")

if port == ""{
	port = "8000"
}

router := gin.New()
router.use(gin.Logger())

}