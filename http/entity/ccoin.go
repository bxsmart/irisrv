package entity

type CCoin struct {
	ID        int64  `gorm: "column:id;type:bigint;primary_key;AUTO_INCREMENT"` // auto-increment by-default by gorm
	Version   string `gorm: "column:version;type:varchar(20)"`
	Salt      string `gorm: "column:salt;type:varchar(50)"`
	Username  string `gorm: "column:username;type:varchar(50)"`
	Password  string `gorm: "column:password;type:varchar(200)"`
	Languages string `gorm: "column:lang;type:varchar(50)"`

	BasicEntity
}

func (e *CCoin) TableName() string {
	return "jys_ccoin"
}
