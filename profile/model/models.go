package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type VolunteerDetails struct{
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name,omitempty"`
	Role   string             `json:"role,omitempty"`
	Email  string              `json:"email,omitempty" bson:"email,omitempty"`
	Mob_No string			  `json:"mob_number,omitempty"`
	Blood_grouop string		  `json:"blood_group,omitempty"`
	Address string			  `json:"address,omitempty"`
} 

