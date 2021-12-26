package main

import (
	"radioGame/internal/usecase"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	radio := make(chan string)

	talkMapAlpha := make(map[string][]string)
	talkMapAlpha["Браво"] = []string{"Как дела в Алматы?"}
	//talkMapAlpha["Чарли"] = []string{"Конец связи"}

	talkMapBravo := make(map[string][]string)
	talkMapBravo["Чарли"] = []string{"В Алматы спокойно", "Как дела в Астане?"}
	talkMapBravo["Браво"] = []string{"Если в Астане холодно, то иди домой"}

	talkMapCharli := make(map[string][]string)
	talkMapCharli["Альфа"] = []string{"В Астане холодно"}
	//talkMapCharli["Браво"] = []string{"Конец связи"}

	alpha := usecase.NewRadist(
		"Альфа",
		talkMapAlpha,
		radio,
		wg,
	)
	bravo := usecase.NewRadist(
		"Браво",
		talkMapBravo,
		radio,
		wg,
	)
	charli := usecase.NewRadist(
		"Чарли",
		talkMapCharli,
		radio,
		wg,
	)

	wg.Add(1)
	go bravo.Listen()
	wg.Add(1)
	go charli.Listen()
	wg.Add(1)
	go alpha.Listen()

	wg.Add(1)
	go alpha.Send()

	go func() {
		wg.Wait()
		close(radio)
	}()

	time.Sleep(1000 * time.Millisecond)

}
