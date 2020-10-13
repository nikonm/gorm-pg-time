####GORM Time

[![Build Status](https://travis-ci.org/nikonm/gorm-pg-time.svg?branch=master)](https://travis-ci.org/nikonm/gorm-pg-time)

#####Time

````
import "github.com/nikonm/gorm-pg-time"

type SomeModel struct {
	gorm.Model
	Time gormpgtime.NullTime
}


tm, _ := time.Parse("15:04:05", "13:07:23")
md := &Model{
	Time: (&NullTime{}).Set(&tm),
}
DB.Create(&md)
// INSERT INTO `models` (`time`) VALUES ("13:07:23")

````
#####TimeTZ
Timezone read from time.Time object and set to timetz field in postgres 
````
import "github.com/nikonm/gorm-pg-time"

type SomeModel struct {
	gorm.Model
	Time gormpgtime.NullTimeTZ
}

tm, _ := time.ParseInLocation("15:04:05", "13:07:23", time.FixedZone("", +3600))
md := &Model{
	Time: (&NullTime{}).Set(&tm),
}
DB.Create(&md)
// INSERT INTO `models` (`time`) VALUES ("13:07:23+01:00")

````
