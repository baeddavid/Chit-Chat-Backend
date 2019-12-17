package data

import (
	"time"
)

type Thread struct {
	Id int
	Uuid string
	Topic string
	userId int
	CreatedAt time.Time
}

type Post struct {
	Id int
	Uuid string
	Body string
	UserId int
	ThreadId int
	CreatedAt time.Time
}

func(thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func(post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func(thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

func(thread *Thread) Posts() (posts []Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, body, user_id, thread_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(createUUID(), body, user.Id, conv.Id, time.Now()).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
