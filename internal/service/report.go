package service

import (
	"ExpenceTracker/internal/repository"
	"time"
)

type Report struct {
	Period string `json:"period"`
	Amount int    `json:"amount"`
}

func GetDailyReport(userID int) (*Report, error) {
	now := time.Now()
	from := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	to := from.Add(24 * time.Hour)

	amount, err := repository.GetExpenseSumDate(userID, from, to)
	if err != nil {
		return nil, err
	}
	return &Report{Period: "Day", Amount: int(amount)}, nil
}

func GetWeeklyReport(userID int) (*Report, error) {
	now := time.Now()
	offset := int(int(now.Weekday()+6) % 7) // чтобы счет началось с понедельника

	from := now.AddDate(0, 0, -offset)
	to := now

	amount, err := repository.GetExpenseSumDate(userID, from, to)
	if err != nil {
		return nil, err
	}
	return &Report{Period: "Week", Amount: int(amount)}, nil
}

func GetMonthlyReport(userID int) (*Report, error) {
	now := time.Now()
	from := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	to := from.AddDate(0, 1, 0)

	amount, err := repository.GetExpenseSumDate(userID, from, to)
	if err != nil {
		return nil, err
	}
	return &Report{Period: "Month", Amount: int(amount)}, nil
}
