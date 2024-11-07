package service

import "errors"

const (
	errMsgIncorrectDate = "invalid date range"
)

var (
	ErrIncorrectDate = errors.New(errMsgIncorrectDate) // ErrIncorrectDate сигнальная ошибка в случае неправильного  диапазона дат
)
