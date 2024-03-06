package main

type AnswerType int

const (
    SingleAnswer AnswerType = iota
    MultipleChoices
    CheckList
)

type User struct {
    UserId      int     `json:"userId"`
    Username    string  `json:"username"`
    Password    string  `json:"password"`
}

type QuizDeck struct {
    DeckId      int     `json:"deckId"`
    Title       string  `json:"title"`
    Quizes      []Quiz  `json:"quizes"`
    CreatedUser int     `json:"createdUser"`
}

type Quiz struct {
    QuizId          int         `json:"quizId"`
    Question        string      `json:"question"`
    AnswerType      AnswerType  `json:"answerType"`
    SingleAnswer    string      `json:"singleAnswer"`
    MultipleChoices []string      `json:"multipleChoices"`
    CheckList       []string      `json:"checkList"`
    ParentDeck      int         `json:"parentDeck"`
}

