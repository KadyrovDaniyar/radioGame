package usecase

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type radist struct {
	name    string
	talkMap map[string][]string
	radio   chan string
	wg      *sync.WaitGroup
}

func NewRadist(
	name string,
	talkMap map[string][]string,
	radio chan string,
	wg *sync.WaitGroup) *radist {
	return &radist{
		name:    name,
		talkMap: talkMap,
		radio:   radio,
		wg:      wg,
	}
}

func (r *radist) Send() {
	defer r.wg.Done()

	for callName, messages := range r.talkMap {

		start := fmt.Sprintf("Говорит %s, вызываю %s", r.name, callName)

		r.radio <- start
		fmt.Println(start)

		for _, value := range messages {
			fmt.Println(value)
			if value == "Конец связи" {
				os.Exit(0)
			}
		}

		delete(r.talkMap, callName)

	}
}

func (r *radist) Listen() {
	defer r.wg.Done()

	for first := range r.radio {

		messageSplit := strings.FieldsFunc(first, split)
		if len(messageSplit) == 4 && messageSplit[3] == r.name {
			fmt.Println("Прием")
			r.wg.Add(1)
			r.Send()
		} else {
			r.radio <- first
		}
	}

}

func split(r rune) bool {
	return r == ' ' || r == ','
}
