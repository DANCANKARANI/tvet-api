package model

func Migration(){
	db.AutoMigrate(
		&Student{},
		&Job{},
		&Course{},
		&Department{},
		&Sponsor{},
		&Image{},
	)
}