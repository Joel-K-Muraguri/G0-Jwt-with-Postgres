package models

import "time"

type User struct{
	ID uint32  
	Nickname string
	Email string
	Password string
	Created_At time.Time
	Updated_At time.Time
}


func (u *User) SaveUser(){
	

}


func (u *User) FindAllUsers(){

	
}


func (u *User) FindAUser(){

	
}


func (u *User) UpdateUser(){

	
}


func (u *User) DeleteUser(){

	
}