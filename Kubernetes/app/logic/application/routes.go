package application

import (
	"github.com/gin-gonic/gin"
	"kubertest/logic/application/rates"
	"net/http"
	"strconv"
)

func BuildUpdateCacheHandler(service rates.IService) gin.HandlerFunc{
	return func(context *gin.Context) {
		err := service.UpdateCache()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal_error",
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
		return
	}
}

func BuildConvertHandler(service rates.IService) gin.HandlerFunc{
	return func(context *gin.Context) {
		from := context.Query("from")
		to := context.Query("to")
		strAmount := context.Query("amount")

		if len(from) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "bad_request",
				"message": "missing from parameter",
			})
			return
		}else if len(to) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "bad_request",
				"message": "missing to parameter",
			})
			return
		}

		amount, err := strconv.ParseFloat(strAmount, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "bad_request",
				"message": "invalid float passed as amount",
			})
			return
		}

		res, err := service.Convert(from,to,amount)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal_error",
				"message": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, gin.H{
			"result": res,
			"unit": to,
		})
		return
	}
}