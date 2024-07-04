package handlers

import (
	"net/http"

	"github.com/TheusLab/ASN-Project/backend/elasticsearch"
	"github.com/TheusLab/ASN-Project/backend/utils"
	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	results, err := elasticsearch.Search(query)
	if err != nil {
		utils.Log.Error().Err(err).Msg("Search error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, results)
}
