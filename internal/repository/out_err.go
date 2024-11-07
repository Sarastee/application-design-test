package repository

import "errors"

const (
	errMsgNoRoomsAvailable = "no rooms available"
)

var (
	ErrNoRoomsAvailable = errors.New(errMsgNoRoomsAvailable) // ErrNoRoomsAvailable сигнальная ошибка в случае отсутствии свободных комнат
)
