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
) (string, bool) {

	value := ctx.Value(UserIDKey)

	id, ok := value.(string)

	return id, ok
}
