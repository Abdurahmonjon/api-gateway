package v1

import (
	"context"
	"fmt"
	pb "github.com/Abdurahmonjon/api-gateway/genproto/gitlab.com/Abdurahmonjon/studentproto"
	"github.com/Abdurahmonjon/api-gateway/pkg/logger"
	"github.com/Abdurahmonjon/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
)

// CreateStudent ...
// @Summary CreateStudent
// @Description This API for creating a new student
// @Tags Students
// @Accept  json
// @Produce  json
// @Param Student request body modules.Student true "RegisterStudentRequest"
// @Success 200 {object} modules.Student
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/student [POST]
func (h handlerV1) CreateStudent(c *gin.Context) {
	var (
		body        pb.RegisterStudentRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	fmt.Println("coming infos:", body.UserName, body.FirstName, body.LastName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	//defer cancel()
	response, err := h.serviceManager.TaskService().RegisterStudent(context.Background(), &body)
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
// @Tags Students
// @Accept  json
// @Produce  json
// @Param id path string true "ID"
// @Success 200 {object} modules.Student
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/student/{id} [GET]
func (h *handlerV1) GetStudent(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Param("id")

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	//defer cancel()

	responcse, err := h.serviceManager.TaskService().GetStudent(
		context.Background(), &pb.GetStudentRequest{
			Id: guid,
		})
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
// @Tags Students
// @Accept  json
// @Produce  json
// @Param id path string true "id"
// @Param Student request body modules.Student true "UpdateStudentRequest"
// @Success 200
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/student/{id} [PUT]
func (h *handlerV1) UpdateStudent(c *gin.Context) {
	var (
		body        pb.UpdateStudentRequest
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	oldId := c.Param("id")
	fmt.Println("id->", oldId)

	err := c.ShouldBindJSON(&body)
	body.Id = oldId
	fmt.Println("body id:", body.Id, body.UserName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", logger.Error(err))
		return
	}

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	//defer cancel()
	response, err := h.serviceManager.TaskService().UpdateStudent(context.Background(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		fmt.Println("id", body.Id)
		h.log.Error("failed to update student", logger.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// DeleteStudent ...
// @Summary DeleteTask
// @Description This API for deleting student
// @Tags Students
// @Accept  json
// @Produce  json
// @Param id path string true "username"
// @Success 200
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/student/{id} [DELETE]
func (h *handlerV1) DeleteStudent(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	id := c.Param("id")
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	//defer cancel()
	response, err := h.serviceManager.TaskService().DeleteStudent(
		context.Background(), &pb.DeleteStudentRequest{Id: id})
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
// @Tags Students
// @Accept  json
// @Produce  json
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success 200 {object} modules.StudentList
// @Failure 400 {object} modules.StandardErrorModel
// @Failure 500 {object} modules.StandardErrorModel
// @Router /v1/students [GET]
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
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	//defer cancel()
	response, err := h.serviceManager.TaskService().GetAllStudents(
		context.Background(), &pb.GetAllStudentsRequest{
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
