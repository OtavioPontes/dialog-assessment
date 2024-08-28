package controllers

import (
	"encoding/json"
	"errors"
	"io"

	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/otaviopontes/api-go/src/authentication"
	"github.com/otaviopontes/api-go/src/database"
	"github.com/otaviopontes/api-go/src/models"
	"github.com/otaviopontes/api-go/src/repositories"
	"github.com/otaviopontes/api-go/src/responses"
)

// @Summary      Create a new post
// @Description  Creates a post associated with the authenticated user.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        title  body    string  true  "Post title"
// @Param        content  body    string  true  "Post content"
// @Success      201   {object}  models.Post
// @Failure      400   {object}  responses.ErrorResponse
// @Failure      401   {object}  responses.ErrorResponse
// @Failure      500   {object}  responses.ErrorResponse
// @Router       /posts [post]
func CreatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post

	err = json.Unmarshal(bodyRequest, &post)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorId = userId

	if err := post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	redis, err := database.ConnectRedis()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db, redis)

	err = repository.Create(userId, post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)

}

// @Summary      Get all posts
// @Description  Retrieves a list of all posts from the database.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Post
// @Failure      500  {object} responses.ErrorResponse
// @Router       /posts [get]
func GetPosts(w http.ResponseWriter, r *http.Request) {

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	redis, err := database.ConnectRedis()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db, redis)

	posts, err := repository.GetPosts()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)

}

// @Summary      Get a post by ID
// @Description  Retrieves a single post by its ID.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      200  {object}  models.Post
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /posts/{id} [get]
func GetPost(w http.ResponseWriter, r *http.Request) {
	postId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	redis, err := database.ConnectRedis()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db, redis)

	post, err := repository.GetPostById(postId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

// @Summary      Update an existing post
// @Description  Updates a post that belongs to the authenticated user.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id    path      string       true  "Post ID"
// @Param        post  body      models.Post  true  "Updated post data"
// @Success      204
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      401  {object}  responses.ErrorResponse
// @Failure      403  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /posts/{id} [put]
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	postId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	redis, err := database.ConnectRedis()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db, redis)

	postSaved, err := repository.GetPostById(postId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postSaved.AuthorId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("it is not possible to update a post that is not yours"))
		return
	}
	bodyRequest, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post

	err = json.Unmarshal(bodyRequest, &post)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = post.Prepare()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = repository.Update(postId, post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// @Summary      Delete a post
// @Description  Deletes a post that belongs to the authenticated user.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      204
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      401  {object}  responses.ErrorResponse
// @Failure      403  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /posts/{id} [delete]
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userId, err := authentication.ExtractUserId(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	postId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	redis, err := database.ConnectRedis()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db, redis)

	postSaved, err := repository.GetPostById(postId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if postSaved.AuthorId != userId {
		responses.Error(w, http.StatusForbidden, errors.New("it is not possible to delete a post that is not yours"))
		return
	}

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = repository.Delete(postId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// @Summary      Like a post
// @Description  Allows a user to like a post by its ID.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      204
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /posts/{id}/like [post]
func LikePost(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	redis, err := database.ConnectRedis()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db, redis)

	err = repository.Like(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}

// @Summary      Dislike a post
// @Description  Allows a user to dislike a post by its ID.
// @Tags         Posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      204
// @Failure      400  {object}  responses.ErrorResponse
// @Failure      500  {object}  responses.ErrorResponse
// @Router       /posts/{id}/dislike [post]
func DislikePost(w http.ResponseWriter, r *http.Request) {
	userId, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	redis, err := database.ConnectRedis()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewPostRepository(db, redis)

	err = repository.Dislike(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)

}
