package repositories

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/otaviopontes/api-go/src/models"
	"github.com/redis/go-redis/v9"
)

type PostRepository interface {
	Create(userId uuid.UUID, post models.Post) error
	GetPostById(id uuid.UUID) (models.Post, error)
	GetPosts() ([]models.Post, error)
	Update(id uuid.UUID, post models.Post) error
	Delete(id uuid.UUID) error
	Like(id uuid.UUID) error
	Dislike(id uuid.UUID) error
}

type Posts struct {
	db    *sql.DB
	redis *redis.Client
}

func NewPostRepository(db *sql.DB, redis *redis.Client) *Posts {
	return &Posts{db, redis}
}

func (repository Posts) Create(userId uuid.UUID, post models.Post) error {
	statement, err := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3);")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(post.Title, post.Content, post.AuthorId)
	if err != nil {
		return err
	}

	repository.redis.Del(context.Background(), "posts")

	return nil
}

func (repository Posts) GetPostById(id uuid.UUID) (models.Post, error) {
	lines, err := repository.db.Query(`
	select p.*, u.nick from
	posts p inner join users u
	on u.id = p.author_id where p.id = $1
	`, id)
	if err != nil {
		return models.Post{}, err
	}
	var post models.Post
	if lines.Next() {

		err := lines.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		)
		if err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Posts) GetPosts() ([]models.Post, error) {
	cachedPosts, err := repository.redis.Get(context.Background(), "posts").Result()
	if err == nil {
		var posts []models.Post
		err := json.Unmarshal([]byte(cachedPosts), &posts)
		if err == nil {
			return posts, nil
		}
	}

	lines, err := repository.db.Query(`
	select p.*, u.nick 
	from posts p 
	join users u on u.id = p.author_id
	order by p.createdat desc;`,
	)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for lines.Next() {
		var post models.Post
		err := lines.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if len(posts) > 0 {
		postsJson, _ := json.Marshal(posts)
		repository.redis.Set(context.Background(), "posts", postsJson, 10*time.Minute)
	}

	return posts, nil
}

func (repository Posts) Update(id uuid.UUID, post models.Post) error {
	statement, err := repository.db.Prepare("update posts set title = $1, content = $2 where id = $3")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(post.Title, post.Content, id)
	if err != nil {
		return err
	}

	repository.redis.Del(context.Background(), "posts")

	return nil
}

func (repository Posts) Delete(id uuid.UUID) error {
	statement, err := repository.db.Prepare("delete from posts where id = $1")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	repository.redis.Del(context.Background(), "posts")

	return nil
}

func (repository Posts) Like(id uuid.UUID) error {
	statement, err := repository.db.Prepare("update posts set likes = likes + 1 where id = $1")
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	repository.redis.Del(context.Background(), "posts")
	return nil
}

func (repository Posts) Dislike(id uuid.UUID) error {
	statement, err := repository.db.Prepare(`
	update posts set likes =
	CASE 
		WHEN likes > 0 THEN likes - 1
		ELSE likes 
	END 
	where id = $1
	`)
	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	repository.redis.Del(context.Background(), "posts")

	return nil
}
