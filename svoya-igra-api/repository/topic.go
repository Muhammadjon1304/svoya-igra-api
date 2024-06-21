package repository

import (
	"database/sql"
	"github.com/muhammadjon1304/svoya-igra-api/model"
	"log"
)

type TopicRepository struct {
	Db *sql.DB
}

func NewTopicRepository(db *sql.DB) TopicRepositoryInterface {
	return &TopicRepository{Db: db}
}

func (t *TopicRepository) DeleteTopic(id uint) bool {
	_, err := t.Db.Exec("DELETE FROM svoya_igra WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return false
	}
	return true

}
func (t *TopicRepository) GetAllTopic() []model.Topic {
	query, err := t.Db.Query("SELECT * from svoya_igra")
	if err != nil {
		log.Fatal(err)
	}
	var topics []model.Topic
	if query != nil {
		for query.Next() {
			var (
				id        uint
				title     string
				question1 string
				answer1   string
				question2 string
				answer2   string
				question3 string
				answer3   string
				question4 string
				answer4   string
				question5 string
				answer5   string
			)
			err := query.Scan(&id, &title, &question1, &answer1, &question2, &answer2, &question3, &answer3, &question4, &answer4, &question5, &answer5)
			if err != nil {
				log.Println(err)
			}
			topic := model.Topic{ID: id, Title: title, Question1: model.Question{question1, answer1, 10}, Question2: model.Question{question2, answer2, 20}, Question3: model.Question{question3, answer3, 30}, Question4: model.Question{question4, answer4, 40}, Question5: model.Question{question5, answer5, 50}}
			topics = append(topics, topic)
		}
	}
	return topics

}
func (t *TopicRepository) GetOneTopic(id uint) model.Topic {
	query, err := t.Db.Query("SELECT * FROM svoya_igra WHERE id=$1", id)
	if err != nil {
		log.Println(err)
		return model.Topic{}
	}
	var topic model.Topic
	if query != nil {
		for query.Next() {
			var (
				id        uint
				title     string
				question1 string
				answer1   string
				question2 string
				answer2   string
				question3 string
				answer3   string
				question4 string
				answer4   string
				question5 string
				answer5   string
			)
			err := query.Scan(&id, &title, &question1, &answer1, &question2, &answer2, &question3, &answer3, &question4, &answer4, &question5, &answer5)
			if err != nil {
				log.Println(err)
			}
			topic = model.Topic{ID: id, Title: title, Question1: model.Question{question1, answer1, 10}, Question2: model.Question{question2, answer2, 20}, Question3: model.Question{question3, answer3, 30}, Question4: model.Question{question4, answer4, 40}, Question5: model.Question{question5, answer5, 50}}
		}
	}
	return topic

}
func (t *TopicRepository) InsertTopic(post model.PostTopic) bool {
	stmt, err := t.Db.Prepare("INSERT INTO svoya_igra(title,question1,answer1,question2,answer2,question3,answer3,question4,answer4,question5,answer5) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err2 := stmt.Exec(post.Title, post.Question1.Question, post.Question1.Answer, post.Question2.Question, post.Question2.Answer, post.Question3.Question, post.Question3.Answer, post.Question4.Question, post.Question4.Answer, post.Question5.Question, post.Question5.Answer)
	if err2 != nil {
		log.Println(err2)
		return false
	}
	return true
}
func (t *TopicRepository) UpdateTopic(id uint, post model.PostTopic) model.Topic {
	_, err := t.Db.Exec("UPDATE svoya_igra SET title=$1,question1=$2,answer1=$3,question2=$4,answer2=$5,question3=$6,answer3=$7,question4=$8,answer4=$9,question5=$10,answer5=$11 WHERE id=$12", post.Title, post.Question1.Question, post.Question1.Answer, post.Question2.Question, post.Question2.Answer, post.Question3.Question, post.Question3.Answer, post.Question4.Question, post.Question4.Answer, post.Question5.Question, post.Question5.Answer, id)
	if err != nil {
		log.Println(err)
		return model.Topic{}
	}
	return t.GetOneTopic(id)
}
