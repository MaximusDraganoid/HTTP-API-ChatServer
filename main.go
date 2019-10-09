// AvitoTry project main.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//да это плохо, но это временно
var numOfChat int  // будет использоваться для хранения id чата
var numOfUsers int // будет использоваться для хранения id пользователя

type User struct {
	id       int
	username string
	//createdAt time.Time  //нужно будет менять, т.к. нужет формат данных сооздания пользователя
}

type Chat struct {
	id           int    //id чата
	name         string //имя чата
	firstUserId  int
	secondUserId int
}

type Message struct {
	userId int
	chatId int
	text   string
	//createdAt time.Time // нужно будет менять, т.к. нужет формат данных сооздания пользователя
}

func startInit() {
	numOfChat = 0
	numOfUsers = 0
}

func addUser(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Add user")
	fmt.Println("Add User")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error of addUser")
	}

	var stringUsers map[string]string //создаем срез для чтения из body, который формата json

	jsonErr := json.Unmarshal(body, &stringUsers)
	if jsonErr != nil {
		fmt.Println("JsonError : ", jsonErr)
	}

	// формирование нового пользователя
	key := stringUsers["username"]
	tempId := numOfUsers
	numOfUsers = numOfUsers + 1
	newUser := User{id: tempId, username: key}

	fmt.Println(newUser)
	fmt.Println(req.Header)
	fmt.Println(req.Host)
	fmt.Println(req.Method)
	fmt.Println("*************************************************************************************")

}

//нормально, работает
func addChat(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Add chat")
	fmt.Println("Add chat")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error of sendMessage")
	}

	var stringInfo map[string]string //создаем срез для чтения из body, который формата json

	jsonErr := json.Unmarshal(body, &stringInfo)
	if jsonErr != nil {
		fmt.Println("JsonError : ", jsonErr)
	}

	//формируем новый чат
	tempChatId := numOfChat
	numOfChat = numOfChat + 1
	tempFirstId, err := strconv.Atoi(stringInfo["firstUser"])
	tempSecondId, err := strconv.Atoi(stringInfo["secondUser"])
	tempChatName := stringInfo["name"]
	var newChat Chat = Chat{id: tempChatId, firstUserId: tempFirstId, secondUserId: tempSecondId, name: tempChatName}

	fmt.Println(newChat)
}

//нормально
func sendMessage(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Send message")
	fmt.Println("Send message")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("error of sendMessage")
	}

	var stringUsers map[string]string //создаем срез для чтения из body, который формата json

	jsonErr := json.Unmarshal(body, &stringUsers)
	if jsonErr != nil {
		fmt.Println("JsonError : ", jsonErr)
	}

	//формирование нового сообщения
	author, err := strconv.Atoi(stringUsers["author"])
	tempChatId, err := strconv.Atoi(stringUsers["chat"])
	tempBody := stringUsers["text"]
	var newMessage Message = Message{userId: author, chatId: tempChatId, text: tempBody}

	fmt.Println(newMessage)
	fmt.Println(req.Header)
	fmt.Println(req.Host)
	fmt.Println(req.Method)
	fmt.Println("*************************************************************************************")

}

func getChatUsers(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Chat of users : ")
	fmt.Println("Get Chat Users")
}

func getAllMessage(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Al message of ... :")
	fmt.Println("Get all message")
}

func welcomeWords(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Welcome to YoungChat")
	fmt.Println("New user join...")
}

func main() {

	startInit()
	http.HandleFunc("/", welcomeWords)
	http.HandleFunc("/users/add", addUser)
	http.HandleFunc("/chats/add", addChat)
	http.HandleFunc("/message/add", sendMessage)
	http.HandleFunc("/chats/get", getChatUsers)
	http.HandleFunc("/message/get", getAllMessage)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		return
	}
	fmt.Println("server is working...")
}
