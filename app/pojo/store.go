package pojo

type Store struct {
	fileStoreId int `gorm:"primaryKey,AUTO_INCREMENT"`
	userId      int `gorm:"INDEX"`
	currentSize int64
	maxSize     int64
	user        User `gorm:"foreignKey:user_id"`
}

func (Store) TableName() string {
	return "file_store"
}
