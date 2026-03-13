package storage

import "cloud-task-service/models"

var Tasks = make(map[int]models.Task)
var CurrentID = 1
