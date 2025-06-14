package util

import (
	"context"
	"errors"

	"github.com/tpyle/ksv/lib/types"
)

type NamespaceConfigKey string

const (
	// NamespaceKey is the key used to store the namespace in the context
	NamespaceKey = NamespaceConfigKey("namespace")
)

func AttachNamespaceToContext(namespace string, ctx context.Context) context.Context {
	return context.WithValue(ctx, NamespaceKey, namespace)
}

func GetNamespaceFromContext(ctx context.Context) string {
	namespace, ok := ctx.Value(NamespaceKey).(string)
	if !ok {
		panic("namespace not found in context")
	}

	return namespace
}

func GetNamespace(ksv *types.KSV, ctx context.Context) (*types.Namespace, error) {
	namespace := GetNamespaceFromContext(ctx)
	if namespace == "" {
		return &ksv.DefaultNamespace, nil
	}

	return nil, errors.ErrUnsupported
}
