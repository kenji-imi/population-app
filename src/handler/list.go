package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kenji-imi/population-app/src/xo"
	"net/http"
)

const (
	listParamDefaultOffset = int(0)
	listParamDefaultLimit  = int(10)
)

func (h *Handler) List(c *gin.Context) {
	offset := getListParamOffset(c)
	limit := getListParamLimit(c)

	rows, err := xo.PopulationsByOffsetLimit(h.DB, offset, limit)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v\n", rows)
	res := []gin.H{}
	for _, row := range rows {
		res = append(res,
			gin.H{
				"ID":               row.ID,
				"PrefCode":         row.PrefCode,
				"PrefName":         row.PrefName,
				"EraName":          row.EraName,
				"EraYear":          row.EraYear,
				"Year":             row.Year,
				"Population":       row.Population,
				"MalePopulation":   row.MalePopulation,
				"FemalePopulation": row.FemalePopulation,
			},
		)
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, res)
}

func getListParamOffset(c *gin.Context) int {
	return getParamInt(c.Query("offset"), listParamDefaultOffset)
}

func getListParamLimit(c *gin.Context) int {
	return getParamInt(c.Query("limit"), listParamDefaultLimit)
}
