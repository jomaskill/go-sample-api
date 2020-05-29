package action

import (
	dbConnection "TesteWeb/framework"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
)

type Action struct {
	Uuid string `json:"uuid" gorm:"type:uuid;primary_key" valid:"_"`
	Name string `json:"name" gorm:"type:varchar(255)" valid:"required"`
	Code string `json:"code" gorm:"type:varchar(255);unique_index" valid:"required"`
	IsActive string `json:"is_active" gorm:"type:bool" valid:"required"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func InitialMigration()  {
	db := dbConnection.Start()

	db.AutoMigrate(&Action{})
}

func NewAction(data Action) (*Action, error){

	/*err := data.validate()
	if err != nil{
		return nil, err
	}*/

	id, _ := uuid.NewV4()
	action := &Action{
		Uuid: id.String(),
		Name: data.Name,
		Code: data.Code,
		IsActive: data.IsActive,
	}

	return action, nil
}

func (action *Action) validate() error{

	_, err := govalidator.ValidateStruct(action)
	if err != nil{
		return err
	}
	return nil
}
