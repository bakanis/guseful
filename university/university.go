package university

import (
	"github.com/coopernurse/gorp"
	"sort"
)

type (
	University struct {
		db *gorp.DbMap
	}

	Schedule []ScheduleItemView

	ScheduleItem struct {
		Id           int64
		PeriodType   int64
		Guru         int64
		Subject      int64
		Audithory    int64
		Corps        int64
		Start        int64
		Duration     int64
		WeekDay      int64
		Group        int64
		TrainingType int64
		Attendance   int64
		Billing      int64
	}

	ScheduleItemView struct {
		ScheduleItem
		PeriodTypeName   string
		GuruName         string
		SubjectName      string
		AudithoryName    string
		CorpsName        string
		StartName        string
		DurationName     string
		WeekDayName      string
		GroupName        string
		TrainingTypeName string
		AttendanceName   string
		BillingName      string
	}

	itemWithTitle struct {
		Id    int64
		Title string
	}

	PeriodType struct {
		itemWithTitle
	}

	Subject struct {
		itemWithTitle
	}

	Audithory struct {
		itemWithTitle
	}

	Corps struct {
		itemWithTitle
	}

	WeekDay struct {
		itemWithTitle
	}

	Group struct {
		itemWithTitle
		Parent int64
	}

	TrainingType struct {
		itemWithTitle
	}

	Attendance struct {
		itemWithTitle
	}

	Billing struct {
		itemWithTitle
	}
)

type ByTime []ScheduleItemView

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Start < a[j].Start }

type Int64 []int64

func (a Int64) Len() int           { return len(a) }
func (a Int64) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64) Less(i, j int) bool { return a[i] < a[j] }

func (s Schedule) MinTime() int64 {
	var min int64
	for i := range s {
		if s[i].Start < min {
			min = s[i].Start
		}
		if i == 0 {
			min = s[i].Start
		}
	}
	return min
}

func (s Schedule) MaxTime() int64 {
	var max int64
	for i := range s {
		if (s[i].Start + s[i].Duration) > max {
			max = s[i].Start + s[i].Duration
		}
	}
	return max
}

func (s Schedule) Startes() []int64 {
	var startes = map[int64]bool{}
	var res []int64
	for _, v := range s {
		startes[v.Start] = true
	}
	for k, _ := range startes {
		res = append(res, k)
	}
	sort.Sort(Int64(res))
	return res
}

func (s Schedule) PeriodTypes() []int64 {
	var periodTypes = map[int64]bool{}
	var res []int64
	for _, v := range s {
		periodTypes[v.PeriodType] = true
	}
	for k, _ := range periodTypes {
		res = append(res, k)
	}
	sort.Sort(Int64(res))
	return res
}