#### install
`go get github.com/Earth-Online/zoomeye-api`

#### doc
https://godoc.org/github.com/Earth-Online/zoomeye-api

#### example
```
package main

 import (
 "fmt"
 "os"
 "github.com/Earth-Online/zoomeye-api"
 }
 
 func main(){
      const username = "you-email"
      const password = "you-password"
      user, err := zoomeye.NewUser(
                username,
                password,
      )
      if err != nil {
                fmt.Printf("login error")
                os.Exit(1)
      }
      info, err := user.HostSearch("title:oneindex", 1, []string{})
      if err != nil {
           fmt.Printf("search error")
           os.Exit(1)
      }
      fmt.Printf("have %d result, Available %d", info.Total, info.Available)
      for _, host := range info.Matches {
          fmt.Print(host.Ip)
      } 
}
 
```
