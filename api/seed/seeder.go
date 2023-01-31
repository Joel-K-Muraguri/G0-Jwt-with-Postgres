package seed

import (
	"log"

	"github.com/Joel-K-Muraguri/go-jwt/api/models"
	"github.com/jinzhu/gorm"

)

var users = []models.User{
	models.User{
		Nickname: "Joel Muraguri",
		Email: "joel@gmail.com",
		Password: "0113358919",
	},
	models.User{
		Nickname: "Mark Mayaka",
		Email: "mark@gmail.com",
		Password: "0123456789",
		
	},
}

func Load(db *gorm.DB){

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table : %v", err)		
	}

	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table : %v", err)
		
	
	}


	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed games table : %v", err)		
		}
	}
}