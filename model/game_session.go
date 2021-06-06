package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const GameSessionCollection = "game_sessions"

type GameSession struct {
	ID         primitive.ObjectID `bson:"_id"`
	MapID      string             `bson:"map_id"`
	Status     string             `bson:"status"`
	StartedAt  time.Time          `bson:"started_at"`
	FinishedAt *time.Time         `bson:"finished_at"`
}

const (
	StatusWaiting  = "waiting"
	StatusReady    = "ready"
	StatusPlaying  = "playing"
	StatusFinished = "finished"
)
