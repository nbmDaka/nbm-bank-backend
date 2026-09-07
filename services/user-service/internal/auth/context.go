package auth

import "context"

type ContextKey string

const UserIDKey ContextKey = "user_id"

func SetUserID(
	ctx context.Context,
	id string,
) context.Context {

	return context.WithValue(
		ctx,
		UserIDKey,
		id,
	)
}

func GetUserID(
	ctx context.Context,
) string {

	value := ctx.Value(UserIDKey)

	id, _ := value.(string)

	return id
}
