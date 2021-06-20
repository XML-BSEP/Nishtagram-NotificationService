package interceptor

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"strings"
)

type authUnaryInterceptor struct {

}


type AuthUnaryInterceptor interface {
	UnaryAuthorizationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
	ExtractUserRole(ctx context.Context) (string, error)
}

func NewAuthUnaryInterceptor() AuthUnaryInterceptor {
	return &authUnaryInterceptor{}
}


func (a *authUnaryInterceptor) UnaryAuthorizationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	role, err := a.ExtractUserRole(ctx)

	if err != nil {
		return nil, err
	}

	if role == "" {
		return nil, err
	}

	fullMethod := info.FullMethod

	ok, err := enforce(role, fullMethod, "*")

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, err
	}

	return handler(ctx, req)

}

func (a *authUnaryInterceptor) ExtractUserRole(ctx context.Context) (string, error) {

	tokenString := a.ExtractToken(ctx)
	if tokenString == nil {
		return "ANONYMOUS", nil
	}

	token, err := jwt.Parse(*tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok  {
		userId, ok := claims["role"].(string)
		if !ok {
			return "ANONYMOUS", err
		}

		return strings.ToUpper(userId), nil
	}
	return  "ANONYMOUS", err
}

func (a *authUnaryInterceptor) ExtractToken(context context.Context) *string {

	headers, ok := metadata.FromIncomingContext(context)

	if !ok {
		return nil
	}

	authHeaders := headers["authorization"]

	if len(authHeaders) != 1 {
		return nil
	}

	return &authHeaders[0]

}

func enforce(role string, obj string, act string) (bool, error) {
	m, _ := os.Getwd()

	if !strings.HasSuffix(m, "src")  {
		splits := strings.Split(m, "src")
		wd := splits[0] + "/src"
		if err := os.Chdir(wd); err != nil {
			return false, err
		}
	}

	enforcer, err := casbin.NewEnforcer("grpc/interceptor/auth_interceptor/rbac_model.conf", "grpc/interceptor/auth_interceptor/rbac_policy.csv")
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return false, fmt.Errorf("failed to load policy from DB: %w", err)
	}

	ok, _ := enforcer.Enforce(role, obj, act)

	return ok, nil
}