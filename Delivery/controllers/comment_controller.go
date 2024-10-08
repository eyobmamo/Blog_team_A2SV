package controllers

import (
	"blogapp/Domain"
	"blogapp/Utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentController struct {
	commentUseCase Domain.CommentUseCase
}

func NewCommentController(usecase Domain.CommentUseCase) *CommentController {
	return &CommentController{
		commentUseCase: usecase,
	}
}

func (cc *CommentController) CommentOnPost(c *gin.Context) {
	claim, err := Getclaim(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var comment = &Domain.Comment{}
	if err := c.ShouldBindJSON(comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// validate comment
	v := validator.New()
	if err := v.Struct(comment); err != nil {
		fmt.Printf(err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid or missing data", "error": err.Error()})
		return
	}
	
	comment.ID = primitive.NewObjectID()
	comment.PostID = objID
	comment.AuthorID = claim.ID
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()

	err, status := cc.commentUseCase.CommentOnPost(c, comment, objID)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment created successfully"})
}

func (cc *CommentController) GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	comment, err, statuscode := cc.commentUseCase.GetCommentByID(c, objID)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"comment": comment})
}

func (cc *CommentController) EditComment(c *gin.Context) {
	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// get comment
	existingComment, err, statuscode := cc.commentUseCase.GetCommentByID(c, objID)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	// check if user is author of comment
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	isAuthor, err := Utils.IsAuthorOrAdmin(*claims, existingComment.AuthorID)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if !isAuthor {
		c.JSON(401, gin.H{"error": "You are not author of this post"})
		return
	}

	var comment = &Domain.Comment{}
	if err := c.ShouldBindJSON(comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err, status := cc.commentUseCase.EditComment(c, objID, comment)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Comment updated successfully",
		"comment": *comment})
}

func (cc *CommentController) GetUserComments(c *gin.Context) {
	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	comments, err, statuscode := cc.commentUseCase.GetUserComments(c, objID)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"comments": comments})
}

func (cc *CommentController) GetMyComments(c *gin.Context) {
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	comments, err, statuscode := cc.commentUseCase.GetUserComments(c, claims.ID)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"comments": comments})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	objID, err := Utils.StringToObjectId(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// get comment
	existingComment, err, statuscode := cc.commentUseCase.GetCommentByID(c, objID)
	if err != nil {
		c.JSON(statuscode, gin.H{"error": err.Error()})
		return
	}

	// check if user is author of comment
	claims, err := Getclaim(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	isAuthor, err := Utils.IsAuthorOrAdmin(*claims, existingComment.AuthorID)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	if !isAuthor {
		c.JSON(401, gin.H{"error": "You are not author of this post"})
		return
	}

	err, status := cc.commentUseCase.DeleteComment(c, objID)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Comment deleted successfully"})
}
