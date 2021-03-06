// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Actor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Briefly  string        `json:"briefly"`
	Category EventCategory `json:"category"`
	Host     *Host         `json:"host"`
	Shows    []*Show       `json:"shows"`
}

type Host struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type NewEvent struct {
	Name      string        `json:"name"`
	Cover     string        `json:"cover"`
	Thumbnail string        `json:"thumbnail"`
	Briefly   string        `json:"briefly"`
	Category  EventCategory `json:"category"`
	Show      []*NewShow    `json:"show"`
}

type NewShow struct {
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
	Price    []*int    `json:"price"`
}

type Show struct {
	ID       string    `json:"id"`
	Time     time.Time `json:"Time"`
	Location string    `json:"location"`
	Price    []int     `json:"price"`
	Link     string    `json:"link"`
	Actors   []*Actor  `json:"actors"`
}

type EventCategory string

const (
	EventCategoryStandup   EventCategory = "standup"
	EventCategoryManzai    EventCategory = "manzai"
	EventCategoryImpromptu EventCategory = "impromptu"
)

var AllEventCategory = []EventCategory{
	EventCategoryStandup,
	EventCategoryManzai,
	EventCategoryImpromptu,
}

func (e EventCategory) IsValid() bool {
	switch e {
	case EventCategoryStandup, EventCategoryManzai, EventCategoryImpromptu:
		return true
	}
	return false
}

func (e EventCategory) String() string {
	return string(e)
}

func (e *EventCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EventCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EventCategory", str)
	}
	return nil
}

func (e EventCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
