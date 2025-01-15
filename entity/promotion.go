package entity
import(
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Email	string `valid:"required~email is required, email~Email is invalid"`
	Phone     string `valid:"required~Phone is required, stringlength(10|10)"`
}