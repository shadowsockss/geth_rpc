package main

import (
    "fmt"
    "github.com/ethereum/go-ethereum/rpc"
)

func creatNewAccount(ip string, port string, password string) string {

    client, err := rpc.Dial("http://" + ip + ":" + port)
    if err != nil {
        fmt.Println("rpc.Dial err", err)
        return ""
    }   

//    var account[]string
//    err = client.Call(&account, "eth_accounts")
    var account string

    err = client.Call(&account, "personal_newAccount", password)
    //err = ec.c.CallContext(ctx, &result, "eth_getBalance", account, "latest")

    if err != nil {
        fmt.Println("client.Call err", err)
        return ""
    }   

 //   fmt.Printf("account[0]: %s\nbalance[0]: %s\n", account, result)
    return account
}

func getBalance(ip string, port string, account string) string {

    client, err := rpc.Dial("http://" + ip + ":" + port)
    if err != nil {
        fmt.Println("rpc.Dial err", err)
        return ""
    }   

//    var account[]string
//    err = client.Call(&account, "eth_accounts")
    var result string
    
    err = client.Call(&result, "eth_getBalance", account, "latest")
    //err = ec.c.CallContext(ctx, &result, "eth_getBalance", account, "latest")

    if err != nil {
        fmt.Println("client.Call err", err)
        return ""
    }   

//    fmt.Printf("account[0]: %s\nbalance[0]: %s\n", account, result)
    return result
}
func main(){
	account := creatNewAccount("localhost","8545","asdfgh123456")
	balance := getBalance("localhost","8545", account)
	fmt.Printf("new_account: %s\nbalance: %s\n", account, balance)
}
