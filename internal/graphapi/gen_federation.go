// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

var (
	ErrUnknownType  = errors.New("unknown type")
	ErrTypeNotFound = errors.New("type not found")
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}

	var sdl []string

	for _, src := range sources {
		if src.BuiltIn {
			continue
		}
		sdl = append(sdl, src.Input)
	}

	return fedruntime.Service{
		SDL: strings.Join(sdl, "\n"),
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) []fedruntime.Entity {
	list := make([]fedruntime.Entity, len(representations))

	repsMap := map[string]struct {
		i []int
		r []map[string]interface{}
	}{}

	// We group entities by typename so that we can parallelize their resolution.
	// This is particularly helpful when there are entity groups in multi mode.
	buildRepresentationGroups := func(reps []map[string]interface{}) {
		for i, rep := range reps {
			typeName, ok := rep["__typename"].(string)
			if !ok {
				// If there is no __typename, we just skip the representation;
				// we just won't be resolving these unknown types.
				ec.Error(ctx, errors.New("__typename must be an existing string"))
				continue
			}

			_r := repsMap[typeName]
			_r.i = append(_r.i, i)
			_r.r = append(_r.r, rep)
			repsMap[typeName] = _r
		}
	}

	isMulti := func(typeName string) bool {
		switch typeName {
		default:
			return false
		}
	}

	resolveEntity := func(ctx context.Context, typeName string, rep map[string]interface{}, idx []int, i int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {
		case "Annotation":
			resolverName, err := entityResolverNameForAnnotation(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Annotation": %w`, err)
			}
			switch resolverName {

			case "findAnnotationByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findAnnotationByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindAnnotationByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Annotation": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "AnnotationNamespace":
			resolverName, err := entityResolverNameForAnnotationNamespace(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "AnnotationNamespace": %w`, err)
			}
			switch resolverName {

			case "findAnnotationNamespaceByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findAnnotationNamespaceByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindAnnotationNamespaceByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "AnnotationNamespace": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Metadata":
			resolverName, err := entityResolverNameForMetadata(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Metadata": %w`, err)
			}
			switch resolverName {

			case "findMetadataByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findMetadataByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindMetadataByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Metadata": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			case "findMetadataByNodeID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["nodeID"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findMetadataByNodeID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindMetadataByNodeID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Metadata": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "MetadataNode":
			resolverName, err := entityResolverNameForMetadataNode(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "MetadataNode": %w`, err)
			}
			switch resolverName {

			case "findMetadataNodeByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findMetadataNodeByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindMetadataNodeByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "MetadataNode": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "ResourceOwner":
			resolverName, err := entityResolverNameForResourceOwner(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "ResourceOwner": %w`, err)
			}
			switch resolverName {

			case "findResourceOwnerByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findResourceOwnerByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindResourceOwnerByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "ResourceOwner": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "Status":
			resolverName, err := entityResolverNameForStatus(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "Status": %w`, err)
			}
			switch resolverName {

			case "findStatusByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findStatusByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindStatusByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "Status": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "StatusNamespace":
			resolverName, err := entityResolverNameForStatusNamespace(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "StatusNamespace": %w`, err)
			}
			switch resolverName {

			case "findStatusNamespaceByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findStatusNamespaceByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindStatusNamespaceByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "StatusNamespace": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}
		case "StatusOwner":
			resolverName, err := entityResolverNameForStatusOwner(ctx, rep)
			if err != nil {
				return fmt.Errorf(`finding resolver for Entity "StatusOwner": %w`, err)
			}
			switch resolverName {

			case "findStatusOwnerByID":
				id0, err := ec.unmarshalNID2goᚗinfratographerᚗcomᚋxᚋgidxᚐPrefixedID(ctx, rep["id"])
				if err != nil {
					return fmt.Errorf(`unmarshalling param 0 for findStatusOwnerByID(): %w`, err)
				}
				entity, err := ec.resolvers.Entity().FindStatusOwnerByID(ctx, id0)
				if err != nil {
					return fmt.Errorf(`resolving Entity "StatusOwner": %w`, err)
				}

				list[idx[i]] = entity
				return nil
			}

		}
		return fmt.Errorf("%w: %s", ErrUnknownType, typeName)
	}

	resolveManyEntities := func(ctx context.Context, typeName string, reps []map[string]interface{}, idx []int) (err error) {
		// we need to do our own panic handling, because we may be called in a
		// goroutine, where the usual panic handling can't catch us
		defer func() {
			if r := recover(); r != nil {
				err = ec.Recover(ctx, r)
			}
		}()

		switch typeName {

		default:
			return errors.New("unknown type: " + typeName)
		}
	}

	resolveEntityGroup := func(typeName string, reps []map[string]interface{}, idx []int) {
		if isMulti(typeName) {
			err := resolveManyEntities(ctx, typeName, reps, idx)
			if err != nil {
				ec.Error(ctx, err)
			}
		} else {
			// if there are multiple entities to resolve, parallelize (similar to
			// graphql.FieldSet.Dispatch)
			var e sync.WaitGroup
			e.Add(len(reps))
			for i, rep := range reps {
				i, rep := i, rep
				go func(i int, rep map[string]interface{}) {
					err := resolveEntity(ctx, typeName, rep, idx, i)
					if err != nil {
						ec.Error(ctx, err)
					}
					e.Done()
				}(i, rep)
			}
			e.Wait()
		}
	}
	buildRepresentationGroups(representations)

	switch len(repsMap) {
	case 0:
		return list
	case 1:
		for typeName, reps := range repsMap {
			resolveEntityGroup(typeName, reps.r, reps.i)
		}
		return list
	default:
		var g sync.WaitGroup
		g.Add(len(repsMap))
		for typeName, reps := range repsMap {
			go func(typeName string, reps []map[string]interface{}, idx []int) {
				resolveEntityGroup(typeName, reps, idx)
				g.Done()
			}(typeName, reps.r, reps.i)
		}
		g.Wait()
		return list
	}
}

func entityResolverNameForAnnotation(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findAnnotationByID", nil
	}
	return "", fmt.Errorf("%w for Annotation", ErrTypeNotFound)
}

