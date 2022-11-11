package v1

import (
	"context"
	"github.com/Abdurahmonjon/api-gateway/pkg/logger"
	"github.com/Abdurahmonjon/api-gateway/pkg/utils"
	pb "github.com/Abdurahmonjon/api-gateway/protos/studentproto"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateStudent ...
// @Summary CreateStudent
// @Description This API for creating a new student
// @Tags Student
// @Accept  json
// @Produce  json
// @Param Student request body modules.Student true "RegisterStudentRequest"
// @Success 200 {object} modules.Student
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/students/ [post]
func (h handlerV1) CreateStudent(c *gin.Context) {
	var (
		body        pb.RegisterStudentRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.TaskService().RegisterStudent(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		h.log.Error("failed to create student", logger.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// GetStudent ...
// @Summary GetStudent
// @Description This API for getting student details
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} modules.Student
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/students/{id} [get]
func (h *handlerV1) GetStudent(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("username")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	responcse, err := h.serviceManager.TaskService().GetStudent(
		ctx, &pb.GetStudentRequest{
			Username: guid})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		h.log.Error("failed to get task", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, responcse)
}

// UpdateStudent ...
// @Summary UpdateStudent
// @Description This API for updating student
// @Tags student
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Param Task request body modules.Student true "UpdateStudentRequest"
// @Success 200
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/students/{id} [put]
func (h *handlerV1) UpdateStudent(c *gin.Context) {
	var (
		body        pb.UpdateStudentRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}
	body.Id = c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.TaskService().UpdateStudent(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		h.log.Error("failed to update student", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteStudent ...
// @Summary DeleteTask
// @Description This API for deleting student
// @Tags task
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/students/{id} [delete]
func (h *handlerV1) DeleteStudent(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	username := c.Param("username")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.TaskService().DeleteStudent(
		ctx, &pb.DeleteStudentRequest{Username: username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete student", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)

}

// ListStudents ...
// @Summary ListStudents
// @Description This API for getting list of students
// @Tags task
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} modules.StudentList
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/students [get]
func (h *handlerV1) ListStudents(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errStr[0],
		})
		h.log.Error("failed to parse query params json" + errStr[0])
		return
	}
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	response, err := h.serviceManager.TaskService().GetAllStudents(
		ctx, &pb.GetAllStudentsRequest{
			Page:  int32(params.Page),
			Limit: int32(params.Limit),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get list of students", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}
