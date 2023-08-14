package handler

import (
	"errors"
)

// const (
// 	autorizationHeader = "Authorization"
// 	userCtx            = "userId"
// )

func (h *Handler) validatePassword(pass1 string, pass2 string) error {
	if pass1 != pass2 {
		return errors.New("mismatched passwords")
	}

	if pass1 == "" || len(pass1) < 10 {
		return errors.New("bad password")
	}
	return nil
}

// func (h *Handler) userIdentity(c *gin.Context) {
// 	header := c.GetHeader(autorizationHeader)
// 	if header == "" {
// 		newErrorResponse(c, http.StatusUnauthorized, "empty authorization header")
// 		return
// 	}

// 	headerParts := strings.Split(header, " ")
// 	if len(headerParts) != 2 {
// 		newErrorResponse(c, http.StatusUnauthorized, "invalid authorization header")
// 	}

// 	userId, err := h.services.Authorization.ParseToken(headerParts[1])
// 	if err != nil {
// 		newErrorResponse(c, http.StatusUnauthorized, err.Error())
// 	}

// 	c.Set(userCtx, userId)
// }

// func getUserId(c *gin.Context) (int, error) {
// 	id, ok := c.Get(userCtx)
// 	if !ok {
// 		newErrorResponse(c, http.StatusInternalServerError, "userId not found")
// 		return 0, errors.New("userId not found")
// 	}

// 	idInt, ok := id.(int)
// 	if !ok {
// 		newErrorResponse(c, http.StatusInternalServerError, "ivalid type userId")
// 		return 0, errors.New("invalid type userId")
// 	}

// 	return idInt, nil
// }