func entityResolverNameForAnnotationNamespace(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findAnnotationNamespaceByID", nil
	}
	return "", fmt.Errorf("%w for AnnotationNamespace", ErrTypeNotFound)
}

func entityResolverNameForMetadata(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findMetadataByID", nil
	}
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["nodeID"]; !ok {
			break
		}
		return "findMetadataByNodeID", nil
	}
	return "", fmt.Errorf("%w for Metadata", ErrTypeNotFound)
}

func entityResolverNameForMetadataNode(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findMetadataNodeByID", nil
	}
	return "", fmt.Errorf("%w for MetadataNode", ErrTypeNotFound)
}

func entityResolverNameForResourceOwner(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findResourceOwnerByID", nil
	}
	return "", fmt.Errorf("%w for ResourceOwner", ErrTypeNotFound)
}

func entityResolverNameForStatus(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findStatusByID", nil
	}
	return "", fmt.Errorf("%w for Status", ErrTypeNotFound)
}

func entityResolverNameForStatusNamespace(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findStatusNamespaceByID", nil
	}
	return "", fmt.Errorf("%w for StatusNamespace", ErrTypeNotFound)
}

func entityResolverNameForStatusOwner(ctx context.Context, rep map[string]interface{}) (string, error) {
	for {
		var (
			m   map[string]interface{}
			val interface{}
			ok  bool
		)
		_ = val
		m = rep
		if _, ok = m["id"]; !ok {
			break
		}
		return "findStatusOwnerByID", nil
	}
	return "", fmt.Errorf("%w for StatusOwner", ErrTypeNotFound)
}
