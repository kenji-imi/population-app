package handler

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	if len(idStr) == 0 {
		log.Fatal("Not Found Delete Id")
	}
	fmt.Println(idStr)

	//row, err := xo.PopulationByID(h.DB, id)
	//if err != nil {
	//	panic(err)
	//}
	//err = row.Delete(h.DB)
	//if err != nil {
	//	panic(err)
	//}
	//c.Status(http.StatusOK)
}
