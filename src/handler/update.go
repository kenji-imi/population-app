package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	if len(idStr) == 0 {
		log.Fatal("Not Found Update Id")
	}
	fmt.Println(idStr)

	//row, err := xo.PopulationByID(h.DB, id)
	//if err != nil {
	//	panic(err)
	//}
	//err = row.Update(h.DB)
	//if err != nil {
	//	panic(err)
	//}
	//c.Status(http.StatusOK)
}
