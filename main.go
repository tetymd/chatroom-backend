package main

import (
    "net/http"
    "encoding/json"
)

type Comment struct {
    Name   string `json:"name"`
    Commnt string `json:"comment"`
}

func getCommentList(w http.ResponseWriter, r *http.Request) {
    comment := Comment{"tetsuya", "hello!"}
    res, _ := json.Marshal(comment)

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}

func main() {
    http.HandleFunc("/api/v1/comments/list", getCommentList)
    http.ListenAndServe(":8080", nil)
}
