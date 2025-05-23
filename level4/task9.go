package main

import "time"

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

type Report struct {
	User
	ReportID int
	Date     string
}

func CreateReport(user User, reportDate string) Report {
	var report Report
	report.ID = user.ID
	report.Name = user.Name
	report.Email = user.Email
	report.Age = user.Age
	report.Date = reportDate
	report.ReportID = time.Now().Nanosecond()
	return report
}

func GenerateUserReports(users []User, reportDate string) []Report {
	var reports []Report
	for _, user := range users {
		reports = append(reports, CreateReport(user, reportDate))
	}
	return reports
}
