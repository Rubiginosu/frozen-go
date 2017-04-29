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
package server

import(
    "initial"
    "encoding/json"
    "fmt"
    "net/http"
)
// 验证panel
func auth(hashCode string) bool {
    // TODO 实现验证方法
    return true
}

type Answer struct {
    Success bool
    Error string
    Data string
}

func getVersion(w http.ResponseWriter,r *http.Request){
    version := initial.VERSION
    r.ParseForm()
    var answer Answer
    if auth(getAuthHashCode(r)) {
        // 验证成功
        answer.Success = true
        answer.Error = ""
        answer.Data = version
    } else {
        authInvalid(&answer)
    }
    bs ,_ := json.Marshal(answer)
    fmt.Fprint(w,string(bs))
}

func getAuthHashCode(r *http.Request) string{
    return r.Form["auth"][0]
}

func authInvalid(answer *Answer){
    answer.Success = false
    answer.Error = "Auth faild"
    answer.Data = ""
}
