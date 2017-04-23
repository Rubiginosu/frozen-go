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
)
// 验证panel
func auth(hashCode string) bool {
    // TODO 实现验证方法
}

type Amswer struct {
    Success bool
    Error string
    Data string
}

func getVersion(w http.ResponseWriter,r *http.Request){
    version := initial.VERSION
    r.ParseForm()
    auth := r.Form['auth'][0]
    var answer answer
    if auth(auth) {
        // 验证成功
        answer.Success = true
        answer.Error = ""
        answer.Data = version
    } else {
    
        authInvalid(&answer)
    }
    
}

func authInvalid(answer *Amswer){
    answer.Success = false
    answer.Error = "Auth faild"
    answer.Data = ""
}
