package api

import (
	"Leadgen/internal/repositories/Building"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddSuccess struct {
	ID int64 `json:"id"`
}
type ErrorResp struct {
	Err string `json:"error"`
}

// BuildingHandlerAdd
//
//	@ID	 building-handler-add
//	@Summary Добавляет Строение, возвращает уникальный номер
//	@Produce json
//	@Produce mpfd
//	@Param name formData string true "Название_записи"
//	@Param city formData string true "Населенный пункт"
//	@Param year formData int false "Год"
//	@Param level formData  int false "Этажность"
//	@Success 201 {object} api.AddSuccess
//	@Failure 400 {object} api.ErrorResp
//	@Router /building/ [post]
func BuildingHandlerAdd(c *gin.Context, db Building.DB) {
	var building Building.Building
	//Проверяем как пришли данные? raw->json или formdata
	err := c.ShouldBind(&building)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResp{Err: err.Error()})
		return
	}
	if building.Title == "" || building.City == "" {
		err = c.BindJSON(&building)
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResp{Err: err.Error()})
		return
	}
	//Добавляем
	Id, err := db.InsertOne(c.Request.Context(), building)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResp{Err: err.Error()})
	} else {
		c.JSON(http.StatusCreated, AddSuccess{ID: Id})
	}
}

// BuildingHandlerGet
//
//	@ID	 building-handler-get
//	@Summary Получает Все объекты с базы. Есть фильтрация
//	@Produce json
//	@Produce mpfd
//	@Param id query int false "id"
//	@Param name query string false "Название_записи"
//	@Param city query string false "Населенный пункт"
//	@Param year query int false "Год"
//	@Param level query  int false "Этажность"
//	@Success 201 {object} []Building.Building
//	@Failure 400 {object} api.ErrorResp
//	@Router /buildings/ [get]
func BuildingHandlerGet(c *gin.Context, db Building.DB) {
	filter := make(map[string]string)
	_ = c.ShouldBind(&filter)
	buildings, err := db.GetBuildings(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else if len(buildings) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No buildings found"})
	} else {
		c.JSON(http.StatusCreated, buildings)
	}

}
