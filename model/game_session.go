package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const (
	SessionWaiting  = "waiting"
	SessionReady    = "ready"
	SessionPlaying  = "playing"
	SessionFinished = "finished"
)

type GameSession struct {
	ID         primitive.ObjectID `bson:"_id"`
	MapID      string             `json:"map_id"`
	Status     string             `json:"status"`
	StartedAt  time.Time          `json:"started_at"`
	FinishedAt *time.Time         `json:"finished_at"`
}
