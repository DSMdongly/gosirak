package gogyo

import (
	"strings"
)

type SchoolKind int

const (
	KinderGarden SchoolKind = 1 + iota
	ElementrySchool
	MiddleSchool
	HighSchool
)

func GetSchoolKind(kind string) SchoolKind {
	kinds := map[string]SchoolKind{
		"kinder-garden":    KinderGarden,
		"elementry-school": ElementrySchool,
		"middle-school":    MiddleSchool,
		"high-school":      HighSchool,
	}

	return kinds[strings.ToLower(kind)]
}

type SchoolRegion string

const (
	Seoul     SchoolRegion = "stu.sen.go.kr"
	Incheon                = "stu.ice.go.kr"
	Busan                  = "stu.pen.go.kr"
	Gwangju                = "stu.gen.go.kr"
	Daejeon                = "stu.dje.go.kr"
	Daegu                  = "stu.dge.go.kr"
	Sejong                 = "stu.sje.go.kr"
	Ulsan                  = "stu.use.go.kr"
	Gyeonggi               = "stu.goe.go.kr"
	Kangwon                = "stu.kew.go.kr"
	Chungbuk               = "stu.cbe.go.kr"
	Chungnam               = "stu.cne.go.kr"
	Gyeongbuk              = "stu.gbe.go.kr"
	Gyeongnam              = "stu.gne.go.kr"
	Jeonbuk                = "stu.jbe.go.kr"
	Jeonnam                = "stu.jne.go.kr"
	Jeju                   = "stu.jje.go.kr"
)

func GetSchoolRegion(region string) SchoolRegion {
	regions := map[string]SchoolRegion{
		"seoul":     Seoul,
		"incheon":   Incheon,
		"busan":     Busan,
		"gwangju":   Gwangju,
		"daejeon":   Daejeon,
		"daegu":     Daegu,
		"sejong":    Sejong,
		"ulsan":     Ulsan,
		"gyeonggi":  Gyeonggi,
		"kangwon":   Kangwon,
		"chungbuk":  Chungbuk,
		"chungnam":  Chungnam,
		"gyeongbuk": Gyeongbuk,
		"gyeongnam": Gyeongnam,
		"jeonbuk":   Jeonbuk,
		"jeonnam":   Jeonnam,
		"jeju":      Jeju,
	}

	return regions[strings.ToLower(region)]
}

type SchoolCode string

type School struct {
	Kind   SchoolKind
	Region SchoolRegion
	Code   SchoolCode
}

func NewSchool(kind SchoolKind, region SchoolRegion, code SchoolCode) *School {
	school := new(School)
	school.Kind = kind
	school.Region = region
	school.Code = code

	return school
}

func (school School) GetDailyMenu(year int, month int, date int) Menu {
	return GetDailyMenu(school, year, month, date)
}

func (school School) GetMonthlyMenus(year int, month int) map[int]Menu {
	return GetMonthlyMenus(school, year, month)
}
