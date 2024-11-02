// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID    `json:"id"`
	Fullname    string       `json:"fullname"`
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Phonenumber string       `json:"phonenumber"`
	Role        int32        `json:"role"`
	Isactive    bool         `json:"isactive"`
	Createdat   time.Time    `json:"createdat"`
	Modifiedat  sql.NullTime `json:"modifiedat"`
}
