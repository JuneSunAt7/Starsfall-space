package tui

import (
	"github.com/rivo/tview"
)

func RunMain() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Файерволл", "Настройка подключений", 'a', nil).
		AddItem("Сканирование сети", "Поиск открытых портов", 'b', nil).
		AddItem("Разработка фильтров", "Блокировка соединений", 'c', nil).
		AddItem("Доступ к данным", "", 'd', nil).
		AddItem("Выход", "Выход из программы", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}