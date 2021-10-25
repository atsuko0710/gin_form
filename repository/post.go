package repository

import "gin_forum/models"

func CreatePost(p models.Post) error {
	return models.CreatePost(p)	
}