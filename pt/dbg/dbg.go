package dbg

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/varunamachi/patti/pt"
	"github.com/varunamachi/teak"
)

func generateUsers(handle func(user *teak.User) error) error {

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	getRandomRole := func() teak.AuthLevel {
		i := r1.Intn(5)
		return teak.AuthLevel(i)
	}
	getRandomTitle := func() string {
		switch r1.Intn(5) {
		case 0:
			return "Ms"
		case 1:
			return "Mrs"
		case 2:
			return "Mr"
		case 3:
			return "Dr"
		case 4:
			return "Prof"
		}
		return "Chief Wizard"
	}

	for i := 0; i < 100; i++ {
		num := ToStrNum(i)
		id := "user_" + num.ID
		user := teak.User{
			UserID:     id,
			Email:      id + "@amail.com",
			Auth:       getRandomRole(),
			FirstName:  "User" + strconv.Itoa(i),
			LastName:   num.Name,
			Title:      getRandomTitle(),
			FullName:   "User" + strconv.Itoa(i) + " " + num.Name, //I'm bit lazy
			State:      teak.Active,
			VerID:      "",
			PwdExpiry:  time.Now().AddDate(1, 0, 0),
			CreatedAt:  time.Now(),
			CreatedBy:  "auto",
			ModifiedAt: time.Now(),
			ModifiedBy: "auto",
			VerfiedAt:  time.Now(),
			Props:      teak.M{},
		}
		if err := handle(&user); err != nil {
			return teak.LogError("pt.gen", err)

		}
	}
	return nil
}

func generateTaskList(handler func(tl *pt.TaskList) error) error {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < 200; i++ {
		userID := fmt.Sprintf("user_%d", r1.Intn(100))
		num := ToStrNum(i)
		tl := pt.TaskList{
			Item: pt.Item{
				Heading:     num.Name,
				Description: num.Name + num.Name,
				Status:      pt.Active,
			},
			UserID:     userID,
			CreatedOn:  time.Now(),
			CreatedBy:  userID,
			ModifiedOn: time.Now(),
			ModifiedBy: userID,
		}
		if err := handler(&tl); err != nil {
			teak.LogError("pt.gen", err)
			break
		}
	}
	return nil
}

// func createTasks() {

// }

// Below code is adapted from:
// https://github.com/siongui/userpages +
// /blob/master/content/articles/2017/12/31/ +
// go-convert-number-to-word-from-1-to-1000%25en.rst

type num struct {
	ID   string
	Name string
}

var numToWord = map[int]num{
	1:  {ID: "one", Name: "One"},
	2:  {ID: "two", Name: "Two"},
	3:  {ID: "three", Name: "Three"},
	4:  {ID: "four", Name: "Four"},
	5:  {ID: "five", Name: "Five"},
	6:  {ID: "six", Name: "Six"},
	7:  {ID: "seven", Name: "Seven"},
	8:  {ID: "eight", Name: "Eight"},
	9:  {ID: "nine", Name: "Nine"},
	10: {ID: "ten", Name: "Ten"},
	11: {ID: "eleven", Name: "Eleven"},
	12: {ID: "twelve", Name: "Twelve"},
	13: {ID: "thirteen", Name: "Thirteen"},
	14: {ID: "fourteen", Name: "Fourteen"},
	15: {ID: "fifteen", Name: "Fifteen"},
	16: {ID: "sixteen", Name: "Sixteen"},
	17: {ID: "seventeen", Name: "Seventeen"},
	18: {ID: "eighteen", Name: "Eighteen"},
	19: {ID: "nineteen", Name: "Nineteen"},
	20: {ID: "twenty", Name: "Twenty"},
	30: {ID: "thirty", Name: "Thirty"},
	40: {ID: "forty", Name: "Forty"},
	50: {ID: "fifty", Name: "Fifty"},
	60: {ID: "sixty", Name: "Sixty"},
	70: {ID: "seventy", Name: "Seventy"},
	80: {ID: "eighty", Name: "Eighty"},
	90: {ID: "ninety", Name: "Ninety"},
}

func convert1to99(n int) (w num) {
	if n < 20 {
		w = numToWord[n]
		return
	}

	r := n % 10
	if r == 0 {
		w = numToWord[n]
	} else {
		w.ID = numToWord[n-r].ID + "_" + numToWord[r].ID
		w.Name = numToWord[n-r].Name + " " + numToWord[r].Name
	}
	return
}

func convert100to999(n int) (w num) {
	q := n / 100
	r := n % 100
	w.ID = numToWord[q].ID + "_" + "hundred"
	w.Name = numToWord[q].Name + " " + "Hundred"
	if r == 0 {
		return
	} else {
		dec := convert1to99(r)
		w.ID = w.ID + "_" + dec.ID
		w.Name = w.Name + " and " + dec.Name
	}
	return
}

func ToStrNum(n int) (w num) {
	if n > 1000 || n < 1 {
		return num{
			Name: "Out of Range",
			ID:   "out_of_range",
		}
	}

	if n < 100 {
		w = convert1to99(n)
		return
	}
	if n == 1000 {
		w = num{
			Name: "One Thousand",
			ID:   "one_thousand",
		}
		return
	}
	w = convert100to999(n)
	return
}
