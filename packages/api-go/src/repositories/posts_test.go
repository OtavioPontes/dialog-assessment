package repositories_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/google/uuid"
	"github.com/otaviopontes/api-go/src/models"
	"github.com/otaviopontes/api-go/src/repositories"
	"github.com/stretchr/testify/assert"
)

func TestCreatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	redis, _ := redismock.NewClientMock()
	assert.NoError(t, err)
	defer db.Close()

	postRepo := repositories.NewPostRepository(db, redis)

	post := models.Post{
		Title:    "First Post",
		Content:  "This is the content of the first post",
		AuthorId: uuid.New(),
	}

	mock.ExpectPrepare("INSERT INTO posts").
		ExpectExec().
		WithArgs(post.Title, post.Content, post.AuthorId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = postRepo.Create(post.AuthorId, post)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetPostById(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	redis, _ := redismock.NewClientMock()
	postRepo := repositories.NewPostRepository(db, redis)
	postId := uuid.New()

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "likes", "created_at", "author_nick"}).
		AddRow(postId, "First Post", "This is the content of the first post", uuid.New(), 0, time.Now(), "author_nick")

	mock.ExpectQuery("select p.*, u.nick from").
		WithArgs(postId).
		WillReturnRows(rows)

	post, err := postRepo.GetPostById(postId)
	assert.NoError(t, err)
	assert.Equal(t, postId, post.Id)
	assert.Equal(t, "First Post", post.Title)
	assert.Equal(t, "This is the content of the first post", post.Content)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetPosts(t *testing.T) {
	db, mock, err := sqlmock.New()
	redis, _ := redismock.NewClientMock()
	assert.NoError(t, err)
	defer db.Close()

	postRepo := repositories.NewPostRepository(db, redis)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "likes", "created_at", "author_nick"}).
		AddRow(uuid.New(), "First Post", "This is the content of the first post", uuid.New(), 0, time.Now(), "author_nick").
		AddRow(uuid.New(), "Second Post", "This is the content of the second post", uuid.New(), 10, time.Now(), "another_author")

	mock.ExpectQuery("select p.*, u.nick from posts").
		WillReturnRows(rows)

	posts, err := postRepo.GetPosts()
	assert.NoError(t, err)
	assert.Len(t, posts, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdatePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	redis, _ := redismock.NewClientMock()
	assert.NoError(t, err)
	defer db.Close()

	postRepo := repositories.NewPostRepository(db, redis)

	postId := uuid.New()
	post := models.Post{
		Title:   "Updated Title",
		Content: "Updated Content",
	}

	mock.ExpectPrepare("update posts set title").
		ExpectExec().
		WithArgs(post.Title, post.Content, postId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = postRepo.Update(postId, post)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	redis, _ := redismock.NewClientMock()
	assert.NoError(t, err)
	defer db.Close()

	postRepo := repositories.NewPostRepository(db, redis)

	postId := uuid.New()

	mock.ExpectPrepare("delete from posts where id").
		ExpectExec().
		WithArgs(postId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = postRepo.Delete(postId)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestLikePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	redis, _ := redismock.NewClientMock()
	assert.NoError(t, err)
	defer db.Close()

	postRepo := repositories.NewPostRepository(db, redis)

	postId := uuid.New()

	mock.ExpectPrepare(`update posts set likes = likes \+ 1 where id = \$1`).
		ExpectExec().
		WithArgs(postId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = postRepo.Like(postId)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDislikePost(t *testing.T) {
	db, mock, err := sqlmock.New()
	redis, _ := redismock.NewClientMock()
	assert.NoError(t, err)
	defer db.Close()

	postRepo := repositories.NewPostRepository(db, redis)

	postId := uuid.New()

	mock.ExpectPrepare(`
	update posts set likes =
	CASE 
		WHEN likes > 0 THEN likes - 1
		ELSE likes 
	END 
	where id = \$1
	`).
		ExpectExec().
		WithArgs(postId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = postRepo.Dislike(postId)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
