package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	TimestampI               = 0
	CurrentlyEnrolledI       = 2
	CurrentlyEnrolledCampusI = 3
	OnWaitlistI              = 4
	NoCampusChoiceI          = 24
	ChildLastNameI           = 25
	ChildFirstNameI          = 26
)

var schoolNames = []string{
	"Blackland Prairie",
	"Brushy Creek",
	"Cactus Ranch",
	"Callison",
	"Caraway",
	"Deep Wood",
	"Double File Trail",
	"Forest North",
	"Great Oaks",
	"Herrington",
	"Jollyville",
	"Live Oak",
	"Pond Springs",
	"Purple Sage",
	"Sommer",
	"Teravista",
	"Union Hill",
	"Voigt",
	"Wells Branch",
}

var schools = map[string][]*Child{
	"Blackland Prairie": make([]*Child, 0, 6),
	"Brushy Creek":      make([]*Child, 0, 6),
	"Cactus Ranch":      make([]*Child, 0, 6),
	"Callison":          make([]*Child, 0, 6),
	"Caraway":           make([]*Child, 0, 6),
	"Deep Wood":         make([]*Child, 0, 6),
	"Double File Trail": make([]*Child, 0, 6),
	"Forest North":      make([]*Child, 0, 6),
	"Great Oaks":        make([]*Child, 0, 6),
	"Herrington":        make([]*Child, 0, 6),
	"Jollyville":        make([]*Child, 0, 6),
	"Live Oak":          make([]*Child, 0, 6),
	"Pond Springs":      make([]*Child, 0, 6),
	"Purple Sage":       make([]*Child, 0, 6),
	"Sommer":            make([]*Child, 0, 6),
	"Teravista":         make([]*Child, 0, 6),
	"Union Hill":        make([]*Child, 0, 6),
	"Voigt":             make([]*Child, 0, 6),
	"Wells Branch":      make([]*Child, 0, 6),
}

type Child struct {
	CurrentlyEnrolled       bool
	CurrentlyEnrolledCampus string
	OnWaitlist              bool
	FirstCampusChoice       string
	SecondCampusChoice      string
	ThridCampusChoice       string
	FourthCampusChoice      string
	FifthCampusChoice       string
	NoCampusChoice          bool
	LastName                string
	FirstName               string
	Timestamp               string
}

func main() {
	children := SortableChildren(getChildren())

	sort.Sort(children)

	for i := 0; i < len(children); i++ {
		child := children[i]

		if child.CurrentlyEnrolled {
			placeChild(child)
		} else if child.OnWaitlist {
			placeChild(child)
		} else {
			placeChild(child)
		}
	}

	fmt.Println(schools)
}

func placeChild(c *Child) {
	if c.NoCampusChoice {
		placeChildRandom(c)
		return
	}

	if len(schools[c.FirstCampusChoice]) < 6 {
		schools[c.FirstCampusChoice] = append(schools[c.FirstCampusChoice], c)
	} else if len(schools[c.SecondCampusChoice]) < 6 {
		schools[c.SecondCampusChoice] = append(schools[c.SecondCampusChoice], c)
	} else if len(schools[c.ThridCampusChoice]) < 6 {
		schools[c.ThridCampusChoice] = append(schools[c.ThridCampusChoice], c)
	} else if len(schools[c.FourthCampusChoice]) < 6 {
		schools[c.FourthCampusChoice] = append(schools[c.FourthCampusChoice], c)
	} else if len(schools[c.FifthCampusChoice]) < 6 {
		schools[c.FifthCampusChoice] = append(schools[c.FifthCampusChoice], c)
	}
}

func placeChildRandom(c *Child) {
	rand.Seed(time.Now().UnixNano())

	rIndex := rand.Intn(len(schoolNames))

	for {
		if len(schools[schoolNames[rIndex]]) < 6 {
			schools[schoolNames[rIndex]] = append(schools[schoolNames[rIndex]], c)
			break
		}
		rIndex = rand.Intn(len(schoolNames))
	}
}

func getChildren() []*Child {
	data, err := getFileData()
	if err != nil {
		log.Fatal(err)
	}

	children := make([]*Child, 0)

	for i := 1; i < len(data); i++ {
		c := &Child{}

		currentlyEnrolledVal := cleanData(data[i][CurrentlyEnrolledI])

		if currentlyEnrolledVal == "yes" || currentlyEnrolledVal == "" {
			c.CurrentlyEnrolled = true
		}

		c.CurrentlyEnrolledCampus = cleanData(data[i][CurrentlyEnrolledCampusI])

		onWaitlistVal := data[i][OnWaitlistI]

		if onWaitlistVal == "yes" || onWaitlistVal == "" {
			c.OnWaitlist = true
		}

		for j := 5; j < NoCampusChoiceI; j++ {
			campusChoice := data[i][j]
			campusName := strings.Trim(data[0][j], "[ ]")

			switch campusChoice {
			case "First Campus Choice":
				c.FirstCampusChoice = campusName
			case "Second Campus Choice":
				c.SecondCampusChoice = campusName
			case "Third Campus Choice":
				c.ThridCampusChoice = campusName
			case "Fourth Campus Choice":
				c.FourthCampusChoice = campusName
			case "Fifth Campus Choice":
				c.FifthCampusChoice = campusName
			}
		}

		noCampusChoiceVal := cleanData(data[i][NoCampusChoiceI])

		if noCampusChoiceVal == "yes" || noCampusChoiceVal == "" {
			c.NoCampusChoice = true
		}

		c.Timestamp = data[i][TimestampI]

		children = append(children, c)
	}

	return children
}

func getFileData() ([][]string, error) {
	file, err := os.Open("twc-app2.csv")
	if err != nil {
		return nil, err
	}

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func cleanData(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

type SortableChildren []*Child

func (a SortableChildren) Len() int           { return len(a) }
func (a SortableChildren) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortableChildren) Less(i, j int) bool { return a[i].Timestamp < a[j].Timestamp }
