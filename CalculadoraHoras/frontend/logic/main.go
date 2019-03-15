package main

import (
	"github.com/Project1/CalculadoraHoras/backend/calculadora"
	"github.com/gopherjs/gopherjs/js"
	"github.com/jaracil/ei"
	"github.com/jaracil/psgo"
	_ "github.com/jaracil/psgo/psjs"
)

var document = js.Global.Get("document")

func main() {
	subscriber := psgo.NewSubscriber(msgSubscriber)
	subscriber.Subscribe("web.calculate.get")
}

func msgSubscriber(msg *psgo.Msg) {
	subscribers := map[string]func(m *psgo.Msg){
		"web.calculate.get": onCalculate,
	}
	subscribers[msg.To](msg)
}

func onCalculate(msg *psgo.Msg) {
	days, hours, mins := calculadora.GetHoursAndMinutes(ei.N(msg.Dat).StringZ())
	msg.Answer(map[string]int{
		"days": days, "hours": hours, "minutes": int(mins),
	}, nil)

}

func getHoursFromInput() string {
	inputHours := document.Call("getElementById", "inputHours")
	return inputHours.Get("value").String()
}

/*
package main

import (
	"github.com/Project1/CalculadoraHoras/backend/calculadora"
	"github.com/gopherjs/gopherjs/js"
	"github.com/jaracil/ei"
	"github.com/jaracil/psgo"
	_ "github.com/jaracil/psgo/psjs"
)


var document = js.Global.Get("document")
var timeToGo *TimeToGo

func main() {
	timeToGo = &TimeToGo{Object: js.Global.Get("Object").New()}
	js.Global.Set("TimeToGo", timeToGo)
	document.Set("PublishHours", PublishHours)

	subscriber := psgo.NewSubscriber(msgSubscriber)
	subscriber.Subscribe("web.calculate")
}

func msgSubscriber(msg *psgo.Msg) {
	subscribers := map[string]func(m *psgo.Msg){
		"web.calculate": onCalculate,
	}
	subscribers[msg.To](msg)
}

func onCalculate(msg *psgo.Msg) {
	println(msg)
	days, hours, mins := calculadora.GetHoursAndMinutes(ei.N(msg.Dat).StringZ())
	println(days, hours, mins) //Publish results
}

func PublishHours() {
	timeToGo.Days, timeToGo.Hours, timeToGo.Minutes = calculadora.GetHoursAndMinutes(getHoursFromInput())
	println(timeToGo.Days)
}

func getHoursFromInput() string {
	inputHours := document.Call("getElementById", "inputHours")
	return inputHours.Get("value").String()
}
*/
