package utils

type Task struct {
    ID     int
    Data   string
    Vector []float32
}

type TaskVector struct {
    ID       int
    TaskName string
    Vector   []float32
}