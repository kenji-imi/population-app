package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kenji-imi/population-app/src/xo"
	"net/http"
)

const (
	readParamDefaultId = int(1)
)

func (h *Handler) Read(c *gin.Context) {
	id := getReadParamId(c)

	row, err := xo.PopulationByID(h.DB, id)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%+v\n", row)
	//fmt.Printf("%+v\n", row.Name)
	res := gin.H{
		"ID":               row.ID,
		"PrefCode":         row.PrefCode,
		"PrefName":         row.PrefName,
		"EraName":          row.EraName,
		"EraYear":          row.EraYear,
		"Year":             row.Year,
		"Population":       row.Population,
		"MalePopulation":   row.MalePopulation,
		"FemalePopulation": row.FemalePopulation,
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, res)
}

func getReadParamId(c *gin.Context) int {
	return getParamInt(c.Param("id"), readParamDefaultId)
}
