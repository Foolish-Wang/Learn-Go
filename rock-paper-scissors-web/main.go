package main

import (
	"encoding/json"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 游戏选项常量
const (
    Rock     = "rock"
    Paper    = "paper"
    Scissors = "scissors"
)

// 游戏结果常量
const (
    Win  = "win"
    Lose = "lose"
    Draw = "draw"
)

// 游戏请求结构
type GameRequest struct {
    PlayerChoice string `json:"playerChoice"`
}

// 游戏响应结构
type GameResponse struct {
    PlayerChoice   string `json:"playerChoice"`
    ComputerChoice string `json:"computerChoice"`
    Result         string `json:"result"`
    Message        string `json:"message"`
}

func init() {
    // 初始化随机数生成器
    rand.Seed(time.Now().UnixNano())
}

func homePage(w http.ResponseWriter, r *http.Request) {
    // 设置正确的内容类型头
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    renderTemplate(w, "index.html", nil)
}

func playGame(w http.ResponseWriter, r *http.Request) {
    // 只接受POST请求
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // 解析请求体中的JSON数据
    var req GameRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // 验证玩家选择
    playerChoice := req.PlayerChoice
    if !isValidChoice(playerChoice) {
        http.Error(w, "Invalid choice", http.StatusBadRequest)
        return
    }

    // 生成电脑选择
    computerChoice := generateComputerChoice()

    // 确定游戏结果
    result, message := determineWinner(playerChoice, computerChoice)

    // 创建响应
    response := GameResponse{
        PlayerChoice:   playerChoice,
        ComputerChoice: computerChoice,
        Result:         result,
        Message:        message,
    }

    // 设置响应头和返回JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func isValidChoice(choice string) bool {
    return choice == Rock || choice == Paper || choice == Scissors
}

func generateComputerChoice() string {
    choices := []string{Rock, Paper, Scissors}
    return choices[rand.Intn(len(choices))]
}

func determineWinner(playerChoice, computerChoice string) (string, string) {
    // 平局情况
    if playerChoice == computerChoice {
        return Draw, "It's a draw!"
    }

    // 玩家获胜情况
    if (playerChoice == Rock && computerChoice == Scissors) ||
        (playerChoice == Paper && computerChoice == Rock) ||
        (playerChoice == Scissors && computerChoice == Paper) {
        return Win, "You win!"
    }

    // 电脑获胜情况
    return Lose, "Computer wins!"
}

func main() {
    // 静态文件服务器
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // API路由
    http.HandleFunc("/", homePage)
    http.HandleFunc("/play", playGame)

    // 启动服务器
    log.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
