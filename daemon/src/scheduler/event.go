/*
Author Axoford12
Team Freeze
Org Rubiginosu
  _____                        ____
|  ___| __ ___ _______ _ __  / ___| ___
| |_ | '__/ _ \_  / _ \ '_ \| |  _ / _ \
|  _|| | | (_) / /  __/ | | | |_| | (_) |
|_|  |_|  \___/___\___|_| |_|\____|\___/

 */

package scheduler

type Event struct {

}

type EventScheduler struct {
	Events map[string][]func(*Event)
}


// 触发这个事件
// name: 对事件的描述
// 例如 ： ServerCreateRequestReceived
// event: 事件的整体记录
// TODO: 现在还没具体实现这个事件，预计将会在以后的版本中实现
func Trigger(name string,event *Event){
	if r,ok := schedule.Event.Events[name];ok {
		for _,f := range r {
			f(event)
		}
	}
}

// 绑定事件处理器
func HandleFunc(name string,function func(*Event)){
	schedule.Event.Events[name] = append(schedule.Event.Events[name],function)
}